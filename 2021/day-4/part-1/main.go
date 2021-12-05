package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	log.Printf("File: %s", file.Name())

	return file
}

func atoi(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return number
}

type LineIterator struct {
	Lines []string
	Index int
}

func createLineIterator(path string) *LineIterator {
	file := openFile(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	li := &LineIterator{}

	for scanner.Scan() {
		li.Lines = append(li.Lines, scanner.Text())
	}

	return li
}

func (li *LineIterator) Next() (string, bool) {
	if len(li.Lines) <= li.Index {
		return "", false
	}

	line := li.Lines[li.Index]
	log.Printf(line)
	li.Index++

	return line, true
}

func parseNumbers(li *LineIterator) []int {
	line, ok := li.Next()
	if !ok {
		log.Fatal("no numbers line found")
	}
	numStrings := strings.Split(line, ",")
	nums := []int{}

	for _, s := range numStrings {
		nums = append(nums, atoi(s))
	}

	return nums
}

type boardSquare struct {
	value  int
	marked bool
}

type boardLine struct {
	squares []boardSquare
}

type board struct {
	lines []boardLine
}

func parseBoardLine(li *LineIterator) boardLine {
	line, ok := li.Next()
	if !ok {
		log.Fatal("board line missing")
	}

	numbers := strings.Fields(line)

	if len(numbers) != 5 {
		log.Fatal("board line didn't have 5 numbers")
	}

	bLine := boardLine{
		squares: []boardSquare{
			{value: atoi(numbers[0])},
			{value: atoi(numbers[1])},
			{value: atoi(numbers[2])},
			{value: atoi(numbers[3])},
			{value: atoi(numbers[4])},
		},
	}

	return bLine
}

func parseBoard(li *LineIterator) board {
	return board{
		lines: []boardLine{
			parseBoardLine(li),
			parseBoardLine(li),
			parseBoardLine(li),
			parseBoardLine(li),
			parseBoardLine(li),
		},
	}
}

func parseBoards(li *LineIterator) []board {
	line, ok := li.Next()
	if !ok {
		log.Fatal("No lines after numbers line")
	}

	if line != "" {
		log.Fatal("Spacer line missing between numbers line and first board")
	}

	boards := []board{}
	boards = append(boards, parseBoard(li))

	line, ok = li.Next()
	for ok {
		if line != "" {
			log.Fatal("Spacer line missing between boards")
		}

		boards = append(boards, parseBoard(li))
		line, ok = li.Next()
	}

	return boards
}

func markNumber(n int, b *board) {
	for _, l := range b.lines {
		for i := range l.squares {
			if l.squares[i].value == n {
                l.squares[i].marked = true
                log.Printf("Marking %d", n)
				return
			}
		}
	}
	return
}

func getRows(b *board) []boardLine {
	return b.lines
}

func checkAll(squares []boardSquare) bool {
	for _, s := range squares {
		if !s.marked {
			return false
		}
	}
	return true
}

func getColumns(b *board) []boardLine {
	columns := []boardLine{}
	for i := 0; i < len(b.lines[0].squares); i++ {
		c := boardLine{}
		for j := 0; j < len(b.lines); j++ {
			c.squares = append(c.squares, b.lines[j].squares[i])
		}
		columns = append(columns, c)
	}
	return columns
}

func checkBoard(b *board) bool {
	for _, r := range getRows(b) {
		if checkAll(r.squares) {
			return true
		}
	}
	for _, c := range getColumns(b) {
		if checkAll(c.squares) {
			return true
		}
	}
	return false
}

func sumUnmarked(b *board) int {
	total := 0

	for _, l := range b.lines {
		for _, s := range l.squares {
			if !s.marked {
				total += s.value
			}
		}
	}
	return total
}

func logDrawnNumbers(numbers []int, index int) {
    str := ""
    for i, n := range numbers {
        if i<= index {
            str += color.GreenString(fmt.Sprintf("%d ", n))
        } else {
            str += fmt.Sprintf("%d ", n)
        }
    }
    log.Print(str)
}

func logBoard(b *board) {
    str := ""
    for _, l := range b.lines {
        str += "\n"
        for _, s := range l.squares {
            if s.marked {
                str += color.GreenString(fmt.Sprintf("%d ", s.value))
            } else {
                str += fmt.Sprintf("%d ", s.value)
            }
        }
    }
    str += "\n"
    log.Print(str)
}


func main() {
	lineIterator := createLineIterator(os.Args[1])

	numbers := parseNumbers(lineIterator)
	boards := parseBoards(lineIterator)

	for i, n := range numbers {
        logDrawnNumbers(numbers, i)
		for _, b := range boards {
			markNumber(n, &b)
            logBoard(&b)
			won := checkBoard(&b)
			if won {
				sum := sumUnmarked(&b)
				log.Printf("Final Number: %d", n)
				log.Printf("Sum: %d", sum)
				log.Printf("Score: %d", sum*n)
				return
			}
		}
	}

	log.Fatal("Got to the end without any board winning")
}
