package primefactors

import (
	"reflect"
	"testing"
)

var testData = map[string]struct {
	value    int
	expected []int
}{
	"Test1":  {1, []int{}},
	"Test2":  {2, []int{2}},
	"Test3":  {3, []int{3}},
	"Test4":  {4, []int{2, 2}},
	"Test5":  {5, []int{5}},
	"Test6":  {6, []int{2, 3}},
	"Test7":  {7, []int{7}},
	"Test8":  {8, []int{2, 2, 2}},
	"Test9":  {9, []int{3, 3}},
	"Test10": {10, []int{2, 5}},
}

func Test(t *testing.T) {
	for key, data := range testData {
		result := GeneratePrimeFactors(data.value)
		if !reflect.DeepEqual(result, data.expected) {
			t.Errorf("[%s] expected %v, got %v", key, data.expected, result)
		}
	}
}
