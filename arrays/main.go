package main

func Sum(input []int) int {
	sum := 0
	for _, v := range input {
		sum += v
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var res []int
	for _, s := range slices {
		res = append(res, Sum(s))
	}
	return res
}

func SumAllTails(slices ...[]int) []int {
	var res []int
	for _, s := range slices {
		tail := s[1:]
		res = append(res, Sum(tail))
	}
	return res
}
