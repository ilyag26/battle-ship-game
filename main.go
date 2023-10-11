package main

import (
	"fmt"
	"math/rand"
	b "seabattle/basic"

	"golang.org/x/exp/slices"
)

var (
	steps        int
	message      string
	winCheckUser int
	winCheckComp int
	checkWinner  bool
	wrong        bool
)

func main() {

	printGame := []string{
		"Welcme to GO Sea Battel",
		"-----------------------",
		"For make any step chose number from 1 to 36",
		"If you will destroy all ships yo will win!"}

	seaFieldUser := [][]string{
		{"~", "~", "~", "~", "~", "~"},
		{"~", "~", "~", "~", "~", "~"},
		{"~", "~", "~", "~", "~", "~"},
		{"~", "~", "~", "~", "~", "~"},
		{"~", "~", "~", "~", "~", "~"},
		{"~", "~", "~", "~", "~", "~"}}

	seaFieldComp := [][]string{
		{"~", "~", "~", "~", "~", "~"},
		{"~", "~", "~", "~", "~", "~"},
		{"~", "~", "~", "~", "~", "~"},
		{"~", "~", "~", "~", "~", "~"},
		{"~", "~", "~", "~", "~", "~"},
		{"~", "~", "~", "~", "~", "~"}}

	//from 0 to 3 for random position of ships
	min := 0
	max := 3
	//from 1 to 36 for random computer's step
	min1 := 1
	max1 := 36
	userShipsRand := rand.Intn(max-min) + min
	compShipsRand := rand.Intn(max-min) + min

	//position for ships
	shipOneCell := []int{1, 10, 12, 30}
	shipTwoCell := [][]int{{10, 11}, {25, 31}, {27, 33}, {12, 18}}
	shipThreeCell := [][]int{{21, 27, 33}, {27, 28, 29}, {15, 16, 17}, {25, 26, 27}}
	shipFourCell := [][]int{{13, 19, 25, 31}, {6, 12, 18, 24}, {1, 2, 3, 4}, {2, 3, 4, 5}}

	//starting game
	b.Print(printGame)
	b.PrintLine()
	UserCoordinate := b.ShipSpawn(userShipsRand, shipOneCell, shipTwoCell, shipThreeCell, shipFourCell, seaFieldUser)
	CompCoordinate := b.IniziArraysShipComp(compShipsRand, shipOneCell, shipTwoCell, shipThreeCell, shipFourCell)
	b.PrintSeaFields(seaFieldUser, seaFieldComp)
	for {
		for {
			if checkWinner {
				break
			}
			wrong = false
			fmt.Print("Make Step:")
			fmt.Scan(&steps)
			if slices.Contains(CompCoordinate, steps) {
				b.StepConverter(steps, seaFieldComp, "O")
				winCheckUser++
				message = "YES!"
			} else {
				b.StepConverter(steps, seaFieldComp, "X")
				message = "NO! WRONG!"
				wrong = true
			}
			b.PrintSeaFields(seaFieldUser, seaFieldComp)
			fmt.Println(message)
			checkWinner = b.CheckWin(winCheckUser, winCheckComp)
			if wrong {
				break
			}
		}
		if checkWinner {
			break
		}
		for {
			if checkWinner {
				break
			}
			wrong = false
			fmt.Println("Computer's Step:")
			randomStepComp := rand.Intn(max1-min1) + min1
			fmt.Println()
			if slices.Contains(UserCoordinate, randomStepComp) {
				b.StepConverter(randomStepComp, seaFieldUser, "-")
				winCheckComp++
				message = "YES!"
			} else {
				b.StepConverter(randomStepComp, seaFieldUser, "X")
				message = "NO! WRONG!"
				wrong = true
			}
			b.PrintSeaFields(seaFieldUser, seaFieldComp)
			fmt.Println(message)
			checkWinner = b.CheckWin(winCheckUser, winCheckComp)
			if wrong {
				break
			}
		}
		if checkWinner {
			break
		}
	}

}
