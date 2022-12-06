package day_6

func Day6(input string) (int, int) {
	var first, second int
	for i, v := range []int{4, 14} {
		if i == 0 {
			first = getChars(input, v)
		} else {
			second = getChars(input, v)
		}

	}

	return first, second
}

func getChars(input string, uniqChars int) int {
	var count int

	buffer := make([]rune, 0, uniqChars)

	for _, c := range input {
		count++
		buffer = append(buffer, c)
		if len(buffer) < uniqChars {
			continue
		}

		unique := map[rune]bool{}
		for _, b := range buffer {
			unique[b] = true
		}

		if len(unique) == uniqChars {
			break
		}

		buffer = append([]rune{}, buffer[1:]...)
	}

	return count
}
