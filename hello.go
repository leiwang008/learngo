package main

import (
	"fmt"
	"log"

	"example.com/morestrings"
	"github.com/google/go-cmp/cmp"
	"github.com/leiwang008/greetings"
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

	fmt.Println(cmp.Diff("hello world", "Hello World"))
}
