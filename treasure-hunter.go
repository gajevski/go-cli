package main

import "fmt"

func main() {
	findTreasure()
}

func findTreasure() {

	x := 0
	y := 0

	treasureX := 5
	treasureY := 5

	fmt.Println(`Welcome to Treasure Quest!
	
Objective: Find the hidden treasure.
	
Controls:
	
"W" to move North
"S" to move South
"A" to move West
"E" to move East

Navigate wisely and watch for clues. Your adventure begins now! Good luck!`)
	for {

		fmt.Print("Direction: ")
		var input string
		fmt.Scan(&input)

		switch input {
		case "W", "w":
			fmt.Println("You're heading north")
			y++
			fmt.Println(y)
		case "S", "s":
			fmt.Println("You're heading south")
			y--
			fmt.Println(y)
		case "A", "a":
			fmt.Println("You're heading west")
			x--
			fmt.Println(x)
		case "D", "d":
			fmt.Println("You're heading east")
			x++
			fmt.Println(x)
		default:
			fmt.Println("Wrong direction!")
		}

		if x == treasureX && y == treasureY {
			fmt.Println("You found a treasure!")
			return
		}
	}
}
