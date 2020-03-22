package algo

// InsertionSort golang implementation of insertion sort algorithm
func InsertionSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		max := s[i]
		j := i
		j--
		for j >= 0 && s[j] > max {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = max
	}

	return s
}
