package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() ([]int, []int) {
	inp, err := os.Open("inputs/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inp.Close()

	ls1 := []int{}
	ls2 := []int{}

	scanner := bufio.NewScanner(inp)
	for scanner.Scan() {
		inpLine := strings.Split(scanner.Text(), "   ")

		inp1, err := strconv.Atoi(inpLine[0])
		if err != nil {
			log.Fatal(err)
		}
		inp2, err := strconv.Atoi(inpLine[1])
		if err != nil {
			log.Fatal(err)
		}

		ls1 = append(ls1, inp1)
		ls2 = append(ls2, inp2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Length of Inputs: ", len(ls1), ", ", len(ls2))

	return ls1, ls2
}

func sortLS(ls []int) []int {
	sort.Slice(ls, func(i, j int) bool {
		return ls[i] < ls[j]
	})
	return ls
}

func calcDistance(ls1, ls2 []int) int {
	totalDistance := 0
	for i := 0; i < len(ls1); i++ {
		distance := 0
		if ls1[i] > ls2[i] {
			distance = ls1[i] - ls2[i]
		} else {
			distance = ls2[i] - ls1[i]
		}
		totalDistance += distance
	}
	return totalDistance
}

func calcSimilarityScore(ls1, ls2 []int) int {
	similarityScore := 0

	for i := 0; i < len(ls1); i++ {
		similar := 0
		for j := 0; j < len(ls2); j++ {
			if ls1[i] == ls2[j] {
				similar++
			} else if ls1[i] < ls2[j] {
				break
			}
		}
		similarityScore += ls1[i] * similar
	}

	return similarityScore
}

func main() {
	ls1, ls2 := readInput()

	ls1 = sortLS(ls1)
	ls2 = sortLS(ls2)

	log.Println("Advent of Code - Day 01:")
	log.Println("Answer Part 1 (total distance):", calcDistance(ls1, ls2))
	log.Println("Answer Part 2 (similarity Score):", calcSimilarityScore(ls1, ls2))
}
