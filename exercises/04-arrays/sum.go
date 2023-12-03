package main

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func SumAll(slicesToSum ...[]int) []int {
	var sums []int

	for _, numbers := range slicesToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums

}

func SumAllTails(slicesToSum ...[]int) []int {
	var sums []int

	for _, numbers := range slicesToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}

	}

	return sums
}
