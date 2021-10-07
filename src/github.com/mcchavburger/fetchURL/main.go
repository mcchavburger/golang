package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		//instantiate for loop against the index, assign the index of the os.Args[1:] array to _ (blank identifier) and the value to url
		resp, err := http.Get(url)
		//asign the response and err variables from the http.get(url) command
		if err != nil {
			//if error isnt empty
			fmt.Fprintf(os.Stderr, "fecth: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		//read the contents of the response body and assign to the variable b, if error assign to err
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
