package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type bag struct {
	color    string
	contents []string
	amounts  []int
}

func main() {
	file, _ := ioutil.ReadFile("input_day_7.txt")
	goal_bag := "shiny gold"

	// Part 1
	bags_input := bagsInput(file)
	count := 0
	for _, bag := range bags_input {
		if checkGoal(goal_bag, bag, bags_input) {
			count++
		}
	}
	fmt.Println("Part 1: ", count) //144

	//Part 2
	count = getBagInBag(goal_bag, bags_input)
	fmt.Println("Part 2: ", count) //5956

}

func bagsInput(input []byte) map[string]bag {
	input_strings := strings.Split(string(input), ".\n")
	bags := make(map[string]bag)
	for _, rule := range input_strings {
		n_contents := strings.Count(rule, ",") + 1
		split_string := strings.Split(rule, " ")
		main_bag := split_string[0] + " " + split_string[1]

		var content_bags []string
		var content_amounts []int
		current_amount_pos, current_color_1, current_color_2 := 4, 5, 6
		for i := 0; i < n_contents; i++ {
			color := split_string[current_color_1] + " " + split_string[current_color_2]
			content_bags = append(content_bags, color)
			amount, _ := strconv.Atoi(split_string[current_amount_pos])
			content_amounts = append(content_amounts, amount)
			current_amount_pos += 4
			current_color_1 += 4
			current_color_2 += 4
		}
		temp_bag := bag{color: main_bag, contents: content_bags, amounts: content_amounts}
		bags[temp_bag.color] = temp_bag
	}
	return bags
}

func checkGoal(goal string, bag bag, bags map[string]bag) bool {
	if len(bag.contents) == 0 {
		return false
	} else {
		temp_bool := false
		for _, temp_bag := range bag.contents {
			if temp_bag == goal {
				if !temp_bool {
					temp_bool = true
				}
			} else {
				if !temp_bool {
					temp_bool = checkGoal(goal, bags[temp_bag], bags)
				}
			}
		}
		return temp_bool
	}
}

func getBagInBag(goal string, bags map[string]bag) int {
	total_count := 0
	for i, bag := range bags[goal].contents {
		total_count += bags[goal].amounts[i]
		total_count += bags[goal].amounts[i] * getBagInBag(bag, bags)
	}
	return total_count
}
