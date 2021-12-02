package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	vectors := getInput("day2/input.txt")
	pos := Position{0, 0, 0}
	pos = ApplyVectors(pos, vectors)
	log.Printf("Pt1 Position Multiple %d", pos.x*pos.y)

	pos = Position{0, 0, 0}
	pos = ApplyAimVectors(pos, vectors)
	log.Printf("Pt2 Position Multiple %d", pos.x*pos.y)
}

type Vector struct {
	direction string
	magnitude int
}

type Position struct {
	x   int
	y   int
	aim int
}

func getInput(filename string) []Vector {
	var vectors []Vector
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitStr := strings.SplitN(scanner.Text(), " ", 2)
		n, err := strconv.Atoi(splitStr[1])
		if err != nil {
			log.Fatal(err)
		}
		vectors = append(vectors, Vector{splitStr[0], n})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return vectors
}

func ApplyVectors(position Position, vectors []Vector) Position {
	for _, v := range vectors {
		if v.direction == "forward" {
			position.x += v.magnitude
		} else if v.direction == "up" {
			position.y -= v.magnitude
		} else if v.direction == "down" {
			position.y += v.magnitude
		} else {
			log.Fatalf("Invalid direction: %s", v.direction)
		}
	}
	return position
}

func ApplyAimVectors(position Position, vectors []Vector) Position {
	for _, v := range vectors {
		if v.direction == "forward" {
			position.x += v.magnitude
			position.y += position.aim * v.magnitude
		} else if v.direction == "up" {
			position.aim -= v.magnitude
		} else if v.direction == "down" {
			position.aim += v.magnitude
		} else {
			log.Fatalf("Invalid direction: %s", v.direction)
		}
	}
	return position
}
