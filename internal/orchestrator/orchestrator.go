package orchestrator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Expression struct {
	ID     string  `json:"id"`
	Expr   string  `json:"expression"`
	Status string  `json:"status"`
	Result float64 `json:"result"`
}

type Task struct {
	ID            string  `json:"id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	OperationTime int     `json:"operation_time"`
}

var (
	expressions = make(map[string]Expression)
	tasks       = make(map[string]Task)
	mutex       = sync.Mutex{}
)

func HandleCalculate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Expression string `json:"expression"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusUnprocessableEntity)
		return
	}

	id := fmt.Sprintf("%d", time.Now().UnixNano())

	mutex.Lock()
	expressions[id] = Expression{
		ID:     id,
		Expr:   req.Expression,
		Status: "pending",
		Result: 0,
	}
	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func HandleGetExpressions(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var exprs []Expression
	for _, expr := range expressions {
		exprs = append(exprs, expr)
	}

	json.NewEncoder(w).Encode(map[string][]Expression{"expressions": exprs})
}

func HandleGetExpressionByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	mutex.Lock()
	expr, exists := expressions[id]
	mutex.Unlock()

	if !exists {
		http.Error(w, "Expression not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]Expression{"expression": expr})
}

func HandleTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mutex.Lock()
		for _, task := range tasks {
			delete(tasks, task.ID)
			json.NewEncoder(w).Encode(map[string]Task{"task": task})
			mutex.Unlock()
			return
		}
		mutex.Unlock()
		http.Error(w, "No tasks available", http.StatusNotFound)

	case http.MethodPost:
		var req struct {
			ID     string  `json:"id"`
			Result float64 `json:"result"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusUnprocessableEntity)
			return
		}

		mutex.Lock()
		expr, exists := expressions[req.ID]
		if !exists {
			mutex.Unlock()
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		expr.Result = req.Result
		expr.Status = "done"
		expressions[req.ID] = expr
		mutex.Unlock()

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
