package main

import (
	"log/slog"
	"testing"
)

func TestSolution(t *testing.T) {
	lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	want := 281

	got := 0
	for _, l := range lines {
		got += Solve(l)
	}
	if got != want {
		slog.Error("test failed", "want", want, "got", got)
		t.FailNow()
	}
}
