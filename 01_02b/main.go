package main

import (
	//"fmt"
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	sp := strings.Split(msg, " ")
	var result string
	for _, c := range sp {
		s := strings.Split(c, "")
		result = ""
		for k, ch := range s {
			for j := k + 1; j > 0; j-- {
				result = result + string(ch)
			}
		}

		print(result)
	}

}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
