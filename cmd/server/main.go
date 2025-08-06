package main

import (
	"net/http"
)

type gauge float64
type counter int64

func main() {
	err := http.ListenAndServe("0.0.0.0:8080", http.FileServer(http.Dir(".")))
	if err != nil {
		panic(err)
	}
}
