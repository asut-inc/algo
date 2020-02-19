package algo

import (
	"reflect"
	"testing"
)

type testCase struct {
	inputSlice    []int
	expectedSlice []int
}

var dataProvider = []testCase{
	{
		inputSlice:    []int{89, 144, 233, 55, 34, 1},
		expectedSlice: []int{1, 34, 55, 89, 144, 233},
	},
	{
		inputSlice:    []int{300000, 100500, 1033, 2000, 899, 13},
		expectedSlice: []int{13, 899, 1033, 2000, 100500, 300000},
	},
}

func TestQuickSort(t *testing.T) {
	for _, data := range dataProvider {
		input := make([]int, len(data.inputSlice))
		input = append(input[:0:0], data.inputSlice...)
		result := QuickSort(data.inputSlice)

		if !reflect.DeepEqual(result, data.expectedSlice) {
			t.Error(
				"test case:", input,
				"expected:", data.expectedSlice,
				"result:", result,
			)
		}
	}
}
