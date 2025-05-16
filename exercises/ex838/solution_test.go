package ex838

import (
	"fmt"
	"leetcode/utils"
	"testing"
	"time"
)

func TestPushDominoes(t *testing.T) {
	var tests = []struct {
		a    string
		want string
	}{
		{"RR.L", "RR.L"},
		{".L.R...LR..L..", "LL.RR.LLRRLL.."},
		{".L.R.", "LL.RR"},
		{"R.R.L", "RRR.L"},
		{"L.LL......", "LLLL......"},
	}

	for ind, tt := range tests {
		testname := fmt.Sprintf("[%d] %s", ind, tt.a)
		t.Run(testname, func(t *testing.T) {
			utils.InTime(1*time.Second, t, func() {
				ans := pushDominoes(tt.a)
				if ans != tt.want {
					t.Errorf("got %#v, want %#v", ans, tt.want)
				}
			})
		})
	}
}

func BenchmarkPushDominoes(t *testing.B) {
	for t.Loop() {
		pushDominoes(".L.R...LR..L..")
	}
}
