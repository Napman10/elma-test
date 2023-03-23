package main

import (
	"errors"
	"fmt"
)

type IntStack []int

var ErrEmptyStack = errors.New("empty stack")

func (s *IntStack) Push(x int) {
	*s = append(*s, x)
}

func (s *IntStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *IntStack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, ErrEmptyStack
	}

	return (*s)[len(*s)-1], nil
}

func (s *IntStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrEmptyStack
	}

	lastIndex := len(*s) - 1

	top, err := s.Top()
	if err != nil {
		return 0, err
	}

	*s = (*s)[:lastIndex]

	return top, nil
}

func (s *IntStack) Max() (int, error) {
	if s.IsEmpty() {
		return 0, ErrEmptyStack
	}

	max := (*s)[0]

	for _, item := range *s {
		if item > max {
			max = item
		}
	}

	return max, nil
}

func ShowTask4() {
	stack := IntStack{4, 6, 7, -3, 0, 2}
	fmt.Printf("1. base %v\n", stack)

	stack.Push(-12)
	fmt.Printf("2. add new val %v\n", stack)

	p, err := stack.Pop()
	if err != nil {
		fmt.Println("%w", err)
		return
	}
	fmt.Printf("3. pop val %d\n", p)
	fmt.Printf("4. after pop %v\n", stack)

	p, err = stack.Pop()
	if err != nil {
		fmt.Println("%w", err)
	}
	fmt.Printf("5. pop val %d\n", p)
	fmt.Printf("6. after pop %v\n", stack)

	max, err := stack.Max()
	if err != nil {
		fmt.Println("%w", err)
		return
	}
	fmt.Printf("7. max val %d\n", max)
}
