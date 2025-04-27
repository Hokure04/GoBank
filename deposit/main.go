package main

import (
	"deposit/operations"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getTransactions(writer http.ResponseWriter, request *http.Request) {
	transactions := operations.GetAllTransactions()
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(transactions)
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
	json.NewEncoder(writer).Encode(newTx)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("GET /transactions", getTransactions)
	router.HandleFunc("POST /transfer", transferMoney)

	fmt.Println("Listening on port 8081")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal(err)
	}
}
