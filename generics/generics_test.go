package generics_test

import (
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/generics"
)

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(generics.StackOfInts)

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfInts.Push(42)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(123)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 42)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(generics.StackOfStrings)

		// check stack is empty
		AssertTrue(t, myStackOfStrings.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfStrings.Push("123")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		// add another thing, pop it back again
		myStackOfStrings.Push("456")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "456")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "123")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})

	t.Run("interface stack DX is horrid", func(t *testing.T) {
		myStackOfInts := new(generics.StackOfInts)

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()

		// get our ints from our interface{}
		reallyFirstNum, ok := firstNum.(int)
		AssertTrue(t, ok) // need to check we definitely got an int out of the interface{}

		reallySecondNum, ok := secondNum.(int)
		AssertTrue(t, ok) // and again!

		AssertEqual(t, reallyFirstNum+reallySecondNum, 3)
	})
}
