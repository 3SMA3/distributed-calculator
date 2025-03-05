package main

import (
	"fmt"
	"net/http"
	"distributed-calculator/internal/orchestrator"
)

func main() {
	http.HandleFunc("/api/v1/calculate", orchestrator.HandleCalculate)
	http.HandleFunc("/api/v1/expressions", orchestrator.HandleGetExpressions)
	http.HandleFunc("/api/v1/expressions/", orchestrator.HandleGetExpressionByID)
	http.HandleFunc("/internal/task", orchestrator.HandleTask)

	fmt.Println("Orchestrator started at :8080")
	http.ListenAndServe(":8080", nil)
}
