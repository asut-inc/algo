package algo

import (
	"reflect"
	"testing"
)

type slice struct {
	data     []int
	expected []int
	name     string
}

func sliceTester(s func([]int) []int, t *testing.T) {
	for _, data := range dataProvider {
		input := make([]int, len(data.data))
		n := copy(input, data.data)
		if n < 1 {
			t.Error("corrupt data")
		}
		result := s(input)

		if !reflect.DeepEqual(result, data.expected) {
			t.Error(
				GetFunctionName(s),
				"\ncase:", data.name,
				"\ninput:", data.data,
				"\nexpected:", data.expected,
				"\nresult:", result,
			)
		}
	}
}

var dataProvider = []slice{
	{
		data:     []int{89, 144, 100, 233, 2, 155, 55, 34, 1},
		expected: []int{1, 2, 34, 55, 89, 100, 144, 155, 233},
		name:     "...mid..., 100 - 233",
	},
	{
		data:     []int{300000, 100500, 1033, 2000, 899, 13},
		expected: []int{13, 899, 1033, 2000, 100500, 300000},
		name:     "reverse",
	},
}
