package day2

type cubeSet struct {
	red   int
	green int
	blue  int
}

type game struct {
	id       int
	cubeSets []cubeSet
}

type minCubeSetGame struct {
	id         int
	minCubeSet cubeSet
}

type gamePower struct {
	id    int
	power int
}

var CUBE_LIMITS = cubeSet{
	red:   12,
	green: 13,
	blue:  14,
}

func (game game) toMinCubeSetGame() minCubeSetGame {
	min := minCubeSetGame{
		id: game.id,
		minCubeSet: cubeSet{
			red:   0,
			green: 0,
			blue:  0,
		},
	}

	for _, cubeSet := range game.cubeSets {
		if cubeSet.red > min.minCubeSet.red {
			min.minCubeSet.red = cubeSet.red
		}

		if cubeSet.green > min.minCubeSet.green {
			min.minCubeSet.green = cubeSet.green
		}

		if cubeSet.blue > min.minCubeSet.blue {
			min.minCubeSet.blue = cubeSet.blue
		}
	}

	return min
}

func (game minCubeSetGame) toGamePower() gamePower {
	return gamePower{
		id:    game.id,
		power: game.minCubeSet.red * game.minCubeSet.green * game.minCubeSet.blue,
	}
}
