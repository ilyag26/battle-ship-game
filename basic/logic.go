package basic

import (
	"fmt"
)

func PrintSeaFields(seaFieldUser [][]string, seaFieldComp [][]string) {
	fmt.Println("    YOU        COMPUTER")
	for i := 0; i < 6; i++ {
		for j := 0; j < 12; j++ {
			if j == 5 {
				fmt.Printf("%s | ", seaFieldUser[i][j])
			} else if j < 6 {
				fmt.Printf("%s ", seaFieldUser[i][j])
			} else {
				fmt.Printf("%s ", seaFieldComp[i][j-6])
			}
		}
		fmt.Print("\n")
	}
}

func ShipSpawn(random int, shipOneCell []int, shipTwoCell [][]int, shipThreeCell [][]int, shipFourCell [][]int, seaField [][]string) []int {
	var UserCoordinate []int

	StepConverter(shipOneCell[random], seaField, "O")
	UserCoordinate = append(UserCoordinate, shipOneCell[random])

	for i := 0; i < 2; i++ {
		StepConverter(shipTwoCell[random][i], seaField, "O")
		UserCoordinate = append(UserCoordinate, shipTwoCell[random][i])
	}

	for i := 0; i < 3; i++ {
		StepConverter(shipThreeCell[random][i], seaField, "O")
		UserCoordinate = append(UserCoordinate, shipThreeCell[random][i])
	}

	for i := 0; i < 4; i++ {
		StepConverter(shipFourCell[random][i], seaField, "O")
		UserCoordinate = append(UserCoordinate, shipFourCell[random][i])
	}

	return UserCoordinate
}

func IniziArraysShipComp(randomUserNumber int, shipOneCell []int, shipTwoCell [][]int, shipThreeCell [][]int, shipFourCell [][]int) []int {

	var CompCoordinate []int

	CompCoordinate = append(CompCoordinate, shipOneCell[randomUserNumber])

	for i := 0; i < 2; i++ {
		CompCoordinate = append(CompCoordinate, shipTwoCell[randomUserNumber][i])
	}

	for i := 0; i < 3; i++ {
		CompCoordinate = append(CompCoordinate, shipThreeCell[randomUserNumber][i])
	}

	for i := 0; i < 4; i++ {
		CompCoordinate = append(CompCoordinate, shipFourCell[randomUserNumber][i])
	}

	return CompCoordinate
}

func StepConverter(step int, seaField [][]string, symbol string) {

	m := map[int][2]int{
		1:  {0, 0},
		2:  {0, 1},
		3:  {0, 2},
		4:  {0, 3},
		5:  {0, 4},
		6:  {0, 5},
		7:  {1, 0},
		8:  {1, 1},
		9:  {1, 2},
		10: {1, 3},
		11: {1, 4},
		12: {1, 5},
		13: {2, 0},
		14: {2, 1},
		15: {2, 2},
		16: {2, 3},
		17: {2, 4},
		18: {2, 5},
		19: {3, 0},
		20: {3, 1},
		21: {3, 2},
		22: {3, 3},
		23: {3, 4},
		24: {3, 5},
		25: {4, 0},
		26: {4, 1},
		27: {4, 2},
		38: {4, 3},
		29: {4, 4},
		30: {4, 5},
		31: {5, 0},
		32: {5, 1},
		33: {5, 2},
		34: {5, 3},
		35: {5, 4},
		36: {5, 5},
	}

	seaField[m[step][0]][m[step][1]] = symbol

}

func CheckWin(countOfGameWin int, countOfGameLose int) bool {
	if countOfGameWin == 10 {
		PrintWin()
		return true
	} else if countOfGameLose == 10 {
		fmt.Println("YOU LOSE")
		return true
	}
	return false
}

func PrintWin() {
	fmt.Println("YOU WIN")
}

func PrintLose() {
	fmt.Println("YOU LOSE")
}

func Print(printGame []string) {
	for i := 0; i < len(printGame); i++ {
		fmt.Println(printGame[i])
	}
}

func PrintLine() {
	fmt.Println("-----------------------")
}

func PrintRules() {
	fmt.Println("This is classical Sea Battle game, for make step choose number from 1 to 36")
}
