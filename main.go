package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("listen on http:localhost:8080")
	http.HandleFunc("/", func(res http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(res, "testing")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
