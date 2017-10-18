package bowling

import "testing"

var testData = map[string]struct {
	throws   []int
	expected int
}{
	"TestAllGutters":  {[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0},
	"TestAllOnes":     {[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 20},
	"TestSpare":       {[]int{5, 5, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 16},
	"TestStrike":      {[]int{10, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 22},
	"TestPerfectGame": {[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 0, 0, 0, 0, 0, 0, 0, 0}, 300},
}

func Test(t *testing.T) {
	for key, value := range testData {
		g := NewGame()
		g.SetThrows(value.throws)
		score := g.GetScore()
		if score != value.expected {
			t.Errorf("[%s]: expected %d, got %d", key, value.expected, score)
		}
	}
}
