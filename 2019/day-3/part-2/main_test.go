package main

import (
	"strconv"
	"testing"
)

func TestWireCrossings(t *testing.T) {
	testData := []struct {
		input  [2]string
		output int
	}{
		{
			input: [2]string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			},
			output: 610,
		},
		{
			input: [2]string{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			},
			output: 410,
		},
	}
	for i, td := range testData {
		td := td

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			expected := td.output
			actual := shortestCrossing(td.input)

			if expected != actual {
				t.Fatalf(
					"resulting distance (%v) was not as expected (%v) for input (%v)",
					actual,
					expected,
					td.input,
				)
			}
		})
	}
}

func TestMain(t *testing.T) {
	t.Parallel()

	correctAnswer := "20286"
	result := do()

	if result != correctAnswer {
		t.Fatalf("Actual result (%v) was incorrect (should have been %v)", result, correctAnswer)
	}
}
