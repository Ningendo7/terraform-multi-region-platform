package chaoscli

import "fmt"

func Usage() {

	fmt.Print(`
Usage:
	chaos <command> <expirement> [flags]
	
Commands:
	
	run
	
Experiments:

	pod-kill
	
Examples:
	chaos run pod-kill --context <kube-context> --namespace demo --label app=demo-app --dry-run
	chaos run pod-kill --context <kube-context> --namespace demo --label app=demo-app
	
`)
}