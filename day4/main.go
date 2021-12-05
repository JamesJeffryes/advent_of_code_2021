package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	calls, boards := getInput("day4/input.txt")
	log.Println(calls)
	log.Println(boards[0])
	log.Println(playBingo(calls, boards, true))
	log.Println(playBingo(calls, boards, false))

}

type Board [][]string

func getInput(filename string) ([]string, []Board) {
	boardSize := 5
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	calls := strings.Split(scanner.Text(), ",")
	scanner.Scan() // skip blank line
	i := 0
	j := 0
	boards := make([]Board, 1)
	boards[0] = make(Board, boardSize)
	for scanner.Scan() {
		if j == boardSize {
			i++
			boards = append(boards, make(Board, boardSize))
			j = 0
		} else {
			boards[i][j] = strings.Fields(scanner.Text())
			j++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return calls, boards
}

func playBingo(calls []string, boards []Board, returnFirst bool) int {
	called := make(map[string]struct{})
	lastScore := 0
	invalidBoards := make(map[int]struct{}, len(boards))
	for i, c := range calls {
		called[c] = struct{}{}
		if i < 4 {
			continue
		} //can't win with less than 4 so why try?
		for bnum, b := range boards {
			if _, ok := invalidBoards[bnum]; ok {
				continue
			}
			columnWin := map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true}
			for _, row := range b {
				rowWin := 0
				for k, val := range row {
					if _, ok := called[val]; ok {
						rowWin++
					} else {
						columnWin[k] = false
					}
				}
				if rowWin == 5 {
					lastScore = scoreBoard(called, b, c)
					if returnFirst {
						return lastScore
					} else {
						invalidBoards[bnum] = struct{}{}
					}
				}
			}
			//check if column win
			for _, v := range columnWin {
				if v {
					lastScore = scoreBoard(called, b, c)
					if returnFirst {
						return lastScore
					} else {
						invalidBoards[bnum] = struct{}{}
					}
				}
			}
		}
	}
	return lastScore
}

func scoreBoard(called map[string]struct{}, b Board, c string) int {
	cnum, err := strconv.Atoi(c)
	if err != nil {
		log.Fatal(err)
	}
	unmarked := 0
	for _, row := range b {
		for _, val := range row {
			if _, ok := called[val]; !ok {
				num, err := strconv.Atoi(val)
				if err != nil {
					log.Fatal(err)
				}
				unmarked += num
			}
		}
	}
	return cnum * unmarked
}
