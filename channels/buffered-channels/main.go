package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "These are the times that try men's souls\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}

}

/*
Step-1 Unbuffered channel with printing phrase
package main

import (
	"strings"
	"fmt"
)

func main() {
	phrase := "These are the times that try men's souls\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string)

	for _, word := range words {
		ch <- word
	}

	for i:=0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}

}
*/

/*
Step-2 Add buffered channel
package main

import (
	"strings"
	"fmt"
)

func main() {
	phrase := "These are the times that try men's souls\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	for i:=0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}

}
*/
