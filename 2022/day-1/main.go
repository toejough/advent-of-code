package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// read input
	day := os.Args[1]
	part := os.Args[2]

	filename := fmt.Sprintf("%s-input-puzzle.txt", day)
	text := mustReadFileText(filename)

	// solve
	var (
		answer string
		err    error
	)

	switch day {
	case "day1":
		switch part {
		case "part1":
			answer, err = solveDay1Part1(text)
			if err != nil {
				panic(err)
			}
		case "part2":
			answer, err = solveDay1Part2(text)
			if err != nil {
				panic(err)
			}
		default:
			log.Fatalf("No solver for %s %s", day, part)
		}
	default:
		log.Fatalf("No solvers for %s ", day)
	}

	// final
	log.Printf("Answer: %s", answer)
}

func mustReadFileText(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := string(bytes)

	return text
}

func splitNoEmpty(s string, sep string) []string {
	list := strings.Split(s, sep)
	noEmpty := []string{}

	for _, text := range list {
		if len(text) == 0 {
			continue
		}

		noEmpty = append(noEmpty, text)
	}

	return noEmpty
}

func solveDay1Part1(text string) (string, error) {
	// split into lists of calories
	hunks := splitNoEmpty(text, "\n\n")
	stringLists := splitHunks(hunks)

	lists, err := convListsOfStringsToListsOfInts(stringLists)
	if err != nil {
		return "", err
	}

	// sum the lists
	sums := sumLists(lists)

	// max
	max, err := max(sums)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", max), nil
}

func solveDay1Part2(text string) (string, error) {
	// split into lists of calories
	hunks := splitNoEmpty(text, "\n\n")
	stringLists := splitHunks(hunks)

	lists, err := convListsOfStringsToListsOfInts(stringLists)
	if err != nil {
		return "", err
	}

	// sum the lists
	sums := sumLists(lists)

	// sort
	sort.Ints(sums)

	// top 3
	top3 := sums[len(sums)-3:]

	// sum
	total := sum(top3)

	return fmt.Sprintf("%d", total), nil
}

var (
	ErrNoRPSEnumForRune       = fmt.Errorf("no RPS enum was found for the rune")
	ErrUnrecognizedRPSEnum    = fmt.Errorf("unrecognized RPS enum")
	ErrUnrecognizedRPSOutcome = fmt.Errorf("unrecognized RPS outcome enum")
)

func solveDay2Part1(text string) (string, error) {
	// split into lines
	lines := splitNoEmpty(text, "\n")
	// parse lines into opponent/you
	type EncodedRPS struct {
		Them rune
		You  rune
	}
	encodedStrategy := []EncodedRPS{}
	for _, line := range lines {
		encodedStrategy = append(encodedStrategy, EncodedRPS{Them: rune(line[0]), You: rune(line[2])})
	}
	// convert into nicer enum representation
	type RPSEnum int
	const (
		Rock RPSEnum = iota
		Paper
		Scissors
	)
	type RPS struct {
		Them RPSEnum
		You  RPSEnum
	}
	strategy := []RPS{}
	for _, encoded := range encodedStrategy {
		rps := RPS{}
		switch encoded.Them {
		case 'A':
			rps.Them = Rock
		case 'B':
			rps.Them = Paper
		case 'C':
			rps.Them = Scissors
		default:
			return "", fmt.Errorf("unable to match %v to an RPS selection for them: %w", encoded.Them, ErrNoRPSEnumForRune)
		}
		switch encoded.You {
		case 'X':
			rps.You = Rock
		case 'Y':
			rps.You = Paper
		case 'Z':
			rps.You = Scissors
		default:
			return "", fmt.Errorf("unable to match %v to an RPS selection for you: %w", encoded.You, ErrNoRPSEnumForRune)
		}
		strategy = append(strategy, rps)
	}
	// enhance with outcome of encounters
	type RPSOutcome int
	const (
		Lost RPSOutcome = iota
		Tied
		Won
	)
	type EvaluatedRPS struct {
		RPS     RPS
		Outcome RPSOutcome
	}
	evaluatedStrategy := []EvaluatedRPS{}
	for _, rps := range strategy {
		evaluated := EvaluatedRPS{RPS: rps}
		switch rps.Them {
		case Rock:
			switch rps.You {
			case Rock:
				evaluated.Outcome = Tied
			case Paper:
				evaluated.Outcome = Won
			case Scissors:
				evaluated.Outcome = Lost
			default:
				return "", fmt.Errorf("unable to evaluate the outcome of a match when you chose %v: %w", rps.You, ErrUnrecognizedRPSEnum)
			}
		case Paper:
			switch rps.You {
			case Rock:
				evaluated.Outcome = Lost
			case Paper:
				evaluated.Outcome = Tied
			case Scissors:
				evaluated.Outcome = Won
			default:
				return "", fmt.Errorf("unable to evaluate the outcome of a match when you chose %v: %w", rps.You, ErrUnrecognizedRPSEnum)
			}
		case Scissors:
			switch rps.You {
			case Rock:
				evaluated.Outcome = Won
			case Paper:
				evaluated.Outcome = Lost
			case Scissors:
				evaluated.Outcome = Tied
			default:
				return "", fmt.Errorf("unable to evaluate the outcome of a match when you chose %v: %w", rps.You, ErrUnrecognizedRPSEnum)
			}
		default:
			return "", fmt.Errorf("unable to evaluate the outcome of a match when they chose %v: %w", rps.Them, ErrUnrecognizedRPSEnum)
		}
		evaluatedStrategy = append(evaluatedStrategy, evaluated)
	}
	// enhance with scores
	type ScoredRPS struct {
		Evaluated EvaluatedRPS
		Score     int
	}
	scoredStrategy := []ScoredRPS{}
	for _, evaluated := range evaluatedStrategy {
		score := 0
		// score selection
		switch evaluated.RPS.You {
		case Rock:
			score += 1
		case Paper:
			score += 2
		case Scissors:
			score += 3
		default:
			return "", fmt.Errorf("unable to score the outcome of a match when you chose %v: %w", evaluated.RPS.You, ErrUnrecognizedRPSEnum)
		}
		// score outcome
		switch evaluated.Outcome {
		case Lost:
			score += 0
		case Tied:
			score += 3
		case Won:
			score += 6
		default:
			return "", fmt.Errorf("unable to score the outcome of a match when the outcome was %v: %w", evaluated.Outcome, ErrUnrecognizedRPSOutcome)
		}
		scoredStrategy = append(scoredStrategy, ScoredRPS{Evaluated: evaluated, Score: score})
	}
	// reduce to scores
	scores := []int{}
	for _, scored := range scoredStrategy {
		scores = append(scores, scored.Score)
	}
	// sum them
	total := sum(scores)

	return fmt.Sprintf("%d", total), nil
}

var ErrNoMaxPossible = fmt.Errorf("no max possible: input list was empty")

func max(list []int) (int, error) {
	if len(list) == 0 {
		return 0, ErrNoMaxPossible
	}

	max := list[0]
	for _, value := range list[1:] {
		if value > max {
			max = value
		}
	}

	return max, nil
}

func sumLists(lists [][]int) []int {
	sums := []int{}

	for _, list := range lists {
		sums = append(sums, sum(list))
	}

	return sums
}

func sum(list []int) int {
	sum := 0
	for _, value := range list {
		sum += value
	}

	return sum
}

func splitHunks(hunks []string) [][]string {
	lists := [][]string{}

	for _, hunk := range hunks {
		list := splitNoEmpty(hunk, "\n")
		lists = append(lists, list)
	}

	return lists
}

func convListsOfStringsToListsOfInts(stringLists [][]string) ([][]int, error) {
	lists := [][]int{}

	for _, stringList := range stringLists {
		list, err := convStringsToInts(stringList)
		if err != nil {
			return nil, fmt.Errorf("while converting string list '%v' to int: %w", list, err)
		}

		lists = append(lists, list)
	}

	return lists, nil
}

func convStringsToInts(stringList []string) ([]int, error) {
	list := []int{}

	for _, s := range stringList {
		value, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("while converting string '%s' to int: %w", s, err)
		}

		list = append(list, value)
	}

	return list, nil
}
