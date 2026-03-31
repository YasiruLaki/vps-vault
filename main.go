package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command. Run 'vps --help' for more information.")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("Insufficient arguments for 'add' command. Run 'vps --help' for more information.")
			return
		}
		name := os.Args[2]
		ip := os.Args[3]
		user := os.Args[4]
		port := 22 
		if len(os.Args) >= 7 {
			fmt.Sscanf(os.Args[5], "%d", &port)
		}
		addVPS(name, ip, user, port)
	case "list":
		listVPS()
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Insufficient arguments for 'remove' command. Run 'vps --help' for more information.")
			return
		}
		name := os.Args[2]
		removeVPS(name)

	case "--help":
		printHelp()
	default:
		fmt.Printf("Unknown command: %s. Run 'vps --help' for more information.\n", command)
	}
}