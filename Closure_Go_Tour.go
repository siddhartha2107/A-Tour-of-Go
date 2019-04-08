package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func f() func(int) int {
	temp1, temp2 := 0, 1
	return func(i int) int {
		// temp1, temp2 := 0, 1
		if i == 0 {
			return temp1
		}
		if i == 1 {

			return temp2
		}
		temp2 = temp2 + temp1
		temp1 = temp2 - temp1
		return temp2
	}
}

func main() {
	x := f()
	for i := 0; i < 10; i++ {
		fmt.Println(x(i))
	}
}
