package main

import (
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	http.HandleFunc("/", handleRoot)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
