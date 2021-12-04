package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	input := getInput("day3/input.txt")
	gamma, epsilon, err := getMostCommon(input)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Pt1 Position Multiple %d", gamma*epsilon)
}

func getInput(filename string) [][]byte {
	var input [][]byte
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func getMostCommon(input [][]byte) (int, int, error) {
	colSum := make([]int, len(input[0]))
	log.Print(colSum)
	for _, bslice := range input {
		log.Print(bslice)
		for i, n := range bslice {
			if n == '1' {
				colSum[i] = colSum[i] + 1
			}
		}
	}
	mostCommon := make([]rune, len(input[0]))
	leastCommon := make([]rune, len(input[0]))
	for i, n := range colSum {
		if n > len(input)/2 {
			mostCommon[i] = '1'
			leastCommon[i] = '0'
		} else {
			mostCommon[i] = '0'
			leastCommon[i] = '1'
		}

	}
	log.Printf("colSum: %d", colSum)
	log.Printf("Gamma: %s", string(mostCommon))
	log.Printf("Epsilon: %s", string(leastCommon))
	gamma, err := strconv.ParseInt(string(mostCommon), 2, 16)
	epsilon, err := strconv.ParseInt(string(leastCommon), 2, 16)
	log.Printf("Gamma: %d", gamma)
	log.Printf("Epsilon: %d", epsilon)

	return int(gamma), int(epsilon), err
}
