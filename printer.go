package main

import (
	"fmt"
	"github.com/fatih/color"
)

func Info() {
	fmt.Println("<><><><><><><><><><><>")
	fmt.Println("<> Chesnokov Arkady <>")
	fmt.Println("<>    AI LAB: 3     <>")
	fmt.Println("<><><><><><><><><><><>")
}

func PrintGreen(message string) {
	color.Green(message)
}
