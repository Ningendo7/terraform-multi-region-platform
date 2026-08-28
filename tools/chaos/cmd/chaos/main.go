package main

import (
	"fmt"
	"os"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/chaos/internal/chaoscli"
)

func main() {
	if err := chaoscli.Execute(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}