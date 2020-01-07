package main

import (
	"log"
	"strconv"
)

func checkDupes(freqs map[int]bool) (int, bool) {
	for i := range freqs {
		if freqs[i] {
			return i, true
		}
	}
	return 0, false
}

func buildFreqs(commands []int, freqs map[int]bool, startFrom int) (map[int]bool, int) {
	freq := startFrom
	for _, i := range commands {
		freq += i
		freqs[freq] = true
	}
	return freqs, freq
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
	freqs := map[int]bool{0: true}
	freq := 0
	freqs, freq = buildFreqs(commands, freqs, 0)
	ch <- freq

	dupe := 0
	done := false
	last := 0
	for !done {
		freqs, last = buildFreqs(commands, freqs, last)
		dupe, done = checkDupes(freqs)
	}
	ch <- dupe

	close(ch)
}
