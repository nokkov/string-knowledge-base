package main

import "github.com/mndrix/golog"

func initKnowledgeBase() golog.Machine {
	kb := golog.NewMachine().Consult(`
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
	`)

	return kb
}
