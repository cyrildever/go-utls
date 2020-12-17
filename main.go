package main

import (
	"fmt"

	"github.com/cyrildever/go-utls/common/logger"
	"github.com/cyrildever/go-utls/model"
)

func main() {
	log := logger.Init("main", "main")

	// Just to make sure everything is ok
	h := model.Hash("")
	if h.NonEmpty() {
		panic("something weird happens")
	}

	// Print special information for library users
	fmt.Println("")
	fmt.Println("COPYRIGHT NOTICE")
	fmt.Println("================")
	fmt.Println("This library holds some modules I found useful in my Go developments. It's available under a MIT license.")
	fmt.Println("")
	fmt.Println("© 2020-2021 Cyril Dever. All rights reserved.")
	fmt.Println("")

	log.Info("Enjoy~")
}
