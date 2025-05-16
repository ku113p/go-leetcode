package ex926

import (
	"fmt"
	"leetcode/utils"
	"testing"
	"time"
)

func TestMinFlipsMonoIncr(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"00110", 1},
		{"11111", 0},
		{"010110", 2},
		{"00011000", 2},
		{"1111001110", 3},
		{"0101100011", 3},
		{"10011111110010111011", 5},
		{utils.ReadStringFromFile("../files/ex926_biginput.txt"), 26641},
	}

	for ind, tt := range tests {
		testname := fmt.Sprintf("[%d] %s...", ind, string([]rune(tt.s[:5])))
		t.Run(testname, func(t *testing.T) {
			utils.InTime(1*time.Second, t, func() {
				ans := minFlipsMonoIncr(tt.s)
				if ans != tt.want {
					t.Errorf("got %d, want %d", ans, tt.want)
				}
			})
		})
	}
}

func BenchmarkMinFlipsMonoIncr(b *testing.B) {
	for b.Loop() {
		minFlipsMonoIncr("0001010101010110110111000")
	}
}
