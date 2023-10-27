package main

import (
	"fmt"
	"github.com/fatih/color"
)

func info() {
	fmt.Println("<><><><><><><><><><><>")
	fmt.Println("<> Chesnokov Arkady <>")
	fmt.Println("<>    AI LAB: 3     <>")
	fmt.Println("<><><><><><><><><><><>")
}

func printGreen(message string) {
	color.Green(message)
}
