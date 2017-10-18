package pancakes

import "testing"

var testData = map[string]struct {
	data     string
	expected int
}{
	"Test: +":      {"+", 0},
	"Test: -":      {"-", 1},
	"Test: +-":     {"+-", 2},
	"Test: +-+-+-": {"+-+-+-", 6},
}

func Test(t *testing.T) {
	for key, value := range testData {
		flips := SortStack(value.data)
		if flips != value.expected {
			t.Errorf("[%s] expected %d, got %d", key, value.expected, flips)
		}
	}
}
