package main

import "fmt"

func main() {
	runCli()
}

func runCli() {
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	fmt.Print(input)
}
