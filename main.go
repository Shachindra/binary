package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	version   = "1.0.0"   // Replace with your version
	codeHash  = "unknown" // Replace with your code hash
	goVersion = "unknown" // Replace with your Go version
)

func main() {
	// Check if the first argument is "version" or "saymyname"
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version":
			// Define CLI flags for version
			jsonFlag := flag.Bool("json", false, "Output the version and code hash in JSON format")
			flag.CommandLine.Parse(os.Args[2:])

			// Handle the --json flag
			if *jsonFlag {
				// Output in JSON format
				fmt.Printf("{\"version\": \"%s\", \"codehash\": \"%s\", \"goVersion\": \"%s\"}\n", version, codeHash, goVersion)
			} else {
				// Output in plain text format
				fmt.Println("Version:", version)
				fmt.Println("Code Hash:", codeHash)
				fmt.Println("Go Version:", goVersion)
			}
			return

		case "saymyname":
			// Define CLI flags for saymyname
			jsonFlag := flag.Bool("json", false, "Output the greeting in JSON format")
			nameFlag := flag.String("name", "", "Your name to greet")
			flag.CommandLine.Parse(os.Args[2:])

			// Handle the --name flag
			if *nameFlag == "" {
				fmt.Println("Error: --name flag is required for saymyname")
				fmt.Println("Usage: ./binary saymyname --name <your_name> [--json]")
				return
			}

			// Handle the --json flag
			if *jsonFlag {
				// Output in JSON format
				fmt.Printf("{\"greeting\": \"Hello %s\"}\n", *nameFlag)
			} else {
				// Output in plain text format
				fmt.Printf("Hello %s\n", *nameFlag)
			}
			return
		}
	}

	// Default message if no valid command is provided
	fmt.Println("Usage:")
	fmt.Println("  ./binary version [--json]")
	fmt.Println("  ./binary saymyname --name <your_name> [--json]")
}
