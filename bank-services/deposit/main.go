package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Hokure04/GoBank/deposit/operations"
)

func getTransactions(writer http.ResponseWriter, request *http.Request) {
	transactions := operations.GetAllTransactions()
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(transactions)
	if err != nil {
		return
	}
}

func transferMoney(writer http.ResponseWriter, request *http.Request) {
	var tx operations.Transaction
	err := json.NewDecoder(request.Body).Decode(&tx)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	newTx, err := operations.Transfer(tx.FromAccount, tx.ToAccount, tx.Amount)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(newTx)
	if err != nil {
		return
	}
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /transactions", getTransactions)
	router.HandleFunc("POST /transfer", transferMoney)

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
