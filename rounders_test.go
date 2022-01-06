package rounders_test

import (
	"fmt"
	"testing"

	"github.com/kklash/rounders"
)

func TestRoundToSigFigs(t *testing.T) {
	fixtures := []struct {
		input, output float64
		sigFigs       int
	}{
		{0, 0, 1},
		{1, 1, 1},
		{1, 1, 2},
		{1.2, 1, 1},
		{1.2, 1.2, 2},
		{1.23, 1, 1},
		{1.23, 1.2, 2},
		{1.23, 1.23, 3},
		{1.23, 1.23, 4},
		{1.263, 1, 1},
		{1.263, 1.3, 2},
		{1.263, 1.26, 3},
		{1.236, 1.24, 3},
		{0.2363, 0.236, 3},
		{0.969856, 0.97, 3},
		{0.969856, 0.9699, 4},
		{0.03774, 0.0377, 3},
		{1025, 1000, 1},
		{1025, 1000, 2},
		{1025, 1030, 3},
		{10611, 10600, 3},
		{-1234, -1230, 3},
		{-0.02838, -0.03, 1},
		{-0.02838, -0.028, 2},
		{-0.02838, -0.0284, 3},
	}

	for _, fixture := range fixtures {
		result := rounders.RoundToSigFigs(fixture.input, fixture.sigFigs)
		if result != fixture.output {
			t.Errorf("Failed to round to %d sig figs\nWanted %v\nGot    %v", fixture.sigFigs, fixture.output, result)
		}
	}
}

func TestTruncToSigFigs(t *testing.T) {
	fixtures := []struct {
		input, output float64
		sigFigs       int
	}{
		{0, 0, 1},
		{1, 1, 1},
		{1, 1, 2},
		{1.2, 1, 1},
		{1.2, 1.2, 2},
		{1.23, 1, 1},
		{1.23, 1.2, 2},
		{1.23, 1.23, 3},
		{1.23, 1.23, 4},
		{1.263, 1, 1},
		{1.263, 1.2, 2},
		{1.263, 1.26, 3},
		{1.236, 1.23, 3},
		{0.2363, 0.236, 3},
		{0.969856, 0.969, 3},
		{0.969856, 0.9698, 4},
		{0.03774, 0.0377, 3},
		{1025, 1000, 1},
		{1025, 1000, 2},
		{1025, 1020, 3},
		{10611, 10600, 3},
		{-1234, -1230, 3},
		{-0.02838, -0.02, 1},
		{-0.02838, -0.028, 2},
		{-0.02838, -0.0283, 3},
	}

	for _, fixture := range fixtures {
		result := rounders.TruncToSigFigs(fixture.input, fixture.sigFigs)
		if result != fixture.output {
			t.Errorf("Failed to truncate to %d sig figs\nWanted %v\nGot    %v", fixture.sigFigs, fixture.output, result)
		}
	}
}

func ExampleRoundToDecimals() {
	fmt.Println(rounders.RoundToDecimals(1.3888, 3))
	// output:
	// 1.389
}

func ExampleTruncToDecimals() {
	fmt.Println(rounders.TruncToDecimals(1.3888, 3))
	// output:
	// 1.388
}
