package day1

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Puzzle1() {
	bs, err := os.ReadFile("pkg/day1/input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	input := string(bs)
	rawInputs := strings.Split(input, "\n")

	var nums []int
	for _, rawInput := range rawInputs {
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
