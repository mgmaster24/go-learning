package app

import (
	"fmt"
	"sort"
)

type App struct {
	Title   string
	Execute func()
}

func RunAppWithSelections(apps map[int]App, title string) {
	fmt.Println(title)
	keys := make([]int, 0, len(apps))
	for k := range apps {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d: %s\n", k, apps[k].Title)
	}

	exercise := 0
	fmt.Print("Please choose the exercise you would like to run: ")
	fmt.Scanln(&exercise)
	if app, ok := apps[exercise]; ok {
		app.Execute()
	}
}
