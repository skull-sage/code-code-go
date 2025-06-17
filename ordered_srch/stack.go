package ordered_srch

// Stack is a generic last-in-first-out (LIFO) data structure
// that supports operations like Push, Pop, and Peek.
// The zero value for Stack is an empty stack ready to use.
type Stack[T any] struct {
	items []T
}

// NewStack creates and returns a new empty Stack.
// Example:
//
//	stack := NewStack[int]()
//	stack.Push(1)
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push adds an item to the top of the stack.
// Time complexity: O(1) amortized
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
// If the stack is empty, returns the zero value of T and false.
// Time complexity: O(1)
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

// Peek returns the top item without removing it from the stack.
// If the stack is empty, returns the zero value of T and false.
// Time complexity: O(1)
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack contains no items.
// Time complexity: O(1)
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
// Time complexity: O(1)
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Clear removes all items from the stack.
// Time complexity: O(1)
func (s *Stack[T]) Clear() {
	s.items = s.items[:0]
}