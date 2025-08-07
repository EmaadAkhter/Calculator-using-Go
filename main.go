package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type eqn struct {
	Num1    int    `json:"num1"`
	Num2    int    `json:"num2"`
	Oprator string `json:"oprator"`
}

type welcome string

func (wc welcome) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Welcome to the go web servers")
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	var eq eqn
	err := json.NewDecoder(r.Body).Decode(&eq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	a := eq.Num1
	b := eq.Num2
	c := eq.Oprator
	var result int
	var id string
	if c == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}
	switch c {
	case "+":
		result = a + b
		id = "Addition"
	case "-":
		result = a - b
		id = "Subtraction"
	case "*":
		result = a * b
		id = "Multiplication"
	case "/":
		if b == 0 {
			http.Error(w, "Division by zero", http.StatusBadRequest)
			return
		}
		result = a / b
		id = "Division"
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "The Result of %s is:%d", id, result)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func main() {
	router := http.NewServeMux()
	var wc welcome
	router.Handle("/", wc)
	router.Handle("/hello", http.HandlerFunc(helloHandler))
	router.Handle("/calc", http.HandlerFunc(CalcHandler))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("server is running on port 8080")
	server.ListenAndServe()
}
