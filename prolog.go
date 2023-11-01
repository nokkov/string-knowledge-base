package main

import (
	"fmt"
	"github.com/mndrix/golog"
	"strings"
)

var knowledgeBase = `
		ingredient(nether_wart).
		ingredient(glowstone_dust).
		ingredient(glistering_melon).
		ingredient(spider_eyes).
		ingredient(magma_cream).
		ingredient(sugar).
		ingredient(blaze_powder).
		ingredient(ghast_tears).
		ingredient(redstone_dust).
		ingredient(gunpowder).
		ingredient(golden_carrots).
		ingredient(rabbits_foot).
		ingredient(turtle_shell).
		ingredient(phantom_membrane).
		ingredient(secret_ingredient).

		%funny ingredients... is it fun huh?!
		ingredient(water).
		ingredient(tears).
		ingredient(mysql).
		ingredient(javac).
		ingredient(linux).
		
		%3 arguments are passed => this is definitely a potion
		isPotion(X, Y, Z, Result) :-
			nonvar(X),
			nonvar(Y),
			nonvar(Z),
			Result = 'Potion!'.

		potion(sugar, spider_eyes, gunpowder).
		potion(sugar, javac, gunpowder).
		potion(linux, javac, tears).
`

func InitKnowledgeBase() golog.Machine {
	kb := golog.NewMachine().Consult(knowledgeBase)
	return kb
}

func AddIngredient(ingredient string) {
	knowledgeBase += fmt.Sprintf("\ningredient(%s).", ingredient)
}

func AddPotion(X, Y, Z string) {
	knowledgeBase += fmt.Sprintf("\npotion(%s, %s, %s).", X, Y, Z)
}

func IsIngredientExists(ingredient string) bool {
	ingredient = "ingredient(" + ingredient + ")."
	return strings.Contains(knowledgeBase, ingredient)
}

func IsPotionExists(X, Y, Z string) bool {
	potion := fmt.Sprintf("potion(%s, %s, %s).", X, Y, Z)
	return strings.Contains(knowledgeBase, potion)
}
