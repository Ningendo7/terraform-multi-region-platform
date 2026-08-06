package main

import (
	"fmt"
	"os"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/platformctl/internal/cli"
)

func main() {
	if err := cli.Execute(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}