package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file_content, _ := ioutil.ReadFile("input_day_4.txt")
	string_content := string(file_content)
	input := strings.Split(string_content, "\n\n")

	check_elements := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}
	count := 0
	for _, key := range input {
		check_count := 0
		for _, element := range check_elements {
			if strings.Contains(string(key), string(element)) {
				check_count++
			}
		}
		if check_count >= len(check_elements) {
			count++
		}
	}

	fmt.Println(count)

}
