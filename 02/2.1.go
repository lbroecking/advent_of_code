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

		if strings.ContainsAny(password, letter) && strings.Count(password, letter) <= policy_max && strings.Count(password, letter) >= policy_min {
			count++
		}
	}
	fmt.Println(count)
}
