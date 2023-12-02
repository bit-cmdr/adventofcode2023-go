package day2

import "regexp"

var gameRegex = regexp.MustCompile(`Game (\d+):\s`)
var colorRegex = regexp.MustCompile(`(\d+)\s(red|green|blue)`)
var numRegex = regexp.MustCompile(`\d+`)
