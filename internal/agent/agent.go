package agent

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Task struct {
	ID            string  `json:"id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	OperationTime int     `json:"operation_time"`
}

func Worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task, err := fetchTask()
		if err != nil {
			fmt.Println("Error fetching task:", err)
			time.Sleep(1 * time.Second)
			continue
		}

		result, err := compute(task)
		if err != nil {
			fmt.Println("Error computing task:", err)
		}
		sendResult(task.ID, result, err)
	}
}

func fetchTask() (*Task, error) {
	resp, err := http.Get("http://localhost:8080/internal/task")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("no tasks available")
	}

	var response struct {
		Task Task `json:"task"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response.Task, nil
}

func compute(task *Task) (float64, error) {
	time.Sleep(time.Duration(task.OperationTime) * time.Millisecond)

	switch task.Operation {
	case "+":
		return task.Arg1 + task.Arg2, nil
	case "-":
		return task.Arg1 - task.Arg2, nil
	case "*":
		return task.Arg1 * task.Arg2, nil
	case "/":
		if task.Arg2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return task.Arg1 / task.Arg2, nil
	default:
		return 0, fmt.Errorf("unknown operation")
	}
}

func sendResult(taskID string, result float64, err error) {
	payload := map[string]interface{}{
		"id":     taskID,
		"result": result,
	}
	if err != nil {
		payload["error"] = err.Error()
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post("http://localhost:8080/internal/task", "application/json", io.NopCloser(strings.NewReader(string(jsonData))))
	if err != nil {
		fmt.Println("Error sending result:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to send result")
	}
}
