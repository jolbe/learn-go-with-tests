package arrays

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d; want %d; given, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3, 4}, []int{5, 6, 7}, []int{8, 9})
	want := []int{10, 18, 17}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d; want %d;", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v; want %v;", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(acc, val int) int { return acc * val }
		AssertEqual(t, Reduce([]int{1, 2, 3, 4, 5}, multiply, 1), 120)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concat := func(acc, val string) string { return acc + val }
		AssertEqual(t, Reduce([]string{"H", "e", "l", "l", "o"}, concat, ""), "Hello")
	})
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(num int) bool { return num%2 == 0 })
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})

	type Person struct {
		Name string
	}

	t.Run("find the best programer", func(t *testing.T) {
		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Fowler"},
			{Name: "Gregor Pifko"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Gregor")
		})

		AssertTrue(t, found)
		AssertEqual(t, king.Name, "Gregor Pifko")
	})
}
