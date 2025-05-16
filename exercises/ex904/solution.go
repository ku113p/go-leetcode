package ex904

func totalFruit(fruits []int) int {
	if len(fruits) < 2 {
		return len(fruits)
	}

	a, b := fruits[0], fruits[1]
	max, counter, bCounter := 0, 2, 1
	if a == b {
		bCounter++
	}

	for _, c := range fruits[2:] {
		switch c {
		case b:
			bCounter++
			counter++
		case a:
			a, b = b, a
			bCounter = 1
			counter++
		default:
			a, b = b, c
			if counter > max {
				max = counter
			}
			counter = bCounter + 1
			bCounter = 1
		}
	}

	if counter > max {
		max = counter
	}

	return max
}
