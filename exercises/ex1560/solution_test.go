package ex1560

import (
	"fmt"
	"leetcode/utils"
	"slices"
	"testing"
	"time"
)

func TestMostVisited(t *testing.T) {
	var tests = []struct {
		n      int
		rounds []int
		want   []int
	}{
		{4, []int{1, 3, 1, 2}, []int{1, 2}},
		{2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2}, []int{2}},
		{7, []int{1, 3, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{3, []int{3, 1}, []int{1, 3}},
	}

	for ind, tt := range tests {
		testname := fmt.Sprintf("[%d] %v|%v", ind, tt.n, tt.rounds)
		t.Run(testname, func(t *testing.T) {
			utils.InTime(1*time.Second, t, func() {
				ans := mostVisited(tt.n, tt.rounds)
				if !slices.Equal(ans, tt.want) {
					t.Errorf("got %#v, want %#v", ans, tt.want)
				}
			})
		})
	}
}

func BenchmarkMostVisited(t *testing.B) {
	for t.Loop() {
		mostVisited(4, []int{1, 3, 1, 2})
	}
}
