package maxnumber

func MaxValueNumber(number []int) int {
	result := 0
	counter := (len(number) - 1)

	for i := 0; i < len(number); i++ {
		if number[i] != number[counter] {
			continue
		}
		if i < counter {
			counter--
			if result < number[i] {
				result = number[i]
			}
		}
	}

	return result
}
