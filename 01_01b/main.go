package main

import (
	"flag"
	"log"

	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {

	t, err := time.Parse(expectedFormat, target)
	log.Println(t)
	if err != nil || time.Now().After(t) {
		log.Fatal(err)
	}

	return t
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	t := time.Until(target).Hours() / 24

	return t
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
