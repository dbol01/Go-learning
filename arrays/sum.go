package sum

// Sum - adds all numbers in an array and returns the result
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}
