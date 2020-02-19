package algo

import (
	"math/rand"
	"time"
)

/* QuickSort golang implementation of quicksort algorithm
*
 */
func QuickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	rand.Seed(time.Now().Unix())
	// left, right slice indexes
	left, right := 0, len(s)-1
	// randomly chosen slice index
	pivotIndex := rand.Intn(len(s) - 1)
	//swap pivot index with last right element in slice
	s[pivotIndex], s[right] = s[right], s[pivotIndex]

	// iterate slice swap elements less than right
	for i, _ := range s {
		if s[i] < s[right] {
			s[i], s[left] = s[left], s[i]
			left++
		}
	}

	s[left], s[right] = s[right], s[left]

	QuickSort(s[:left])
	QuickSort(s[right+1:])

	return s
}
