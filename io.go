package main

import (
	"strings"
)

func ExecuteCommand(command string) {
	tokens := strings.Split(command, " ")

	switch tokens[0] {
	case "ing:":
		switch tokens[1] {

		case "!":
			PrintAllIngredients()

		case "?":
			PrintRandomPotion()
		}

	}

}
