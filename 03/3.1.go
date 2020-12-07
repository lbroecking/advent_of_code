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

	currentPostition := make([]int, 2)
	position := make([]int, 2)
	position[0] = 3
	position[1] = 1
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

	fmt.Println(count)
}
