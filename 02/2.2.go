package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("passwords")

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

	count := 0

	for _, x := range txtlines {
		split := strings.Split(x, ":")
		password := split[1]
		policy := strings.Split(split[0], "-")
		policy_split := strings.Split(policy[1], " ")

		policy_min, _ := strconv.Atoi(policy[0])
		policy_max, _ := strconv.Atoi(policy_split[0])

		letter := policy_split[1]

		var passwordSlice []string
		for _, r := range password {
			c := string(r)
			passwordSlice = append(passwordSlice, c)
		}

		// (first OR last letter correct) AND (first AND last letter not equal)
		if passwordSlice == nil {
			return
		}
		if (passwordSlice[policy_min] == letter || passwordSlice[policy_max] == letter) && (passwordSlice[policy_min] != passwordSlice[policy_max]) {
			count++
		}
	}
	fmt.Println(count)
}
