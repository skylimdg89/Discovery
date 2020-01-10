package main

import (
	"fmt"
	"os"
)

//var(
//	name, course, module = "DK", "Docker Deep Dive", 3.2
//)
func main() {
	name := os.Getenv("USER")
	course := "Docker Deep Dive"
	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourse(&course)
	fmt.Println("\nYou are now watching course", course)

}

func changeCourse(course *string) string {

	*course = "First Look: Native Docker Clustering"

	fmt.Println("\nTrying to change your course to", *course)

	return *course
}
