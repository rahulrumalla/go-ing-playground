package fibonacci

// 0, 1, 1, 3, 5, 8, 13, 21

func Generate(sequenceLength int) []int {
	var sequence []int
	for i := 0; i < sequenceLength-1; i++ {
		var sum int

		// handle first 2 numbers in the sequence
		if i == 0 || i == 1 {
			sum = i
		} else {
			sum = sequence[i-1] + sequence[i-2]
		}

		sequence = append(sequence, sum)
	}
	return sequence
}
