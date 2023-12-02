package main

import (
	"fmt"
	"os"

	"github.com/bit-cmdr/adventofcode2023-go/pkg/day1"
	"github.com/bit-cmdr/adventofcode2023-go/pkg/day2"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No puzzle specified")
		return
	}

	fmt.Printf("Running %s\n", args[0])
	switch args[0] {
	case "day1":
		fmt.Println("Running puzzle 1")
		day1.Puzzle1()
		fmt.Println("Running puzzle 2")
		day1.Puzzle2()

	case "day2":
		fmt.Println("Running puzzle 1")
		day2.Puzzle1()
		fmt.Println("Running puzzle 2")
		day2.Puzzle2()

	default:
		fmt.Println("No known puzzle specified")
	}
}
