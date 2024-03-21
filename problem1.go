// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var input, counter int

	// this will store diagonal number total of matrix with N rows
	// with N as key and diagonal number total as value
	mapN := make(map[int]int)

	// insert the smallest possible N
	mapN[1] = 1
	mapN[2] = 10

	// receive input
	fmt.Println("Please input N value")
	fmt.Scan(&input)

	if input%2 == 0 {
		counter = 4
	} else {
		counter = 3
	}

	// loop until counter reach input value
	for ; counter <= input; counter = counter + 2 {
		// store the value of diagonal number total in map
		mapN[counter] = mapN[counter-2] + 10*(counter-1) + 4*(counter-2)*(counter-2)
	}

	// print the diagonal number total
	fmt.Println(mapN[input])
}
