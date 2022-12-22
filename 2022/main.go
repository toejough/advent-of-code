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
		answer, err = solveDay1(part, text)
	case "day2":
		switch part {
		case "part1":
			answer, err = solveDay2Part1(text)
			if err != nil {
				panic(err)
			}
		case "part2":
			answer, err = solveDay2Part2(text)
			if err != nil {
				panic(err)
			}
		default:
			log.Fatalf("No solver for %s %s", day, part)
		}
	default:
		log.Fatalf("No solvers for %s ", day)
	}

	if err != nil {
		panic(err)
	}

	// final
	log.Printf("Answer: %s", answer)
}

var ErrMissingSolver = fmt.Errorf("missing solver")

func solveDay1(part string, text string) (string, error) {
	switch part {
	case "part1":
		return solveDay1Part1(text)
	case "part2":
		return solveDay1Part2(text)
	default:
		return "", fmt.Errorf("no solver for day1 %s: %w", part, ErrMissingSolver)
	}
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

var ErrNotEnoughItems = fmt.Errorf("not enough items")

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
	const numbersToSum = 3
	if len(sums) < numbersToSum {
		return "", fmt.Errorf("can't take the top three of a list of only %d: %w", len(sums), ErrNotEnoughItems)
	}

	top3 := sums[len(sums)-numbersToSum:]

	// sum
	total := sum(top3)

	return fmt.Sprintf("%d", total), nil
}

var (
	ErrNoRPSEnumForRune        = fmt.Errorf("no RPS enum was found for the rune")
	ErrUnrecognizedRPSEnum     = fmt.Errorf("unrecognized RPS enum")
	ErrUnrecognizedRPSOutcome  = fmt.Errorf("unrecognized RPS outcome enum")
	ErrNoRPSOutcomeEnumForRune = fmt.Errorf("no RPS outcome enum was found for the rune")
)

type EncodedRPS struct {
	Them rune
	You  rune
}

type RPSEnum int

const (
	Rock RPSEnum = iota
	Paper
	Scissors
)

type RPSMatch struct {
	Them RPSEnum
	You  RPSEnum
}

type RPSOutcomeMatch struct {
	Them    RPSEnum
	Outcome RPSOutcome
}

type RPSOutcome int

const (
	Lost RPSOutcome = iota
	Tied
	Won
)

type DecidedRPSMatch struct {
	Them    RPSEnum
	You     RPSEnum
	Outcome RPSOutcome
}

type ScoredRPS struct {
	Evaluated DecidedRPSMatch
	Score     int
}

func solveDay2Part1(text string) (string, error) {
	// split into lines
	lines := splitNoEmpty(text, "\n")
	// parse lines into opponent/you
	encodedMatches, err := parseEncodedMatchLines(lines)
	if err != nil {
		return "", fmt.Errorf("unable to solve after match parsing failure: %w", err)
	}
	// convert into nicer enum representation
	matches, err := decodeMatches(encodedMatches)
	if err != nil {
		return "", fmt.Errorf("unable to solve after match decoding failure: %w", err)
	}
	// enhance with outcome of encounters
	decidedMatches, err := decideMatches(matches)
	if err != nil {
		return "", fmt.Errorf("unable to solve after match decision failure: %w", err)
	}
	// enhance with scores
	scoredStrategy, err := scoreRPSMatches(decidedMatches)
	if err != nil {
		return "", fmt.Errorf("unable to solve after scoring failure: %w", err)
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

func solveDay2Part2(text string) (string, error) {
	// split into lines
	lines := splitNoEmpty(text, "\n")
	// parse lines into opponent/you
	encodedMatches, err := parseEncodedMatchLines(lines)
	if err != nil {
		return "", fmt.Errorf("unable to solve after match parsing failure: %w", err)
	}
	// convert into nicer enum representation
	decidedMatches, err := decodeMatchesAsOutcomes(encodedMatches)
	if err != nil {
		return "", fmt.Errorf("unable to solve after match decoding failure: %w", err)
	}
	// enhance with roll to reach outcome
	deducedMoves, err := deduceMoves(decidedMatches)
	if err != nil {
		return "", fmt.Errorf("unable to solve after match decision failure: %w", err)
	}
	// enhance with scores
	scoredStrategy, err := scoreRPSMatches(deducedMoves)
	if err != nil {
		return "", fmt.Errorf("unable to solve after scoring failure: %w", err)
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

func deduceMoves(outcomeMatches []RPSOutcomeMatch) ([]DecidedRPSMatch, error) {
	outcomeMap := map[RPSEnum]map[RPSOutcome]RPSEnum{
		Rock: {
			Lost: Scissors,
			Tied: Rock,
			Won:  Paper,
		},
		Paper: {
			Lost: Rock,
			Tied: Paper,
			Won:  Scissors,
		},
		Scissors: {
			Lost: Paper,
			Tied: Scissors,
			Won:  Rock,
		},
	}
	decidedMatches := []DecidedRPSMatch{}

	for _, match := range outcomeMatches {
		outcomeAndMe, themOk := outcomeMap[match.Them]
		if !themOk {
			return nil, fmt.Errorf("unable to deduce moves with %v: %w", match.Them, ErrUnrecognizedRPSEnum)
		}

		you, outcomeOK := outcomeAndMe[match.Outcome]
		if !outcomeOK {
			return nil, fmt.Errorf("unable to deduce moves with %v: %w", match.Outcome, ErrNoRPSOutcomeEnumForRune)
		}

		decidedMatches = append(
			decidedMatches,
			DecidedRPSMatch{Them: match.Them, You: you, Outcome: match.Outcome},
		)
	}

	return decidedMatches, nil
}

func decodeMatchesAsOutcomes(encodedStrategy []EncodedRPS) ([]RPSOutcomeMatch, error) {
	decidedMatches := []RPSOutcomeMatch{}

	for _, encoded := range encodedStrategy {
		rps := RPSOutcomeMatch{Them: Rock, Outcome: Lost}

		switch encoded.Them {
		case 'A':
			rps.Them = Rock
		case 'B':
			rps.Them = Paper
		case 'C':
			rps.Them = Scissors
		default:
			return nil, fmt.Errorf("unable to match %v to an RPS selection for them: %w", encoded.Them, ErrNoRPSEnumForRune)
		}

		switch encoded.You {
		case 'X':
			rps.Outcome = Lost
		case 'Y':
			rps.Outcome = Tied
		case 'Z':
			rps.Outcome = Won
		default:
			return nil, fmt.Errorf("unable to match %v to an RPS outcome for you: %w", encoded.You, ErrNoRPSOutcomeEnumForRune)
		}

		decidedMatches = append(decidedMatches, rps)
	}

	return decidedMatches, nil
}

func scoreRPSMatches(outcomeMatches []DecidedRPSMatch) ([]ScoredRPS, error) {
	scoredStrategy := []ScoredRPS{}

	for _, evaluated := range outcomeMatches {
		score := 0

		// score on choice
		switch evaluated.You {
		case Rock:
			score++
		case Paper:
			score += 2
		case Scissors:
			score += 3
		default:
			return nil, fmt.Errorf(
				"unable to score the outcome of a match when you chose %v: %w",
				evaluated.You,
				ErrUnrecognizedRPSEnum,
			)
		}

		// score on outcome
		switch evaluated.Outcome {
		case Lost:
			score += 0
		case Tied:
			score += 3
		case Won:
			score += 6
		default:
			return nil, fmt.Errorf(
				"unable to score the outcome of a match when the outcome was %v: %w",
				evaluated.Outcome,
				ErrUnrecognizedRPSOutcome,
			)
		}

		scoredStrategy = append(scoredStrategy, ScoredRPS{Evaluated: evaluated, Score: score})
	}

	return scoredStrategy, nil
}

func decideMatches(matches []RPSMatch) ([]DecidedRPSMatch, error) {
	decidedMatches := []DecidedRPSMatch{}

	for i, rps := range matches {
		decidedMatch, err := decideRPSMatch(rps)
		if err != nil {
			return nil, fmt.Errorf("unable to decide matches after failure with match %d: %w", i, err)
		}

		decidedMatches = append(decidedMatches, decidedMatch)
	}

	return decidedMatches, nil
}

func decideRPSMatch(rps RPSMatch) (DecidedRPSMatch, error) {
	decidedMatch := DecidedRPSMatch{Them: rps.Them, You: rps.You, Outcome: Lost}

	var (
		result RPSOutcome
		err    error
	)

	switch rps.Them {
	case Rock:
		result, err = decideVsRock(rps.You)
	case Paper:
		result, err = decideVsPaper(rps.You)
	case Scissors:
		result, err = decideVsScissors(rps.You)
	default:
		return DecidedRPSMatch{}, fmt.Errorf(
			"unable to evaluate the outcome of a match when they chose %v: %w",
			rps.Them,
			ErrUnrecognizedRPSEnum,
		)
	}

	if err != nil {
		return decidedMatch, fmt.Errorf("unable to decide match: %w", err)
	}

	decidedMatch.Outcome = result

	return decidedMatch, nil
}

func decideVsRock(you RPSEnum) (RPSOutcome, error) {
	switch you {
	case Rock:
		return Tied, nil
	case Paper:
		return Won, nil
	case Scissors:
		return Lost, nil
	default:
		return Lost, fmt.Errorf(
			"unable to evaluate the outcome of a match when you chose %v: %w",
			you,
			ErrUnrecognizedRPSEnum,
		)
	}
}

func decideVsPaper(you RPSEnum) (RPSOutcome, error) {
	switch you {
	case Rock:
		return Lost, nil
	case Paper:
		return Tied, nil
	case Scissors:
		return Won, nil
	default:
		return Lost, fmt.Errorf(
			"unable to evaluate the outcome of a match when you chose %v: %w",
			you,
			ErrUnrecognizedRPSEnum,
		)
	}
}

func decideVsScissors(you RPSEnum) (RPSOutcome, error) {
	switch you {
	case Rock:
		return Won, nil
	case Paper:
		return Lost, nil
	case Scissors:
		return Tied, nil
	default:
		return Lost, fmt.Errorf(
			"unable to evaluate the outcome of a match when you chose %v: %w",
			you,
			ErrUnrecognizedRPSEnum,
		)
	}
}

func decodeMatches(encodedStrategy []EncodedRPS) ([]RPSMatch, error) {
	strategy := []RPSMatch{}

	for _, encoded := range encodedStrategy {
		rps := RPSMatch{Them: Rock, You: Rock}

		switch encoded.Them {
		case 'A':
			rps.Them = Rock
		case 'B':
			rps.Them = Paper
		case 'C':
			rps.Them = Scissors
		default:
			return nil, fmt.Errorf("unable to match %v to an RPS selection for them: %w", encoded.Them, ErrNoRPSEnumForRune)
		}

		switch encoded.You {
		case 'X':
			rps.You = Rock
		case 'Y':
			rps.You = Paper
		case 'Z':
			rps.You = Scissors
		default:
			return nil, fmt.Errorf("unable to match %v to an RPS selection for you: %w", encoded.You, ErrNoRPSEnumForRune)
		}

		strategy = append(strategy, rps)
	}

	return strategy, nil
}

func parseEncodedMatchLines(lines []string) ([]EncodedRPS, error) {
	encodedStrategy := []EncodedRPS{}

	const runesToExpect = 3
	for i, line := range lines {
		if len(line) < runesToExpect {
			return nil, fmt.Errorf("unable to parse line %d, because it only has %d runes: %w", i, len(line), ErrNotEnoughItems)
		}

		encodedStrategy = append(encodedStrategy, EncodedRPS{Them: rune(line[0]), You: rune(line[2])})
	}

	return encodedStrategy, nil
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
