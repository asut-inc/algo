package algo

import (
	"reflect"
	"testing"
)

type slice struct {
	cases    []int
	expected []int
}

func sliceTester(s func([]int) []int, t *testing.T) {
	for _, data := range dataProvider {
		input := make([]int, len(data.cases))
		n := copy(input, data.cases)
		if n < 1 {
			t.Error("corrupt data")
		}
		result := s(input)

		if !reflect.DeepEqual(result, data.expected) {
			t.Error(
				GetFunctionName(s),
				"\ninput:", data.cases,
				"\nexpected:", data.expected,
				"\nresult:", result,
			)
		}
	}
}

var dataProvider = []slice{
	{
		cases:    []int{89, 144, 233, 55, 34, 1},
		expected: []int{1, 34, 55, 89, 144, 233},
	},
	{
		cases:    []int{300000, 100500, 1033, 2000, 899, 13},
		expected: []int{13, 899, 1033, 2000, 100500, 300000},
	},
}
