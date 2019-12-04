package main

import "testing"

func TestMain(t *testing.T) {
	t.Parallel()

	correctAnswer := 3305301
	result := do()

	if result != correctAnswer {
		t.Fatalf("Actual result (%v) was incorrect (should have been %v)", result, correctAnswer)
	}
}
