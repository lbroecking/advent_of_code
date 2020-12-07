package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file_content, _ := ioutil.ReadFile("input_day_6.txt")
	string_content := string(file_content)
	input := strings.Split(string_content, "\n\n")

	count := 0
	check_elements := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, line := range input {
		for _, element := range check_elements {
			if strings.Contains(string(line), string(element)) {
				count++
			}
		}
	}
	fmt.Println("Part 1:", count) // 7120

	//Part 2
	count = 0
	for _, line := range input {
		len_input := len(strings.Split(string(line), "\n"))
		var characters [26]int
		for _, character := range line {
			if character != 10 {
				characters[character-97] += 1
				if characters[character-97] == len_input {
					count++
				}
			}
		}
	}

	fmt.Println("Part 2:", count) // 3570
}
