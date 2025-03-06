package agent_test

import (
	"testing"
	"time"

	"github.com/3SMA3/distributed-calculator/internal/agent"
	"github.com/stretchr/testify/assert"
)

func TestComputeTask(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		task := agent.Task{
			Arg1:          2.5,
			Arg2:          3.5,
			Operation:     "+",
			OperationTime: 100,
		}

		result, err := agent.Compute(task)
		assert.NoError(t, err)
		assert.Equal(t, 6.0, result)
	})

	t.Run("Subtraction", func(t *testing.T) {
		task := agent.Task{
			Arg1:          5.0,
			Arg2:          2.5,
			Operation:     "-",
			OperationTime: 100,
		}

		result, err := agent.Compute(task)
		assert.NoError(t, err)
		assert.Equal(t, 2.5, result)
	})

	t.Run("Multiplication", func(t *testing.T) {
		task := agent.Task{
			Arg1:          3.0,
			Arg2:          4.0,
			Operation:     "*",
			OperationTime: 100,
		}

		result, err := agent.Compute(task)
		assert.NoError(t, err)
		assert.Equal(t, 12.0, result)
	})

	t.Run("Division", func(t *testing.T) {
		task := agent.Task{
			Arg1:          10.0,
			Arg2:          2.0,
			Operation:     "/",
			OperationTime: 100,
		}

		result, err := agent.Compute(task)
		assert.NoError(t, err)
		assert.Equal(t, 5.0, result)
	})

	t.Run("Division by zero", func(t *testing.T) {
		task := agent.Task{
			Arg1:          10.0,
			Arg2:          0.0,
			Operation:     "/",
			OperationTime: 100,
		}

		result, err := agent.Compute(task)
		assert.Error(t, err)
		assert.Equal(t, "division by zero", err.Error())
		assert.Equal(t, 0.0, result)
	})

	t.Run("Invalid operation", func(t *testing.T) {
		task := agent.Task{
			Arg1:          10.0,
			Arg2:          5.0,
			Operation:     "%",
			OperationTime: 100,
		}

		result, err := agent.Compute(task)
		assert.Error(t, err)
		assert.Equal(t, "unknown operation", err.Error())
		assert.Equal(t, 0.0, result)
	})

	t.Run("Operation time", func(t *testing.T) {
		task := agent.Task{
			Arg1:          2.0,
			Arg2:          3.0,
			Operation:     "+",
			OperationTime: 100,
		}

		start := time.Now()
		_, err := agent.Compute(task)
		duration := time.Since(start)

		assert.NoError(t, err)
		assert.True(t, duration >= 100*time.Millisecond, "Operation should take at least 100ms")
	})
}
