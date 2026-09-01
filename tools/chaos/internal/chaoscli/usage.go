package chaoscli

import "fmt"

func Usage() {

	fmt.Print(`
Usage:
	chaos <command> [args]
	chaos run <experiment> [flags]

Commands:

	run       run an experiment
	restore   restore any nodes still cordoned by a previous node-failure run

Experiments:

	pod-kill      delete one pod matching a label selector
	node-failure  cordon and evict a set of real nodes, by count or by AZ

Examples:
	chaos run pod-kill --context <ctx> --namespace demo --label app=demo-app --dry-run
	chaos run node-failure --context <ctx> --region us-east-1 --count 1 --dry-run
	chaos run node-failure --context <ctx> --region us-east-1 --az us-east-1a --yes
	chaos restore --context <ctx>

`)
}
