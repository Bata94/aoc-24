package main

import (
	"bufio"
	"log"
	"os"
)

func readInput() [][]rune {
	file, err := os.Open("inputs/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ret := [][]rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		innerLs := []rune{}
		for _, r := range scanner.Text() {
			innerLs = append(innerLs, r)
		}
		ret = append(ret, innerLs)
	}

	return ret
}

func checkNeighbors(data [][]rune, r rune, x, y int) [][]int {
	ret := [][]int{}

	for _, neighbor := range listNeighbors(data, x, y) {
		if data[neighbor[1]][neighbor[0]] == r {
			ret = append(ret, neighbor)
		}
	}

	return ret
}

func listNeighbors(data [][]rune, x, y int) [][]int {
	ret := [][]int{}

	for yOffset := -1; yOffset <= 1; yOffset++ {
		neighborY := y + yOffset
		if neighborY < 0 || neighborY >= len(data) {
			continue
		}

		for xOffset := -1; xOffset <= 1; xOffset++ {
			neighborX := x + xOffset
			if neighborX < 0 || neighborX >= len(data[0]) {
				continue
			}

			ret = append(ret, []int{neighborX, neighborY})
		}
	}

	return ret
}

func main() {
	data := readInput()
	part1 := 0

	for y, line := range data {
		for x, r := range line {
			if r == 'X' {
				neighborM := checkNeighbors(data, 'M', x, y)
				for _, n := range neighborM {
					offsets := []int{x - n[0], y - n[1]}
					if (y+3*offsets[1] < 0 || y+3*offsets[1] >= len(data)) || (x+3*offsets[0] < 0 || x+3*offsets[0] >= len(data[0])) {
						continue
					}

					if data[y+2*offsets[1]][x+2*offsets[0]] == 'A' && data[y+3*offsets[1]][x+3*offsets[0]] == 'S' {
						part1++
					}
				}
			}
		}
	}

	log.Println("Advent of Code - Day 04:")
	log.Println("Answer Part 1 (Word search XMAS): ", part1)
	log.Println("Answer Part 2 (): ", nil)
}
