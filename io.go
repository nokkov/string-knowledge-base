package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RemoveControlSeq(command string) string {
	command = strings.ReplaceAll(command, "\r", "")
	command = strings.ReplaceAll(command, "\n", "")

	return command
}

func ExecuteCommand(command string) {
	tokens := strings.Split(command, " ")

	switch tokens[0] {
	case "ing:":
		switch tokens[1] {
		//FIXME: fix PrintAllIngredients() & PrintRandomPotion()
		case "!":
			PrintAllIngredients()

		case "?":
			PrintRandomPotion()

		case "*":
			reader := bufio.NewReader(os.Stdin)
			ingredient, _ := reader.ReadString('\n')

			if !IsIngredientExists(ingredient) {
				AddIngredient(ingredient)
				PrintGreen("Successfully added a new ingredient!")
			} else {
				PrintRed("Ingredient already exists!")
			}

		case "***":
			reader := bufio.NewReader(os.Stdin)
			potion, _ := reader.ReadString('\n')
			ingredients := strings.Split(potion, " ")

			if len(ingredients) != 3 {
				errString := fmt.Sprintf("Wrong amount of args, expected 3, got %d", len(ingredients))
				PrintRed(errString)
			} else {
				if !IsPotionExists(ingredients[0], ingredients[1], ingredients[2]) {
					//TODO: check ingredients
					AddPotion(ingredients[0], ingredients[1], ingredients[2])
					PrintGreen("Successfully added a new potion!")
				} else {
					PrintRed("Potion already exists!")
				}
			}
		}

	case "stop":
		PrintRed("Terminating program...")
		os.Exit(0)

	default:
		PrintRed("Unknown command!")
	}

}
