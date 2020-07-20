package iteration

// Repeat - Takes a char and returns is repeated 5 times
func Repeat(char string, numberOfTimes int) string {
	var repeated string
	for i := 0; i < numberOfTimes; i++ {
		repeated += char
	}

	return repeated
}
