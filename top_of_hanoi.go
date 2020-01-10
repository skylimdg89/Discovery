package main

import "fmt"

func hanoi(n int) {
	fmt.Println("number of disks: ", n)
	move(n, 1, 2, 3)
}

func move(n, from, to, via int) {
	if n <= 0 {
		return
	}
	move(n-1, from, via, to)
	fmt.Println(from, "->", to)
	move(n-1, via, to, from)

}

func main() {
	fmt.Println("Top of Hanoi")
	hanoi(3)
}
