package kube

import (
	"context"
	"encoding/json"
	"fmt"
	"net/netip"
)

type k8sNode struct {
	Metadata struct {
		Name        string            `json:"name"`
		Annotations map[string]string `json:"annotations"`
	} `json:"metadata"`
	Status struct {
		Addresses []struct {
			Type    string `json:"type"`
			Address string `json:"address"`
		} `json:"addresses"`
	} `json:"status"`
}

func (n k8sNode) internalIP() (string, bool) {
	for _, a := range n.Status.Addresses {
		if a.Type == "InternalIP" {
			return a.Address, true
		}
	}
	return "", false
}

func getNodes(ctx context.Context, kubeContext string) ([]k8sNode, error) {
	out, err := Run(ctx, kubeContext, "get", "nodes", "-o", "json")
	if err != nil {
		return nil, err
	}

	var list struct {
		Items []k8sNode `json:"items"`
	}
	if err := json.Unmarshal([]byte(out), &list); err != nil {
		return nil, fmt.Errorf("parsing node list: %w", err)
	}

	return list.Items, nil
}

// nodeNamesInCIDRs returns the name of every node whose internal IP falls
// inside any of the given CIDR blocks.
func nodeNamesInCIDRs(nodes []k8sNode, cidrs []string) ([]string, error) {
	prefixes := make([]netip.Prefix, 0, len(cidrs))
	for _, c := range cidrs {
		p, err := netip.ParsePrefix(c)
		if err != nil {
			return nil, fmt.Errorf("parsing CIDR %q: %w", c, err)
		}
		prefixes = append(prefixes, p)
	}

	var matched []string
	for _, n := range nodes {
		ip, ok := n.internalIP()
		if !ok {
			continue
		}

		addr, err := netip.ParseAddr(ip)
		if err != nil {
			continue
		}

		for _, p := range prefixes {
			if p.Contains(addr) {
				matched = append(matched, n.Metadata.Name)
				break
			}
		}
	}

	return matched, nil
}

func nodeNameWithAnnotation(nodes []k8sNode, key, value string) []string {
	var matched []string
	for _, n := range nodes {
		if n.Metadata.Annotations[key] == value {
			matched = append(matched, n.Metadata.Name)
		}
	}
	return matched
}

func cordonNode(ctx context.Context, kubeContext, name string) error {
	_, err := Run(ctx, kubeContext, "cordon", name)
	return err
}

func uncordonNode(ctx context.Context, kubeContext, name string) error {
	_, err := Run(ctx, kubeContext, "uncordon", name)
	return err
}

func annotateNode(ctx context.Context, kubeContext, name, key, value string) error {
	_, err := Run(ctx, kubeContext, "annotate", "node", name, fmt.Sprintf("%s=%s", key, value), "--overwrite")
	return err
}

func removeNodeAnnotation(ctx context.Context, kubeContext, name, key string) error {
	_, err := Run(ctx, kubeContext, "annotate", "node", name, key+"-")
	return err
}

// evictPods deletes every non-DaemonSet pod currently running on node,
// forcing the scheduler to place them elsewhere — cordoning alone only
// stops *new* pods landing there, it doesn't move what's already running.
func evictPods(ctx context.Context, kubeContext, nodeName string) error {
	out, err := Run(ctx, kubeContext, "get", "pods", "-A",
		"--field-selector", "spec.nodeName="+nodeName,
		"-o", "json")
	if err != nil {
		return err
	}

	var podList struct {
		Items []struct {
			Metadata struct {
				Name            string `json:"name"`
				Namespace       string `json:"namespace"`
				OwnerReferences []struct {
					Kind string `json:"kind"`
				} `json:"ownerReferences"`
			} `json:"metadata"`
		} `json:"items"`
	}
	if err := json.Unmarshal([]byte(out), &podList); err != nil {
		return fmt.Errorf("parsing pod list for node %s: %w", nodeName, err)
	}

	for _, p := range podList.Items {
		isDaemonSet := false
		for _, ref := range p.Metadata.OwnerReferences {
			if ref.Kind == "DaemonSet" {
				isDaemonSet = true
				break
			}
		}
		if isDaemonSet {
			continue
		}

		if _, err := Run(ctx, kubeContext, "delete", "pod", p.Metadata.Name, "-n", p.Metadata.Namespace, "--wait=false"); err != nil {
			return fmt.Errorf("evicting pod %s/%s: %w", p.Metadata.Namespace, p.Metadata.Name, err)
		}
	}

	return nil
}
