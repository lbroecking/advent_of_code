package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input_day_5.txt")

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

	max := 0
	seats := make(map[int]bool)
	// F "Front": lower half 0-63
	// B "Back": upper half 63-127
	// R "Right": upper half 4-7
	// L "Left": lower half 0-4
	for _, line := range txtlines {
		row, col := getRowAndCole(line)
		sid := row*8 + col
		seats[sid] = true
		if sid > max {
			max = sid
		}
	}
	fmt.Println("Part 1: ", max)

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			seatId := i*8 + j
			if !seats[seatId] {
				if seats[seatId+1] && seats[seatId-1] {
					fmt.Println("Part 2: ", seatId)
				}
			}
		}
	}

}

func getRowAndCole(s string) (int, int) {
	rows := make([]int, 128)
	for i := range rows {
		rows[i] += i
	}

	cols := make([]int, 8)
	for i := range cols {
		cols[i] += i
	}

	for _, r := range s {
		c := string(r)
		switch c {
		case "F":
			rows = rows[:len(rows)/2]
		case "B":
			rows = rows[len(rows)/2:]
		case "L":
			cols = cols[:len(cols)/2]
		case "R":
			cols = cols[len(cols)/2:]
		}
	}
	return rows[0], cols[0]
}
