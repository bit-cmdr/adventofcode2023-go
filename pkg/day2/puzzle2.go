package day2

import (
	"fmt"
	"os"
	"strings"
)

func Puzzle2() {
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

	var minCubeSetGames []minCubeSetGame
	for _, game := range games {
		minCubeSetGames = append(minCubeSetGames, game.toMinCubeSetGame())
	}

	fmt.Printf("Min cube set games: %+v\n", minCubeSetGames)

	var gamePowers []gamePower
	for _, minCubeSetGame := range minCubeSetGames {
		gamePowers = append(gamePowers, minCubeSetGame.toGamePower())
	}

	fmt.Printf("Game powers: %+v\n", gamePowers)

	sum := 0
	for _, game := range gamePowers {
		sum += game.power
	}

	fmt.Println(sum)
}
