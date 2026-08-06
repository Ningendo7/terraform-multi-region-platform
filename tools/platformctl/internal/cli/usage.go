package cli

import "fmt"

func Usage() {

	fmt.Println(`
Usage:
	platformctl <command> <scope> <stack>

Commands:

	init
	plan
	apply
	destroy

Examples:
	platformctl init global iam
	platformctl plan global route53
	platformctl apply region us-east-1 eks
	platformctl destroy global iam

`)
}
