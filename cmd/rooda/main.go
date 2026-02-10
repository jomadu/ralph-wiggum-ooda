package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	Version   = "dev"
	CommitSHA = "unknown"
	BuildDate = "unknown"
)

func main() {
	versionFlag := flag.Bool("version", false, "Print version information")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("rooda %s\n", Version)
		fmt.Printf("Commit: %s\n", CommitSHA)
		fmt.Printf("Built: %s\n", BuildDate)
		os.Exit(0)
	}

	fmt.Fprintln(os.Stderr, "rooda: OODA loop implementation not yet available")
	os.Exit(1)
}
