package main

import (
	"fmt"
	"os"
	"os/exec"
)

var jsonPath = "./vps_data.json"

func addVPS(name, ip, user string, port int) {
	if name == "" || ip == "" || user == "" {
		fmt.Println("All fields (name, ip, username, port) are required to add a VPS. Run 'vps --help' for more information.")
		return
	}

	if port <= 0 || port > 65535 {
		fmt.Println("Port must be a valid number between 1 and 65535. Run 'vps --help' for more information.")
		return
	}

	if port == 0 {
		port = 22
	}

	newVPS := VPSData{
		Name: name,
		IP:   ip,
		User: user,
		Port: port,
	}

	vpsList := loadVPS(jsonPath)
	if nameExists(name, vpsList) {
		fmt.Printf("Error: a VPS with the name '%s' already exists. Choose a different name. Run 'vps --help' for more information.\n", name)
		return
	}

	vpsList = append(vpsList, newVPS)
	if err := saveVPS(vpsList, jsonPath); err != nil {
		fmt.Println("Error: failed to save VPS data.")
		return
	}

	fmt.Printf("Success: VPS '%s' added. Run 'vps list' to see your updated VPS list.\n", name)
}

func listVPS() {
	vpsList := loadVPS(jsonPath)
	if len(vpsList) == 0 {
		fmt.Println("No VPS entries found. Try adding a VPS first using the 'add' command. Run 'vps --help' for more information.")
		return
	}

	fmt.Println("Your VPS List:")
	for _, vps := range vpsList {
		fmt.Printf("- Name: %s, IP: %s, User: %s, Port: %d\n", vps.Name, vps.IP, vps.User, vps.Port)
	}
}

func removeVPS(name string) {
	if name == "" {
		fmt.Println("The name field is required to remove a VPS. Run 'vps --help' for more information.")
		return
	}

	vpsList := loadVPS(jsonPath)
	var updatedList []VPSData
	found := false

	for _, vps := range vpsList {
		if vps.Name != name {
			updatedList = append(updatedList, vps)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Printf("Error: no VPS found with the name '%s'. Run 'vps list' to see your current VPS entries.\n", name)
		return
	}

	if err := saveVPS(updatedList, jsonPath); err != nil {
		fmt.Println("Error: failed to save VPS data.")
		return
	}

	fmt.Printf("Success: VPS '%s' removed. Run 'vps list' to see your updated VPS list.\n", name)
}

func connectVPS(name string) {
	if name == "" {
		fmt.Println("The name field is required to connect to a VPS. Run 'vps --help' for more information.")
		return
	}

	vpsList := loadVPS(jsonPath)
	var targetVPS *VPSData

	for _, vps := range vpsList {
		if vps.Name == name {
			targetVPS = &vps
			break
		}
	}

	if targetVPS == nil {
		fmt.Printf("Error: no VPS found with the name '%s'. Run 'vps list' to see your current VPS entries.\n", name)
		return
	}

	fmt.Printf("Info: connecting to VPS '%s' at %s:%d as user '%s'...\n", targetVPS.Name, targetVPS.IP, targetVPS.Port, targetVPS.User)

	remote := fmt.Sprintf("%s@%s", targetVPS.User, targetVPS.IP)
	cmd := exec.Command("ssh", "-p", fmt.Sprintf("%d", targetVPS.Port), remote)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: failed to connect to VPS '%s'. Check your SSH configuration and try again.\n", targetVPS.Name)
		return
	}

	fmt.Printf("Info: SSH session ended for VPS '%s'.\n", targetVPS.Name)
}

func printHelp() {
	helpText := `
VPS Manager - A simple command-line tool to manage your VPS connections.

Usage:
  vps <command> [arguments]

Commands:
  add <name> <ip> <username> [port]   Add a new VPS entry. Port is optional and defaults to 22.
  list                                 List all saved VPS entries.
  remove <name>                       Remove a VPS entry by name.
  connect <name>                      Connect to a VPS using its name.

Examples:
	vps add myserver 203.0.113.10 ubuntu 22
  vps list
  vps remove myserver
  vps connect myserver


Made with ❤️ by Yasiru Lakintha. For support or to contribute, visit our GitHub repository.

Thank you for using VPS Manager! For more information, visit our GitHub repository or contact support.
`
	fmt.Println(helpText)
}
