package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "suchi", Age: 28}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Print("error")
		return
	}

	fmt.Print("serving")
	fileserve := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserve)

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		http.ServeFile(w, r, "./static/form.html")
	})
	http.HandleFunc("/form/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(data))
	})
	http.ListenAndServe(":8000", nil)
}
