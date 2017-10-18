package fizzbuzz

import "testing"

var testData = map[string]struct {
	value    int
	expected string
}{
	"TestNoMatch": {1, ""},
	"TestFizz":    {3, FIZZ},
	"TestBuzz":    {5, BUZZ},
	"TestFizzBuzz": {15, FIZZ_BUZZ},
}

func Test(t *testing.T) {
	for key, data := range testData {
		result := ParseFizzBuzz(data.value)
		if result != data.expected {
			t.Errorf("[%s]: expected %s, got %s", key, data.expected, result)
		}
	}
}
