//duplicate 1
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// creates a new map type variable named counts
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		//for each input we are putting the value in the counts map and using ++ to iterate up through the index
	}
	//Note: ignoring potential erros from input.Err()
	for line, n := range counts {
		//above assigning the variables line and n using range, which will assign the value to the varible line and the index to the variable n
		if n > 1 {
			//checking to see if the index is more that 1, <= 1 means array is empty
			fmt.Printf("%d\t%s\n", n, line)
			//printing out the variables.
		}
	}
}
