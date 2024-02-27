package lib

import (
	"fmt"
	"slices"
	"testing"
)

func TestRepository_Average(t *testing.T) {
	r := Repository{
		prices: []Price{
			{
				Amount:    101,
				Timestamp: 12_345,
			},
			{
				Amount:    100,
				Timestamp: 12_347,
			},
			{
				Amount:    5,
				Timestamp: 40_960,
			},
			{
				Amount:    102,
				Timestamp: 12_346,
			},
		},
	}

	testCases := []struct {
		Description string
		Query
		ExpectedAverage int32
	}{
		{
			Description: "Normal usage",
			Query: Query{
				MinTime: 12_288,
				MaxTime: 16_384,
			},
			ExpectedAverage: 101,
		},
		{
			Description: "No prices in range",
			Query: Query{
				MinTime: 999_999,
				MaxTime: 1_000_000,
			},
			ExpectedAverage: 0,
		},
	}

	fmt.Printf("%v\n", r.prices)
	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			avg := r.Average(tc.MinTime, tc.MaxTime)
			if avg != tc.ExpectedAverage {
				t.Fatalf("avegare don't match: expected %d but got %d", tc.ExpectedAverage, avg)
			}
		})
	}
	fmt.Printf("%v\n", r.prices)
}

func TestRepository_Sorted(t *testing.T) {
	r := Repository{
		prices: []Price{
			{
				Amount:    101,
				Timestamp: 12_345,
			},
			{
				Amount:    100,
				Timestamp: 12_347,
			},
			{
				Amount:    5,
				Timestamp: 40_960,
			},
			{
				Amount:    102,
				Timestamp: 12_346,
			},
		},
	}

	var initial = make([]Price, len(r.prices))
	copy(initial, r.prices)

	expected := []Price{
		{

			Amount:    101,
			Timestamp: 12_345,
		},
		{
			Amount:    102,
			Timestamp: 12_346,
		},
		{
			Amount:    100,
			Timestamp: 12_347,
		},
		{
			Amount:    5,
			Timestamp: 40_960,
		},
	}

	sorted := r.Sorted()

	if !slices.Equal(sorted, expected) {
		t.Fatalf("slices don't match: expected %v but got %v", expected, sorted)
	}

	if !slices.Equal(r.prices, initial) {
		t.Fatal("method should not mutate the original slice")
	}
}
