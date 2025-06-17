package ordered_srch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("integer stack operations", func(t *testing.T) {
		stack := NewStack[int]()

		// Test empty stack
		assert.True(t, stack.IsEmpty())
		assert.Equal(t, 0, stack.Size())

		// Test push operations
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		assert.False(t, stack.IsEmpty())
		assert.Equal(t, 3, stack.Size())

		// Test peek operation
		val, ok := stack.Peek()
		assert.True(t, ok)
		assert.Equal(t, 3, val)
		assert.Equal(t, 3, stack.Size()) // Size should not change after peek

		// Test pop operations
		val, ok = stack.Pop()
		assert.True(t, ok)
		assert.Equal(t, 3, val)

		val, ok = stack.Pop()
		assert.True(t, ok)
		assert.Equal(t, 2, val)

		val, ok = stack.Pop()
		assert.True(t, ok)
		assert.Equal(t, 1, val)

		// Test empty stack after all pops
		assert.True(t, stack.IsEmpty())
		val, ok = stack.Pop()
		assert.False(t, ok)
		assert.Equal(t, 0, val)
	})

	t.Run("string stack operations", func(t *testing.T) {
		stack := NewStack[string]()

		stack.Push("hello")
		stack.Push("world")

		val, ok := stack.Pop()
		assert.True(t, ok)
		assert.Equal(t, "world", val)

		val, ok = stack.Peek()
		assert.True(t, ok)
		assert.Equal(t, "hello", val)
	})

	t.Run("clear operation", func(t *testing.T) {
		stack := NewStack[int]()

		// Push some items
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		assert.Equal(t, 3, stack.Size())

		// Clear the stack
		stack.Clear()
		assert.True(t, stack.IsEmpty())
		assert.Equal(t, 0, stack.Size())

		// Ensure we can still use the stack after clearing
		stack.Push(4)
		assert.Equal(t, 1, stack.Size())
		val, ok := stack.Pop()
		assert.True(t, ok)
		assert.Equal(t, 4, val)
	})
}