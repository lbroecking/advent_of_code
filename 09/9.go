package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input_day_9.txt")

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

	preamble_size := 25
	var lines []int
	var previous_numbers []int

	for k := 5; k < len(txtlines); k++ {
		j, _ := strconv.Atoi(txtlines[k])
		lines = append(lines, j)
	}

	for i, number := range lines {
		if i <= preamble_size {
			continue
		} else {
			previous_numbers = lines[i-preamble_size : i]
		}
		if !checkValidity(number, previous_numbers) {
			fmt.Println("The answer to puzzle 1 is :", number)                      //133015568
			fmt.Println("The answer to puzzle 2 is :", findWeakness(number, lines)) //16107959
			break
		}
	}
}

func checkValidity(number int, pre_numbers []int) bool {
	for _, key := range pre_numbers {
		for _, key1 := range pre_numbers {
			if key+key1 == number && key != key1 {
				return true
			}
		}
	}
	return false
}

func findWeakness(number int, input_lines []int) int {
	for i, key := range input_lines {
		list := []int{key}
		for _, key1 := range input_lines[i+1:] {
			list = append(list, key1)

			result := 0
			for _, v := range list {
				result += v
			}

			if result == number {
				min, max := findMinMax(list)
				return min + max
			}
		}
	}
	return 0
}

func findMinMax(array []int) (int, int) {
	max := array[0]
	min := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
