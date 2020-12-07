package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input_day_3.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	_ = file.Close()

	position := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	product := 1

	for i, _ := range position {
		product *= countTrees(txtlines, position[i])
	}

	fmt.Println(product)
}

func countTrees(txtlines []string, position []int) int {

	currentPostition := make([]int, 2)
	width := len(txtlines[0])
	count := 0
	for currentPostition[1] < len(txtlines) {
		y := currentPostition[1]
		line := txtlines[y]
		c := line[currentPostition[0]]
		if c == '#' {
			count++
		}

		currentPostition[0] = (currentPostition[0] + position[0]) % width
		currentPostition[1] += position[1]
	}

	return count
}
