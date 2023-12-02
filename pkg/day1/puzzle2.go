package day1

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type wordNums struct {
	word    string
	numWord string
}

var KNOWN_WORD_NUMS = []wordNums{
	{
		word:    "one",
		numWord: "o1e",
	},
	{
		word:    "two",
		numWord: "t2o",
	},
	{
		word:    "three",
		numWord: "t3ree",
	},
	{
		word:    "four",
		numWord: "f4ur",
	},
	{
		word:    "five",
		numWord: "f5ve",
	},
	{
		word:    "six",
		numWord: "s6x",
	},
	{
		word:    "seven",
		numWord: "s7ven",
	},
	{
		word:    "eight",
		numWord: "e8ght",
	},
	{
		word:    "nine",
		numWord: "n9ne",
	},
}

func Puzzle2() {
	bs, err := os.ReadFile("pkg/day1/input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	input := string(bs)
	rawInputs := strings.Split(input, "\n")

	var nums []int
	for _, rawInput := range rawInputs {
		for _, knownWordNum := range KNOWN_WORD_NUMS {
			rawInput = strings.Replace(rawInput, knownWordNum.word, knownWordNum.numWord, -1)
		}

		rx := regexp.MustCompile(`\d`)
		matches := rx.FindAllString(rawInput, -1)
		if len(matches) == 0 {
			continue
		}

		fn, _ := strconv.Atoi(matches[0])
		ln, _ := strconv.Atoi(matches[len(matches)-1])
		nums = append(nums, fn*10+ln)
	}

	fmt.Println(nums)

	sum := 0
	for _, n := range nums {
		sum += n
	}

	fmt.Println(sum)
}
