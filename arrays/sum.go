// Package arrays showcases how to use arrays with generics
package arrays

func Sum(numbers []int) int {
	add := func(acc, val int) int { return acc + val }
	return Reduce(numbers, add, 0)
}

func SumAll(numbers ...[]int) []int {
	sum := func(acc, val []int) []int {
		acc = append(acc, Sum(val))
		return acc
	}
	return Reduce(numbers, sum, []int{})
}

func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, val []int) []int {
		if len(val) == 0 {
			return append(acc, 0)
		} else {
			return append(acc, Sum(val[1:]))
		}
	}
	return Reduce(numbers, sumTail, []int{})
}
