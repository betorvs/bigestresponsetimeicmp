package config

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

var (
	// DestinationHost variable
	DestinationHost string
)

// info to print how to use it
func info() {
	fmt.Println("Usage: sudo biggestresponsetime REMOTE_HOST")
	fmt.Println("You must run this command with root permissions")
	fmt.Println("")
}

func init() {
	if len(os.Args) == 1 {
		info()
		os.Exit(1)
	}
	DestinationHost = os.Args[1]
	if DestinationHost == "version" {
		fmt.Println("biggestresponsetime Version: ", Version)
		os.Exit(0)
	}
	if DestinationHost == "help" {
		info()
		os.Exit(0)
	}
	cmd := exec.Command("id", "-u")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error trying to find permissions")
		os.Exit(1)
	}
	i, err := strconv.Atoi(string(output[:len(output)-1]))
	if err != nil {
		fmt.Println("Error trying to read permissions")
		os.Exit(1)
	}
	if i != 0 {
		fmt.Println("Error: You must run this command with root permissions. Use: sudo")
		os.Exit(2)
	}
}
