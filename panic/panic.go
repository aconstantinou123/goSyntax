package main

import (
	"net/http"
)

func main() {
	// simple web handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// panics if error
		panic(err.Error())
	}
}
