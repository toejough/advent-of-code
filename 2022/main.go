package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	Rock RPSEnum = iota
	Paper
	Scissors
)

const (
	Lost RPSOutcome = iota
	Tied
	Won
)

var (
	ErrMissingSolver           = fmt.Errorf("missing solver")
	ErrNoCommonItemFound       = fmt.Errorf("no common item found")
	ErrNoMaxPossible           = fmt.Errorf("no max possible: input list was empty")
	ErrNoRPSEnumForRune        = fmt.Errorf("no RPS enum was found for the rune")
	ErrNoRPSOutcomeEnumForRune = fmt.Errorf("no RPS outcome enum was found for the rune")
	ErrNoScoreMappedForItem    = fmt.Errorf("no score mapped for item")
	ErrNotEnoughItems          = fmt.Errorf("not enough items")
	ErrOddRucksackLength       = fmt.Errorf("odd rucksack item count")
	ErrUnrecognizedRPSEnum     = fmt.Errorf("unrecognized RPS enum")
	ErrUnrecognizedRPSOutcome  = fmt.Errorf("unrecognized RPS outcome enum")
)

type (
	RPSEnum    int
	RPSOutcome int
)

type EncodedRPS struct {
	Them rune
	You  rune
}

type RPSMatch struct {
	Them RPSEnum
	You  RPSEnum
}

type RPSOutcomeMatch struct {
	Them    RPSEnum
	Outcome RPSOutcome
}

type DecidedRPSMatch struct {
	Them    RPSEnum
	You     RPSEnum
	Outcome RPSOutcome
}

type ScoredRPS struct {
	Evaluated DecidedRPSMatch
	Score     int
}

type compartment struct {
	items []rune
}

type rucksack struct {
	compartment1, compartment2 compartment
}

func main() {
	// read input
	day := os.Args[1]
	part := os.Args[2]

	filename := fmt.Sprintf("%s-input-puzzle.txt", day)

	text, err := readFileText(filename)
	if err != nil {
		log.Fatal(err)
	}

	// solve
	answer, err := solve(day, part, text)
	if err != nil {
		log.Fatal(err)
	}

	// print output
	log.Printf("Answer: %s", answer)
}

func readFileText(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("couldn't read text from %s: %w", filename, err)
	}

	text := string(bytes)

	return text, nil
}

func solve(day string, part string, text string) (string, error) {
	solverMap := map[string]map[string]func(string) (string, error){
		"day1": {
			"part1": solveDay1Part1,
			"part2": solveDay1Part2,
		},
		"day2": {
			"part1": solveDay2Part1,
			"part2": solveDay2Part2,
		},
		"day3": {
			"part1": solveDay3Part1,
			// "part2": solveDay3Part2,
		},
	}

	partMap, dayOk := solverMap[day]
	if !dayOk {
		return "", fmt.Errorf("cannot solve day %s: %w", day, ErrMissingSolver)
	}

	solver, partOk := partMap[part]
	if !partOk {
		return "", fmt.Errorf("cannot solve day %s part %s: %w", day, part, ErrMissingSolver)
	}

	answer, err := solver(text)
	if err != nil {
		return "", fmt.Errorf("solving failed: %w", err)
	}

	return answer, nil
}

func maps[F any, T any](from []F, transform func(F) (T, error)) (to []T, err error) {
	for i, item := range from {
		var transformed T

		transformed, err = transform(item)
		if err != nil {
			return nil, fmt.Errorf("could not map item %d: %w", i, err)
		}

		to = append(to, transformed)
	}

	return
}

func mapsNoErr[F any, T any](from []F, transform func(F) T) (to []T) {
	for _, item := range from {
		to = append(to, transform(item))
	}

	return
}

func solveDay1Part1(text string) (string, error) {
	// split into lists of calories
	hunks := splitNoEmpty(text, "\n\n")
	stringLists := mapsNoErr(hunks, func(s string) []string {
		return splitNoEmpty(s, "\n")
	})

	lists, err := maps(stringLists, func(stringList []string) ([]int, error) {
		return maps(stringList, strconv.Atoi)
	})
	if err != nil {
		return "", err
	}

	// sum the lists
	sums := mapsNoErr(lists, sum)

	// max
	if len(sums) == 0 {
		return "", fmt.Errorf("unable to solve: %w", ErrNoMaxPossible)
	}

	max := reduce(sums[0], sums[1:], returnLarger)

	return fmt.Sprintf("%d", max), nil
}

func returnLarger(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func solveDay1Part2(text string) (string, error) {
	// split into lists of calories
	hunks := splitNoEmpty(text, "\n\n")
	stringLists := mapsNoErr(hunks, func(s string) []string {
		return splitNoEmpty(s, "\n")
	})

	lists, err := maps(stringLists, func(stringList []string) ([]int, error) {
		return maps(stringList, strconv.Atoi)
	})
	if err != nil {
		return "", err
	}

	// sum the lists
	sums := mapsNoErr(lists, sum)

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

func solveDay2Part1(text string) (string, error) {
	// split into lines
	lines := splitNoEmpty(text, "\n")
	// parse lines into opponent/you
	encodedMatches, err := maps(lines, parseEncodedMatchLine)
	if err != nil {
		return "", fmt.Errorf("unable to solve after match parsing failure: %w", err)
	}
	// convert into nicer enum representation
	matches, err := maps(encodedMatches, decodeMatch)
	if err != nil {
		return "", fmt.Errorf("unable to solve after match decoding failure: %w", err)
	}
	// enhance with outcome of encounters
	decidedMatches, err := maps(matches, decideRPSMatch)
	if err != nil {
		return "", fmt.Errorf("unable to solve after match decision failure: %w", err)
	}
	// enhance with scores
	scoredStrategy, err := maps(decidedMatches, scoreRPSMatch)
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
	encodedMatches, err := maps(lines, parseEncodedMatchLine)
	if err != nil {
		return "", fmt.Errorf("unable to solve after match parsing failure: %w", err)
	}
	// convert into nicer enum representation
	decidedMatches, err := maps(encodedMatches, decodeMatchAsOutcome)
	if err != nil {
		return "", fmt.Errorf("unable to solve after match decoding failure: %w", err)
	}
	// enhance with roll to reach outcome
	deducedMoves, err := maps(decidedMatches, deduceMove)
	if err != nil {
		return "", fmt.Errorf("unable to solve after match decision failure: %w", err)
	}
	// enhance with scores
	scoredStrategy, err := maps(deducedMoves, scoreRPSMatch)
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

func solveDay3Part1(text string) (string, error) {
	// split into lines
	lines := splitNoEmpty(text, "\n")

	// parse lines into contents in compartments
	rucksacks, err := maps(lines, parseRucksackLine)
	if err != nil {
		return "", fmt.Errorf("unable to solve: %w", err)
	}

	// reduce to common items
	commonItems, err := maps(rucksacks, identifyCommonItem)
	if err != nil {
		return "", fmt.Errorf("unable to solve: %w", err)
	}

	// score the items
	scores, err := maps(commonItems, scoreRucksackItem)
	if err != nil {
		return "", fmt.Errorf("unable to solve: %w", err)
	}

	// sum them
	total := sum(scores)

	return fmt.Sprintf("%d", total), nil
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

func scoreRucksackItem(item rune) (score int, err error) {
	const (
		lowercaseStartingScore = 1
		uppercaseStartingScore = 27
	)

	if item >= 'a' && item <= 'z' {
		distance := item - 'a'
		score = lowercaseStartingScore + int(distance)

		return score, nil
	} else if item >= 'A' && item <= 'Z' {
		distance := item - 'A'
		score = uppercaseStartingScore + int(distance)

		return score, nil
	}

	return score, fmt.Errorf("unable to score %#v: %w", item, ErrNoScoreMappedForItem)
}

func identifyCommonItem(sack rucksack) (item rune, err error) {
	for _, item := range sack.compartment1.items {
		if contains(sack.compartment2.items, item) {
			return item, nil
		}
	}

	return item, fmt.Errorf(
		"unable to reduce to common item between %v and %v: %w",
		sack.compartment1.items,
		sack.compartment2.items,
		ErrNoCommonItemFound,
	)
}

func contains(r []rune, item rune) bool {
	for _, item2 := range r {
		if item == item2 {
			return true
		}
	}

	return false
}

func parseRucksackLine(line string) (r rucksack, err error) {
	size := len(line)

	const numCompartments = 2

	compartmentSize := size / numCompartments
	if size != compartmentSize*2 {
		return r, fmt.Errorf("unable to parse rucksack line with length (%d): %w", size, ErrOddRucksackLength)
	}

	r.compartment1.items = ([]rune)(line[:compartmentSize])
	r.compartment2.items = ([]rune)(line[compartmentSize:])

	return r, nil
}

func deduceMove(match RPSOutcomeMatch) (DecidedRPSMatch, error) {
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

	outcomeAndMe, themOk := outcomeMap[match.Them]
	if !themOk {
		return DecidedRPSMatch{}, fmt.Errorf("unable to deduce moves with %v: %w", match.Them, ErrUnrecognizedRPSEnum)
	}

	you, outcomeOK := outcomeAndMe[match.Outcome]
	if !outcomeOK {
		return DecidedRPSMatch{}, fmt.Errorf("unable to deduce moves with %v: %w", match.Outcome, ErrNoRPSOutcomeEnumForRune)
	}

	return DecidedRPSMatch{Them: match.Them, You: you, Outcome: match.Outcome}, nil
}

func decodeMatchAsOutcome(encoded EncodedRPS) (RPSOutcomeMatch, error) {
	them, err := decodeThemRPS(encoded)
	if err != nil {
		return RPSOutcomeMatch{}, err
	}

	youRPSMap := map[rune]RPSOutcome{
		'X': Lost,
		'Y': Tied,
		'Z': Won,
	}

	you, youOK := youRPSMap[encoded.You]
	if !youOK {
		return RPSOutcomeMatch{}, fmt.Errorf(
			"unable to match %v to an RPS outcome for you: %w",
			encoded.You,
			ErrNoRPSOutcomeEnumForRune,
		)
	}

	return RPSOutcomeMatch{Them: them, Outcome: you}, nil
}

func decodeThemRPS(encoded EncodedRPS) (RPSEnum, error) {
	themRPSMap := map[rune]RPSEnum{
		'A': Rock,
		'B': Paper,
		'C': Scissors,
	}

	them, themOK := themRPSMap[encoded.Them]
	if !themOK {
		return 0, fmt.Errorf(
			"unable to match %v to an RPS selection for them: %w",
			encoded.Them,
			ErrNoRPSEnumForRune,
		)
	}

	return them, nil
}

func scoreRPSMatch(outcomeMatch DecidedRPSMatch) (ScoredRPS, error) {
	score := 0

	// score on choice
	switch outcomeMatch.You {
	case Rock:
		score++
	case Paper:
		score += 2
	case Scissors:
		score += 3
	default:
		return ScoredRPS{}, fmt.Errorf(
			"unable to score the outcome of a match when you chose %v: %w",
			outcomeMatch.You,
			ErrUnrecognizedRPSEnum,
		)
	}

	// score on outcome
	switch outcomeMatch.Outcome {
	case Lost:
		score += 0
	case Tied:
		score += 3
	case Won:
		score += 6
	default:
		return ScoredRPS{}, fmt.Errorf(
			"unable to score the outcome of a match when the outcome was %v: %w",
			outcomeMatch.Outcome,
			ErrUnrecognizedRPSOutcome,
		)
	}

	return ScoredRPS{Evaluated: outcomeMatch, Score: score}, nil
}

func decideRPSMatch(rps RPSMatch) (DecidedRPSMatch, error) {
	decidedMatch := DecidedRPSMatch{Them: rps.Them, You: rps.You, Outcome: Lost}

	if rps.Them == rps.You {
		decidedMatch.Outcome = Tied
		return decidedMatch, nil
	}

	// win map: given them, you need what to win?
	winMap := map[RPSEnum]RPSEnum{
		Rock:     Paper,
		Paper:    Scissors,
		Scissors: Rock,
	}

	need, ok := winMap[rps.Them]
	if !ok {
		return DecidedRPSMatch{}, fmt.Errorf(
			"unable to evaluate the outcome of a match when they chose %v: %w",
			rps.Them,
			ErrUnrecognizedRPSEnum,
		)
	}

	if need == rps.You {
		decidedMatch.Outcome = Won
	} else {
		decidedMatch.Outcome = Lost
	}

	return decidedMatch, nil
}

func decodeMatch(encoded EncodedRPS) (RPSMatch, error) {
	them, err := decodeThemRPS(encoded)
	if err != nil {
		return RPSMatch{}, err
	}

	youRPSMap := map[rune]RPSEnum{
		'X': Rock,
		'Y': Paper,
		'Z': Scissors,
	}

	you, youOK := youRPSMap[encoded.You]
	if !youOK {
		return RPSMatch{}, fmt.Errorf(
			"unable to match %v to an RPS selection for you: %w",
			encoded.You,
			ErrNoRPSEnumForRune,
		)
	}

	return RPSMatch{Them: them, You: you}, nil
}

func parseEncodedMatchLine(line string) (EncodedRPS, error) {
	const runesToExpect = 3
	if len(line) < runesToExpect {
		return EncodedRPS{}, fmt.Errorf("needed %d runes but got %d: %w", runesToExpect, len(line), ErrNotEnoughItems)
	}

	return EncodedRPS{Them: rune(line[0]), You: rune(line[2])}, nil
}

func sum(list []int) int {
	return reduce(0, list, func(a, b int) int { return a + b })
}

func reduce[T any](initial T, toReduce []T, reducer func(a, b T) T) T {
	result := initial
	for _, item := range toReduce {
		result = reducer(result, item)
	}

	return result
}
