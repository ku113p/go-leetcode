package ex1560

func mostVisited(n int, rounds []int) []int {
	if len(rounds) <= 1 {
		return rounds
	}

	repeated := []int{}
	start := rounds[0]
	end := rounds[len(rounds)-1]

	var stop int
	if end < start {
		stop = n
		for i := 1; i <= end; i++ {
			repeated = append(repeated, i)
		}
	} else {
		stop = end
	}

	for i := start; i <= stop; i++ {
		repeated = append(repeated, i)
	}

	return repeated
}
