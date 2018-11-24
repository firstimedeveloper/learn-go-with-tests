package sum

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(collection ...[]int) []int {
	var result []int
	for _, v := range collection {
		result = append(result, Sum(v))
	}
	return result
}
