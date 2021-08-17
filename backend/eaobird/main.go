package eaobird

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
)

//go:embed eao_bird_circle.txt
var birdCircle string

func Print() {
	bird, err := strconv.Unquote(birdCircle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bird)
	fmt.Println()
	fmt.Println()
	fmt.Println("Poesitory - A project by the Edgar Allan Ohms, FRC Team 5276")
	fmt.Println()
	fmt.Println()
}
