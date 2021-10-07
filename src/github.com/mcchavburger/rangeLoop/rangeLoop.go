//range looop
//prints command line arguements
//example "go run .\rangeLoop.go "test" "Something Else" " would print the inputs
//same cexecution as basicLoop but shorter hand

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	//set the variabls s and sep to empty strings, dont need to define variable as string as "" is the default String value
	//same as var s string = ""
	for _, arg := range os.Args[1:] {
		//range produces a pair of values, the index and the value of the element at that index
		//above we assigning the index to _ and the value at that index to arg (_, arg := range os.Args[1:])
		//_ is the "blank identifier", go does not permit unused local variables, however in this code as we do not need the index, we assign to blank and avoid an error
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
