package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

type Node struct {
	Next  *Node
	Value interface{}
}

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Greyland", "William", "Danny"}
	messages, err := greetings.Hello2(names)

	// Request a greeting message
	// message, err := greetings.Hello("Greyland")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	// var first *Node
	// visited := make(map[*Node]bool)
	// for n := first; n != nil; n = n.Next {
	// 	if visited[n] {
	// 		fmt.Println("cycle detected")
	// 		break
	// 	}
	// 	visited[n] = true
	// 	fmt.Println(n.Value)
	// }
}
