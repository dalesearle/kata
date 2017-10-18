package bowling

type game struct {
	throws []int
}

func NewGame() *game {
	return new(game)
}

func (g *game) SetThrows(throws []int) {
	g.throws = throws
}
func (g *game) GetScore() int {
	score := 0
	throwIndex := 0
	for frame := 0; frame < 10; frame++ {
		if g.throws[throwIndex] == 10 {
			score += g.throws[throwIndex] + g.throws[throwIndex +1] + g.throws[throwIndex + 2]
			throwIndex++
		}else if isSpare(g, throwIndex) {
			score += 10 + g.throws[throwIndex+2]
			throwIndex += 2
		} else {
			score += g.throws[throwIndex] + g.throws[throwIndex+1]
			throwIndex += 2
		}
	}
	return score
}

func isSpare(g *game, throwIndex int) bool {
	return g.throws[throwIndex]+g.throws[throwIndex+1] == 10
}
