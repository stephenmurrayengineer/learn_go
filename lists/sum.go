package main

func Sum(numbers []int) (sum int) {
	sum = 0
	// for <idx>,<val> in <array>
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	// create slice using make function:
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
