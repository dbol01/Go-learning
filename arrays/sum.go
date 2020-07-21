package arrays

// Sum - adds all numbers in an array and returns the result
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll - takes any number of slices and returns a slice with totals
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails - takes any number of slices and returns the tails, totaled (all elements besides first/head)
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
