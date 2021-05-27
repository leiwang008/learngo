package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"example.com/morestrings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	//log.SetFlags(0)

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())

	//define a slice of names
	names := []string{"Tom", "Sandy", "Jack"}
	msgs, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(msgs)
		for _, name := range names {
			fmt.Println(morestrings.ReverseRunes(msgs[name]))
		}
	}
}
