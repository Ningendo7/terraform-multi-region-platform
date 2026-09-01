package main

import (
	"fmt"
	"os"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/probe/internal/probecli"
)

func main() {
	if err := probecli.Execute(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
