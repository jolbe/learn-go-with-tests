package main

import (
	"fmt"

	"github.com/gregor-pifko/learn-go-with-tests/generics"
)

func main() {
	stack := generics.StackOfInts{}
	fmt.Println("is empty:", stack.IsEmpty())
	stack.Push(42)
	fmt.Println(stack, stack.IsEmpty())

	elem, ok := stack.Pop()
	fmt.Println(elem)
	elem, ok = stack.Pop()
	fmt.Println(elem)
	if !ok {
		fmt.Println(ok)
	}

	stack.Push(42)
	stack.Push("string inside stack of ints")
	fmt.Println(stack)
}
