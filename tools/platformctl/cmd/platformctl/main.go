package main

import (
	"fmt"
	"os"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/platformctl/internal/checks"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("usage: platformctl <command>")
		return
	}
	
	switch os.Args[1] {

	case "doctor":

		fmt.Println("Platform Environment Check")
		fmt.Println("----------------------------")

		checks.CheckCommand("terraform")
		checks.CheckCommand("aws")
		checks.CheckCommand("kubectl")
		checks.CheckCommand("helm")
		checks.CheckCommand("git")
		
	default:

		fmt.Println("unknown command:", os.Args[1])
		
	}
}