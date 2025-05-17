package ex153

import (
	"fmt"
	"leetcode/utils"
	"testing"
	"time"
)

func TestFindMin(t *testing.T) {
	var tests = []struct {
		a    []int
		want int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
	}

	for ind, tt := range tests {
		testname := fmt.Sprintf("[%d] %v", ind, tt.a)
		t.Run(testname, func(t *testing.T) {
			utils.InTime(1*time.Second, t, func() {
				ans := findMin(tt.a)
				if ans != tt.want {
					t.Errorf("got %#v, want %#v", ans, tt.want)
				}
			})
		})
	}
}

func BenchmarkFindMin(t *testing.B) {
	for t.Loop() {
		findMin([]int{3, 4, 5, 1, 2})
	}
}
