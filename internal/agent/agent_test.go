package agent_test

import (
	"testing"

	"github.com/3SMA3/distributed-calculator/internal/agent"
	"github.com/stretchr/testify/assert"
)

func TestComputeExpression(t *testing.T) {
	t.Run("Simple addition", func(t *testing.T) {
		result, err := agent.ComputeExpression("2+2")
		assert.NoError(t, err)
		assert.Equal(t, 4.0, result)
	})

	t.Run("Multiplication before addition", func(t *testing.T) {
		result, err := agent.ComputeExpression("2+2*2")
		assert.NoError(t, err)
		assert.Equal(t, 6.0, result)
	})

	t.Run("With parentheses", func(t *testing.T) {
		result, err := agent.ComputeExpression("(2+3)*4")
		assert.NoError(t, err)
		assert.Equal(t, 20.0, result)
	})

	t.Run("Division", func(t *testing.T) {
		result, err := agent.ComputeExpression("10/(2+3)")
		assert.NoError(t, err)
		assert.Equal(t, 2.0, result)
	})

	t.Run("Division by zero", func(t *testing.T) {
		_, err := agent.ComputeExpression("10/0")
		assert.Error(t, err)
		assert.Equal(t, "division by zero", err.Error())
	})

	t.Run("Invalid expression", func(t *testing.T) {
		_, err := agent.ComputeExpression("2+*3")
		assert.Error(t, err)
	})
}
