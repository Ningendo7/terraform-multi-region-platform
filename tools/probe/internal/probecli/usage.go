package probecli

import "fmt"

func Usage() {

	fmt.Print(`
Usage:
	probe <command> [flags]

Commands:

	check     check one or more URLs, report pass/fail per target
	recover   watch one URL, measure how long an outage lasts once one starts

Examples:
	probe check --url https://demo.tendo.pro --url https://demo-us-east-2.tendo.pro
	probe check --url https://app.tendo.pro --interval 10s
	probe recover --url https://app.tendo.pro --max-wait 5m

`)
}
