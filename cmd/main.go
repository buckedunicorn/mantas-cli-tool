package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/buckedunicorn/mantas-cli-tool/internal/fileops"
	"github.com/buckedunicorn/mantas-cli-tool/internal/netops"
	"github.com/buckedunicorn/mantas-cli-tool/internal/sysdiag"
)

func main() {
	mainFlagSet := flag.NewFlagSet("mantas-cli-tool", flag.ExitOnError)
	mainFlagSet.Usage = customUsage

	var fileCmd, netCmd, sysCmd bool
	mainFlagSet.BoolVar(&fileCmd, "file", false, "File operations: create, read, update, delete")
	mainFlagSet.BoolVar(&netCmd, "net", false, "Network operations: ping, fetch")
	mainFlagSet.BoolVar(&sysCmd, "sys", false, "System diagnostics: mem, cpu")

	if err := mainFlagSet.Parse(os.Args[1:]); err != nil {
		fmt.Println("Error parsing flags:", err)
		return
	}

	args := mainFlagSet.Args()
	if len(args) == 0 {
		mainFlagSet.Usage()
		return
	}

	if fileCmd {
		fileops.HandleFileOperations(args)
	} else if netCmd {
		if len(args) < 2 {
			fmt.Println("Network operation requires a command and target")
			return
		}
		netops.HandleNetworkOperations(args)
	} else if sysCmd {
		sysdiag.HandleSystemDiagnostics(args)
	} else {
		mainFlagSet.Usage()
	}
}

func customUsage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Println("Example commands:")
	fmt.Println("  -file create <filename>")
	fmt.Println("  -net ping <host>")
	fmt.Println("  -sys mem")
}
