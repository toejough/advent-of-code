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
		if i <= index {
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

type point struct {
	x, y int
}

type segment struct {
	first, second point
}

func parsePoint(s string) point {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		log.Fatal("length of point parts was not 2")
	}

	return point{
		x: atoi(parts[0]),
		y: atoi(parts[1]),
	}
}

func parseSegment(line string) segment {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		log.Fatal("length of line segment parts was not 3")
	}

	return segment{
		first:  parsePoint(parts[0]),
		second: parsePoint(parts[2]),
	}
}

func parseSegments(li *LineIterator) []segment {
	line, ok := li.Next()
	segments := []segment{}
	for ok {
		segments = append(segments, parseSegment(line))
		line, ok = li.Next()
	}
	return segments
}

type ventMap struct {
	coordinateCounts map[int]map[int]int
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func expandPoints(s segment) []point {
    log.Printf("segment: %v", s)
	points := []point{}

	xdist := abs(s.first.x - s.second.x)
	ydist := abs(s.first.y - s.second.y)

	numPoints := max(xdist, ydist) + 1

	var xmult, ymult int

	if xdist != 0 {
		xmult = (s.second.x - s.first.x) / xdist
	}
	if ydist != 0 {
		ymult = (s.second.y - s.first.y) / ydist
	}

	for i := 0; i < numPoints; i++ {
		points = append(points, point{x: s.first.x + xmult*i, y: s.first.y + ymult*i})
	}
    log.Printf("points: %v", points)
    log.Println()
	return points
}

func applySegmentsToMap(m *ventMap, segments []segment) {
	for _, s := range segments {
		points := expandPoints(s)
		for _, p := range points {
			_, ok := m.coordinateCounts[p.x]
			if !ok {
				m.coordinateCounts[p.x] = map[int]int{}
			}
			m.coordinateCounts[p.x][p.y]++
		}
	}
}

func countCoordinatesOver2(v *ventMap) (count int) {
	for _, y := range v.coordinateCounts {
		for _, v := range y {
			if v > 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	lineIterator := createLineIterator(os.Args[1])

	segments := parseSegments(lineIterator)
	vMap := &ventMap{coordinateCounts: map[int]map[int]int{}}

	applySegmentsToMap(vMap, segments)

	count := countCoordinatesOver2(vMap)

	log.Printf("Count: %d", count)
}
