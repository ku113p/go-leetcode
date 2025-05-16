package ex904

import (
	"fmt"
	"leetcode/utils"
	"testing"
	"time"
)

func TestTotalFruit(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{[]int{1, 2, 1}, 3},
		{[]int{0, 1, 2, 2}, 3},
		{[]int{1, 2, 3, 2, 2}, 4},
	}

	for ind, tt := range tests {
		testname := fmt.Sprintf("[%d] %v", ind, tt.a)
		t.Run(testname, func(t *testing.T) {
			utils.InTime(1*time.Second, t, func() {
				ans := totalFruit(tt.a)
				if ans != tt.want {
					t.Errorf("got %d, want %d", ans, tt.want)
				}
			})
		})
	}
}

func BenchmarkMinFlipsMonoIncr(b *testing.B) {
	for b.Loop() {
		totalFruit([]int{1, 2, 1})
	}
}
