package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input_day_1.txt")

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

	for i := 0; i < len(txtlines); i++ {
		j, _ := strconv.Atoi(txtlines[i])
		numbers = append(numbers, j)
	}

	//1.1
	for _, num := range numbers {
		for _, num2 := range numbers {
			if num+num2 == 2020 {
				fmt.Println(num * num2)
			}
		}
	}

	//1.2
	for _, num := range numbers {
		for _, num2 := range numbers {
			for _, num3 := range numbers {
				if num+num2+num3 == 2020 {
					fmt.Println(num * num2 * num3)
				}
			}
		}
	}
}
