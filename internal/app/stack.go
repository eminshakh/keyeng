package app

import "strings"

type Stack struct {
	elements []string
}

func New(capacity int) *Stack {
	return &Stack{
		elements: make([]string, 0, capacity),
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Empty() {
	s.elements = s.elements[0:0]
}

func (s *Stack) Push(e string) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		var zero string
		return zero, false
	}
	e := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return e, true
}

func (s *Stack) Peek() (string, bool) {
	if s.IsEmpty() {
		var zero string
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *Stack) Len() int {
	return len(s.elements)
}

func (s *Stack) String() string {
	builder := strings.Builder{}
	for _, element := range s.elements {
		builder.WriteString(element)
	}
	return builder.String()
}
