# chaos

A small CLI for running chaos experiments against this platform's EKS clusters — built to prove resilience claims rather than assume them.

## Build

	cd tools/chaos
	go build -o chaos ./cmd/chaos

## Safety model

Every experiment requires either `--dry-run` (selects and logs targets, touches nothing) or explicit `--yes` — no interactive confirmation prompt, since this tool is meant to be scriptable, and a prompt blocking on stdin is the wrong failure mode there. Always run with `--dry-run` first against a cluster you haven't tested against before.

Anything involving randomness (`pod-kill`, `node-failure --count`) accepts `--seed` for reproducibility — if something goes wrong, you can replay exactly which target got picked.

## Experiments

### pod-kill

Deletes one pod matching a label selector.

	chaos run pod-kill --context <ctx> --namespace demo --label app=demo-app --dry-run
	chaos run pod-kill --context <ctx> --namespace demo --label app=demo-app --yes

### node-failure

Cordons and evicts a set of real nodes — either a random `--count`, or every node in a given `--az` (that's also how this simulates an availability-zone failure: `--az` with no `--count` targets every node in that zone). Nodes restore automatically after `--window` (default 5m), or earlier via `chaos restore` from a separate terminal.

Refuses to target more than 50% of the cluster's nodes unless `--force` is passed.

	chaos run node-failure --context <ctx> --region us-east-1 --count 1 --dry-run
	chaos run node-failure --context <ctx> --region us-east-1 --az us-east-1a --yes
	chaos restore --context <ctx>

Only meaningful against real, persistent nodes (Karpenter-provisioned EC2) — Fargate nodes are ephemeral per-pod, so cordoning one has no lasting effect.

## Every command needs

	--context   the kubeconfig context of the target cluster (`kubectl config get-contexts`)
