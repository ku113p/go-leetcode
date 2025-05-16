package ex838

// 0 1 2 3 4 5

const pushLeft = 'L'
const pushRight = 'R'
const noPush = '.'

type lastPushedDomino struct {
	vector rune
	index  int
}

func newLastPushedDomino(r rune, i int) *lastPushedDomino {
	return &lastPushedDomino{r, i}
}

func pushDominoes(dominoes string) string {
	symbols := []rune(dominoes)

	var prevActive *lastPushedDomino

	for ind, r := range symbols {
		switch r {
		case pushRight:
			if prevActive != nil && prevActive.vector == pushRight {
				segmentSize := ind - prevActive.index - 1
				for i := range segmentSize {
					symbols[ind-segmentSize+i] = pushRight
				}
			}
			prevActive = newLastPushedDomino(pushRight, ind)
		case pushLeft:
			if prevActive != nil {
				if prevActive.vector != pushLeft {
					segmentSize := ind - prevActive.index - 1
					for i := range segmentSize / 2 {
						symbols[ind-i-1] = pushLeft
						symbols[ind-segmentSize+i] = pushRight
					}
				} else {
					segmentSize := ind - prevActive.index - 1
					for i := range segmentSize {
						symbols[ind-segmentSize+i] = pushLeft
					}
				}
			} else if ind != 0 {
				for i := range ind {
					symbols[i] = pushLeft
				}
			}
			prevActive = newLastPushedDomino(pushLeft, ind)
		case noPush:
			isLast := ind == len(symbols)-1
			needChange := isLast && prevActive != nil && prevActive.vector == pushRight
			if needChange {
				segmentSize := ind - prevActive.index
				for i := range segmentSize {
					symbols[ind-segmentSize+i+1] = pushRight
				}
			}
		}
	}

	return string(symbols)
}
