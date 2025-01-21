package main

import (
	"fmt"
	greetings "github.com/ladydn/Greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings ")
	log.SetFlags(0)

	names := []string{"Lady", "Dn", "Golang"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// Run the Hello function from the greetings package
	//message, err := greetings.Hello("Lady")
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(messages)
}
