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
	switch {
	case step == 1:
		seaField[0][0] = symbol
	case step == 2:
		seaField[0][1] = symbol
	case step == 3:
		seaField[0][2] = symbol
	case step == 4:
		seaField[0][3] = symbol
	case step == 5:
		seaField[0][4] = symbol
	case step == 6:
		seaField[0][5] = symbol
	case step == 7:
		seaField[1][0] = symbol
	case step == 8:
		seaField[1][1] = symbol
	case step == 9:
		seaField[1][2] = symbol
	case step == 10:
		seaField[1][3] = symbol
	case step == 11:
		seaField[1][4] = symbol
	case step == 12:
		seaField[1][5] = symbol
	case step == 13:
		seaField[2][0] = symbol
	case step == 14:
		seaField[2][1] = symbol
	case step == 15:
		seaField[2][2] = symbol
	case step == 16:
		seaField[2][3] = symbol
	case step == 17:
		seaField[2][4] = symbol
	case step == 18:
		seaField[2][5] = symbol
	case step == 19:
		seaField[3][0] = symbol
	case step == 20:
		seaField[3][1] = symbol
	case step == 21:
		seaField[3][2] = symbol
	case step == 22:
		seaField[3][3] = symbol
	case step == 23:
		seaField[3][4] = symbol
	case step == 24:
		seaField[3][5] = symbol
	case step == 25:
		seaField[4][0] = symbol
	case step == 26:
		seaField[4][1] = symbol
	case step == 27:
		seaField[4][2] = symbol
	case step == 28:
		seaField[4][3] = symbol
	case step == 29:
		seaField[4][4] = symbol
	case step == 30:
		seaField[4][5] = symbol
	case step == 31:
		seaField[5][0] = symbol
	case step == 32:
		seaField[5][1] = symbol
	case step == 33:
		seaField[5][2] = symbol
	case step == 34:
		seaField[5][3] = symbol
	case step == 35:
		seaField[5][4] = symbol
	case step == 36:
		seaField[5][5] = symbol
	}
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
