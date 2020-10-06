package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	tests := []struct {
		name     string
		data     []interface{}
	}{
		{
			name:     "Right data",
			data:     []interface{}{2, 3, 4},
		},
		{
			name:     "Wrong data",
			data:     []interface{}{4, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack(3)

			for _, i := range tt.data {
				assert.NoError(t, stack.Push(i), "Error adding new element to stack")
			}

			stack.Pop()

			stack.Print()
		})
	}

}
