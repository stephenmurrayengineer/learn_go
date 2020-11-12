package main

func Sum(numbers []int) (sum int) {
	sum = 0
	// for <idx>,<val> in <array>
	for _, number := range numbers {
		sum += number
	}
	return sum
}
