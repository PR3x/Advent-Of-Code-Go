package main

import (
	"log"
	"strconv"
)

func checkDupes(freqs []int) (int, bool) {
	for i, j := range freqs {
		for _, k := range freqs[0:i] {
			if j == k {
				return k, true
			}
		}
	}
	return 0, false
}

func buildFreqs(commands []int, startFrom int) []int {
	freqs := make([]int, 0)
	freq := startFrom
	for _, i := range commands {
		freq += i
		freqs = append(freqs, freq)
	}
	return freqs
}

// Day1 runs the problem for Day 1
func Day1(ch chan int) {
	// Part 1
	lines := ReadLines("2018day1.txt")
	commands := make([]int, 0)
	for _, s := range lines {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		commands = append(commands, i)
	}
	freqs := append([]int{0}, buildFreqs(commands, 0)...)
	ch <- freqs[len(freqs)-1]

	dupe := 0
	done := false
	for !done {
		dupe, done = checkDupes(freqs)
		freqs = append(freqs, buildFreqs(commands, freqs[len(freqs)-1])...)
	}
	ch <- dupe

	close(ch)
}
