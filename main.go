package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("octopus.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Pink background
	fmt.Print("\033[48;5;225m")
	fmt.Print("\033[2J\033[H")

	// Title
	fmt.Println("\033[38;5;206m")
	fmt.Println("  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("         🐙  Welcome to the Octopus Terminal  🐙")
	fmt.Println("  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("\033[0m")

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			fmt.Println()
			continue
		}
		// Color each character
		for _, ch := range line {
			switch ch {
			case '█':
				fmt.Print("\033[38;5;0m██") // black eyes
			case '▓':
				fmt.Print("\033[38;5;213m▓▓") // dark pink tentacle tips
			case '▒':
				fmt.Print("\033[38;5;206m▒▒") // light pink body
			default:
				fmt.Print(string(ch))
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("\033[38;5;206m  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("            Have a squishy day! 🩷")
	fmt.Println("  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\033[0m")
}
