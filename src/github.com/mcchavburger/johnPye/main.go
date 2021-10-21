package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var term string = "intel"

func main() {
	url := "https://www.johnpyeauctions.co.uk/search-auctions.asp?term=" + term
	resp, err := http.Get(url)
	//asign the response and err variables from the http.get(url) command
	if err != nil {
		//if error isnt empty
		fmt.Fprintf(os.Stderr, "fecth: %v\n", err)
		os.Exit(1)
	}
	out, err := os.Create(term + ".txt")
	if err != nil {
		//panic
	}
	io.Copy(out, resp.Body)
}
