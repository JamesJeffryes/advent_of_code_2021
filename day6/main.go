package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fishCounts := getInput("day6/input.txt")
	log.Println(fishCounts)
	log.Println(breedFish(fishCounts, 256))
}

func getInput(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fishCounts := make([]int, 9)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fishSlice := strings.Split(scanner.Text(), ",")
	for _, fish := range fishSlice {
		n, err := strconv.Atoi(fish)
		if err != nil {
			log.Fatal(err)
		}
		fishCounts[n]++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fishCounts
}

func breedFish(fishCounts []int, generations int) int {
	newFish := 0
	for g := 0; g < generations; g++ {
		for i := 8; i >= 0; i-- {
			if i == 0 {
				fishCounts[6] += fishCounts[i]
				fishCounts[8] = fishCounts[i]
			}
			currentCount := fishCounts[i]
			fishCounts[i] = newFish
			newFish = currentCount
		}
		log.Println(fishCounts)
	}
	total := 0
	for _, n := range fishCounts {
		total += n
	}
	return total
}
