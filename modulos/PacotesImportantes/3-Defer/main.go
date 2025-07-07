package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	defer fmt.Println("Primeira Linha")
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
	ExemploDeferChamadaHttp()
}

func ExemploDeferChamadaHttp() {
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	req.Body.Close()
}
