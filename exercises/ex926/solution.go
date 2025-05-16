package ex926

func minFlipsMonoIncr(s string) int {
	symbols := []rune(s)
	zeros, ones := 0, 0
	for _, r := range symbols {
		if r == '0' {
			zeros++
		} else {
			ones++
		}
	}
	leftInd, rightInd := 0, len(symbols)-1

	trim := func() {
		for leftInd < rightInd && symbols[leftInd] == '0' {
			zeros--
			leftInd++
		}
		for rightInd > leftInd && symbols[rightInd] == '1' {
			ones--
			rightInd--
		}
	}

	counter := 0
	trim()

	for leftInd < rightInd {
		counter += 1
		if ones > zeros {
			zeros--
			rightInd--
		} else {
			ones--
			leftInd++
		}
		trim()
	}

	return counter
}
