package algo

// InsertionSort golang implementation of insertion sort algorithm
func InsertionSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		cur := s[i]
		j := i - 1
		for j >= 0 && s[j] > cur {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = cur
	}

	return s
}
