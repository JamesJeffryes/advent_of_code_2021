package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	readings := getReadings("day1/input.txt")
	log.Printf("Detected %d increases", countIncreases(readings))
	log.Printf("Detected %d increases with window of 3", countWindowIncreases(readings, 3))
}

func getReadings(filename string) []int {
	readings := []int{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		readings = append(readings, n)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return readings
}

func countIncreases(readings []int) int {
	increases := 0
	prev := readings[0]
	for _, n := range readings {
		if n > prev {
			increases++
		}
		prev = n
	}
	return increases
}

func countWindowIncreases(readings []int, window int) int {
	increases := 0
	if window > len(readings) {
		return 0
	}
	prev := sumSlice(readings[0:window])
	for i := range readings {
		if  i + window > len(readings) {
			break
		}
		n := sumSlice(readings[i:i+window])
		if n > prev {
			increases++
		}
		prev = n
	}
	return increases
}

func sumSlice(s []int) int {
	sum := 0
	for _, n := range s {
		sum = sum + n
	}
	return sum
}