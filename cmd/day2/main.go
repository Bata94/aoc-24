package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() [][]int {
	file, err := os.Open("inputs/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ls := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inpLine := strings.Split(scanner.Text(), " ")
		innerLs := []int{}

		for _, x := range inpLine {
			inp, err := strconv.Atoi(x)
			if err != nil {
				log.Fatal(err)
			}
			innerLs = append(innerLs, inp)
		}
		ls = append(ls, innerLs)
	}

	log.Println("Length of Input: ", len(ls))

	return ls
}

func ruleCheck(r []int) bool {
	incr := false

	if r[0] == r[1] {
		return false
	} else if r[0] < r[1] {
		incr = true
	}

	for i := 0; i < len(r)-1; i++ {
		if r[i] == r[i+1] || (r[i] > r[i+1] && incr) || (r[i] < r[i+1] && !incr) {
			return false
		}
	}

	for i := 0; i < len(r)-1; i++ {
		diff := r[i] - r[i+1]
		if diff < 0 {
			diff = -diff
		}

		if diff > 3 {
			return false
		}
	}

	return true
}

func calcNumSafe(data [][]int, tolerateSingleBadLevel bool) int {
	safeReports := 0

	for _, r := range data {
		if ruleCheck(r) {
			safeReports++
		} else if tolerateSingleBadLevel {
			for i := 0; i < len(r); i++ {
				newR := make([]int, 0, len(r)-1)
				newR = append(newR, r[:i]...)
				newR = append(newR, r[i+1:]...)
				if ruleCheck(newR) {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports
}

func main() {
	lsReports := readInput()

	log.Println("Advent of Code - Day 02:")
	log.Println("Answer Part 1 (Num safe reports): ", calcNumSafe(lsReports, false))
	log.Println("Answer Part 2 (Num safe reports, tolerate a single bad level): ", calcNumSafe(lsReports, true))
}
