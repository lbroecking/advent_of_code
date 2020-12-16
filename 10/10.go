package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input_day_10.txt")

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

	var numbers []int

	for k := 0; k < len(txtlines); k++ {
		j, _ := strconv.Atoi(txtlines[k])
		numbers = append(numbers, j)
	}

	sort.Ints(numbers)

	fmt.Println("Part 1:", part1(numbers)) //2380
	fmt.Println("Part 2:", part2(numbers)) //48358655787008

}

func part1(numbers []int) int {
	countOne := 1
	countThree := 1

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i+1]-numbers[i] == 1 {
			countOne++
		} else if numbers[i+1]-numbers[i] == 3 {
			countThree++
		}
	}

	return countOne * countThree
}

func part2(numbers []int) int {
	variants := map[int]int{0: 1}

	for _, num := range numbers {
		variants[num] = variants[num-1] + variants[num-2] + variants[num-3]
	}

	return variants[numbers[len(numbers)-1]]
}
