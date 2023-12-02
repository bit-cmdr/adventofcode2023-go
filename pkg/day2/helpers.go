package day2

import (
	"strconv"
	"strings"
)

func toGame(str string) game {
	game := game{}
	gameMatch := gameRegex.FindAllString(str, -1)
	if len(gameMatch) == 0 {
		return game
	}

	gameNumMatch := numRegex.FindAllString(gameMatch[0], -1)
	if len(gameNumMatch) == 0 {
		return game
	}
	gameId, _ := strconv.Atoi(gameNumMatch[0])
	game.id = gameId

	cubeSetsStrings := strings.Split(str, gameMatch[0])
	cubeSetStrings := strings.Split(cubeSetsStrings[1], ";")

	for _, cubeSetString := range cubeSetStrings {
		game.cubeSets = append(game.cubeSets, toCubeSet(cubeSetString))
	}

	return game
}

func toCubeSet(str string) cubeSet {
	cs := cubeSet{
		red:   0,
		green: 0,
		blue:  0,
	}

	colorMatches := colorRegex.FindAllString(str, -1)
	if len(colorMatches) == 0 {
		return cs
	}

	for _, colorMatch := range colorMatches {
		numMatches := numRegex.FindAllString(colorMatch, -1)
		if len(numMatches) == 0 {
			continue
		}

		if strings.Contains(colorMatch, "red") {
			cs.red, _ = strconv.Atoi(numMatches[0])
		} else if strings.Contains(colorMatch, "green") {
			cs.green, _ = strconv.Atoi(numMatches[0])
		} else if strings.Contains(colorMatch, "blue") {
			cs.blue, _ = strconv.Atoi(numMatches[0])
		}
	}

	return cs
}
