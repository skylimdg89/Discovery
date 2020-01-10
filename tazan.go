package main

import "fmt"

func tazan(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println("price is ", i*10, " price is ", (i*10)+10)
	}
}

func main() {
	tazan(5)
}
