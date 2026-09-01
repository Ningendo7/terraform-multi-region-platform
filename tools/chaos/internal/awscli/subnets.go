package awscli

import (
	"context"
	"encoding/json"
	"fmt"
)

// CIDRsForAZ returns the CIDR block of every subnet in the given
// availability zone. Not filtered by project tag — the caller already
// scopes to a specific cluster via kubeContext, and every node it
// checks is necessarily inside that cluster's own VPC already.
func CIDRsForAZ(ctx context.Context, region, az string) ([]string, error) {
	out, err := Run(ctx, region,
		"ec2", "describe-subnets",
		"--filters", "Name=availability-zone,Values="+az,
		"--query", "Subnets[].CidrBlock",
		"--output", "json",
	)
	if err != nil {
		return nil, err
	}

	var cidrs []string
	if err := json.Unmarshal([]byte(out), &cidrs); err != nil {
		return nil, fmt.Errorf("parsing subnet CIDRs: %w", err)
	}

	if len(cidrs) == 0 {
		return nil, fmt.Errorf("no subnets found in %s - is the AZ name correct?", az)
	}

	return cidrs, nil
}
