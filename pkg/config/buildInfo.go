package config

import (
	"fmt"
)

var (
	// These values will be filled at build time, these are just defaults.
	VersionMajor = "x"
	VersionMinor = "x"
	BuildNumber  = "x"
	BuildTime    = "-"
	BuildBranch  = "-"
	BuildCommit  = "-"
	Builder      = "a person, manually"
	BuildMachine = "-"
)

func PrintBuildInfo() {
	// Build number is major.minor.buildID
	fmt.Println("=========================================================")
	fmt.Printf(" Version: %s.%s.%s\n", VersionMajor, VersionMinor, BuildNumber)
	fmt.Printf("Built at: %s\n", BuildTime)
	fmt.Printf("  Branch: %s\n", BuildBranch)
	fmt.Printf("Commit: %s\n", BuildCommit)
	fmt.Printf("Built by: %s\n", Builder)
	fmt.Printf("      on: %s\n", BuildMachine)
	fmt.Println("=========================================================")
}
