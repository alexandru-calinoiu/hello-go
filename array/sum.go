package array

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			numbersToSum[i] = []int{}
		} else {
			numbersToSum[i] = numbers[1:]
		}
	}
	sums = SumAll(numbersToSum...)
	return
}
