package day2

import (
	"fmt"
	"os"
	"strings"
)

func Puzzle1() {
	bs, err := os.ReadFile("pkg/day2/input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	input := string(bs)
	rawInputs := strings.Split(input, "\n")

	var games []game
	for _, rawInput := range rawInputs {
		games = append(games, toGame(rawInput))
	}

	fmt.Printf("All games: %+v\n", games)

	validGames := filterValidGames(games)

	fmt.Printf("Valid games: %+v\n", validGames)

	sum := 0
	for _, game := range validGames {
		sum += game.id
	}

	fmt.Println(sum)
}

func filterValidGames(games []game) []game {
	var validGames []game
	for _, game := range games {
		valid := true
		for _, cubeSet := range game.cubeSets {
			if cubeSet.red > CUBE_LIMITS.red || cubeSet.green > CUBE_LIMITS.green || cubeSet.blue > CUBE_LIMITS.blue {
				valid = false
			}
		}

		if valid {
			validGames = append(validGames, game)
		}
	}

	return validGames
}
