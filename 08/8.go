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
	file, err := os.Open("input_day_8.txt")

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

	part1(txtlines)
	_, instructions, err := part2(txtlines)
	if err != nil {
		panic(err)
	}

	for _, i := range instructions {
		l := txtlines[i]
		cmd := l[:3]

		if cmd == "nop" {
			txtlines[i] = strings.ReplaceAll(l, "nop", "jmp")
		} else if cmd == "jmp" {
			txtlines[i] = strings.ReplaceAll(l, "jmp", "nop")
		} else {
			continue
		}

		acc, _, err := part2(txtlines)
		if err != nil {
			panic(err)
		}

		if acc != -1 {
			fmt.Printf("Modified line %d from '%s' to '%s'\n", i+1, l, txtlines[i])
			fmt.Println("Part 2:", acc)
			break
		} else {
			txtlines[i] = l
		}
	}

}

func part1(txtlines []string) {
	m := map[int]struct{}{}

	i := 0

	accumulator := 0
	for i < len(txtlines) {
		if _, ok := m[i]; ok {
			fmt.Println("Part 1:", accumulator) //1521
			return
		}
		key := txtlines[i][:3]
		value := txtlines[i][4:]

		switch key {
		case "acc":
			acc, _ := strconv.Atoi(value)
			accumulator += acc
		case "jmp":
			jmp, _ := strconv.Atoi(value)
			i += jmp
			continue
		}
		m[i] = struct{}{}
		i++
	}
}

func part2(txtlines []string) (accumulator int, run []int, err error) {
	m := map[int]struct{}{}

	i := 0
	for i < len(txtlines) {
		run = append(run, i)
		if _, ok := m[i]; ok {
			return -1, run, nil
		}
		m[i] = struct{}{}

		key := txtlines[i][:3]
		value := txtlines[i][4:]

		switch key {
		case "nop":
			i++
		case "acc":
			acc, _ := strconv.Atoi(value)
			if err != nil {
				return 0, nil, err
			}
			accumulator += acc
			i++
		case "jmp":
			jmp, _ := strconv.Atoi(value)
			if err != nil {
				return 0, nil, err
			}
			i += jmp
		}
	}
	return
}
