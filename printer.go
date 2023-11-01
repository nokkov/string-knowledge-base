package main

import (
	"fmt"
	"github.com/fatih/color"
)

// Info Print author's name, lab's name
func Info() {
	fmt.Println("<><><><><><><><><><><>")
	fmt.Println("<> Chesnokov Arkady <>")
	fmt.Println("<>    AI LAB: 3     <>")
	fmt.Println("<><><><><><><><><><><>")
}

// PrintGreen Print green text to console
func PrintGreen(message string) {
	color.Green(message)
}

// PrintRed Print red text to console
func PrintRed(message string) {
	color.Red(message)
}

// PrintRules Print rules how to make requests to knowledge base
func PrintRules() {
	PrintRed("HOW TO\n")
	PrintRed("ing: ingredients (1..n) => poison exists or not\n")
	PrintRed("ing: ! => list of all ingredients\n")
	PrintRed("ing: ? => receipt of random potion\n")
	PrintRed("ing: * => new ingredient\n")
	PrintRed("ing: *** => new potion\n")
	PrintRed("stop => exit program\n")
}
