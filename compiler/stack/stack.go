package stack

import (
	"fmt"
)

// Stack ,,,
type Stack struct {
	top      *component
	capacity int
	length   int
}

// Push ...
func (s *Stack) Push(data interface{}) error {

	if s.capacity <= s.length {
		return &stackOverflowError{}
	}

	s.top = &component{
		data: data,
		next: s.top,
	}

	s.length++
	return nil
}

// Pop ...
func (s *Stack) Pop() {
	if s.top != nil {
		s.top = s.top.next
		s.length--
	}
}

// Top ...
func (s *Stack) Top() interface{} {
	return s.top.data
}

// Print ...
func (s *Stack) Print() {
	if s.top == nil {
		fmt.Println("{ }")
		return
	}

	var temp *component = s.top

	fmt.Print("{ ")
	for {
		fmt.Print(temp.data, " ")

		if temp.next == nil {
			break
		}

		temp = temp.next
	}
	fmt.Println("}")
}

// NewStack ...
func NewStack(cap int) *Stack {

	return &Stack{
		top:      nil,
		capacity: cap,
		length:   0,
	}
}

type component struct {
	data interface{}
	next *component
}
