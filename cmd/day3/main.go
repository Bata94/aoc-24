package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readInput() string {
	file, err := os.Open("inputs/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	str := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str += scanner.Text()
	}

	return str
}

func readMemory(data string, allInstructions bool) int {
	sum := 0
	nextInstruction := true

	for i := 0; i < len(data); i++ {
		if !allInstructions && (data[i] == 'd' && data[i+1] == 'o') {
			strDo := ""

			for j := i; j < i+7; j++ {
				strDo += string(data[j])
				if data[j] == ')' {
					i = j
					break
				}
			}

			if strDo == "do()" {
				nextInstruction = true
			} else if strDo == "don't()" {
				nextInstruction = false
			}
		}

		if nextInstruction && (data[i] == 'm' && data[i+1] == 'u' && data[i+2] == 'l' && data[i+3] == '(') {
			strMul := ""

			for j := i; j < i+12; j++ {
				strMul += string(data[j])
			}

			num1Sel := true
			num1Str := ""
			num2Str := ""

			for j := i + 4; j < i+12; j++ {
				if data[j] == ')' {
					break
				} else if data[j] == ',' {
					num1Sel = false
					continue
				}

				if num1Sel {
					num1Str += string(data[j])
				} else {
					num2Str += string(data[j])
				}

				if data[j] == ')' {
					i = j
					break
				}
			}

			num1, err1 := strconv.Atoi(num1Str)
			num2, err2 := strconv.Atoi(num2Str)
			if err1 != nil || err2 != nil {
				continue
			}

			sum += num1 * num2
		}
	}
	return sum
}

func main() {
	data := readInput()

	log.Println("Advent of Code - Day 03:")
	log.Println("Answer Part 1 (Memory multiplication): ", readMemory(data, true))
	log.Println("Answer Part 2 (Memory multiplication, only do prefixed): ", readMemory(data, false))
}
