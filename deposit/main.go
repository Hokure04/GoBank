package deposit

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer, "Pong!")
	})
	fmt.Println("Listening on port 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
