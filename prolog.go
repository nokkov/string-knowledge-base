package main

import "github.com/mndrix/golog"

func initKnowledgeBase() golog.Machine {
	kb := golog.NewMachine().Consult(`
		ingredient(nether_wart).
		ingredient(glowstone_dust)
	`)

	return kb
}
