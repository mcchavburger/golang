//prints command line arguements
//example "go run .\basicLoop.go "test" "Something Else" " would print the inputs
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	//fmt.Println("Total number of inputs recieved: " + string(len(os.Args)))
	//Above didnt work as string(int) converts to a rune
	fmt.Println("Total number of inputs recieved: " + strconv.Itoa(len(os.Args)))
	//i starts as 1, because element 0 contains the go executable
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("I Value is: " + strconv.Itoa(i) + " with value: " + os.Args[i])
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
