package gol_tourofgo

import gol_app "go-learning.com/learning/app"

func RunTOG() {
	togApp := map[int]gol_app.App{
		1: {Title: "Loops And Functions", Execute: runLaF},
	}

	gol_app.RunAppWithSelections(togApp, "Tour of GO!")
}
