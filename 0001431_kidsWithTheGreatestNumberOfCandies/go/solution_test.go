package solution_test

import (
	"reflect"
	"testing"

	solution "github.com/XOzymandiasz/leetcode-solutions/0001431_kidsWithTheGreatestNumberOfCandies/go"
)

func TestKidsWithCandies(t *testing.T) {
	testCases := []struct {
		name         string
		candies      []int
		extraCandies int
		want         []bool
	}{
		{
			"base",
			[]int{4, 5, 3, 1, 2},
			2,
			[]bool{true, true, true, false, false},
		},
		{
			"no extra candies",
			[]int{4, 5, 3, 1, 2},
			0,
			[]bool{false, true, false, false, false},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := solution.KidsWithCandies(tc.candies, tc.extraCandies); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("given: candies = %v; extraCandies = %v, got: %v, want: %v",
					tc.candies,
					tc.extraCandies,
					got,
					tc.want,
				)
			}
		})
	}
}
