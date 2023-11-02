package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Info()
	PrintRules()

	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")

		command, _ := reader.ReadString('\n')
		command = RemoveControlSeq(command)

		ExecuteCommand(command)
	}
}
