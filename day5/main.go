package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getInput("day5/input.txt")
	overlaps := countOverlaps(lines)
	log.Println(overlaps)

}

type Point struct {
	x int
	y int
}

func getInput(filename string) [][2]Point {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make([][2]Point, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pt := [2]Point{}
		ptTxts := strings.Split(scanner.Text(), " -> ")
		for i, txt := range ptTxts {
			spTxt := strings.Split(txt, ",")
			x, err := strconv.Atoi(spTxt[0])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(spTxt[1])
			if err != nil {
				log.Fatal(err)
			}
			pt[i] = Point{x, y}
		}
		lines = append(lines, pt)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func countOverlaps(lines [][2]Point) int {
	overlaps := 0
	lineMap := make(map[string]int)
	for _, line := range lines {
		xStep := getStep(line[0].x, line[1].x)
		yStep := getStep(line[0].y, line[1].y)
		stop := getLen(line[0].x, line[1].x)
		if xStep == 0 {
			stop = getLen(line[0].y, line[1].y)
		}
		for i := 0; i <= stop; i++ {
			x := line[0].x + xStep*i
			y := line[0].y + yStep*i
			ptStr := fmt.Sprintf("%d,%d", x, y)
			if q := lineMap[ptStr]; q == 1 {
				overlaps++
			}
			lineMap[ptStr]++
		}
	}
	return overlaps
}

func getStep(a, b int) int {
	if a > b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}

func getLen(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
