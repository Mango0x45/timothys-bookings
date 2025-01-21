package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("Hello, World")
	http.HandleFunc("/hello", getHello)
	err := http.ListenAndServe(":6969", nil)
	fmt.Print(err)

}

func getHello(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
