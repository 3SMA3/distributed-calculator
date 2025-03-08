package main

import (
	"fmt"
	"net/http"

	"github.com/3SMA3/distributed-calculator/internal/orchestrator"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/calculate", orchestrator.HandleCalculate).Methods("POST")
	r.HandleFunc("/api/v1/expressions", orchestrator.HandleGetExpressions).Methods("GET")
	r.HandleFunc("/api/v1/expressions/{id}", orchestrator.HandleGetExpressionByID).Methods("GET")
	r.HandleFunc("/internal/task", orchestrator.HandleTask).Methods("GET", "POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	fmt.Println("Orchestrator started at :8080")
	http.ListenAndServe(":8080", r)
}
