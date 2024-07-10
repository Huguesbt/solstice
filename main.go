package main

import (
	"flag"
	"fmt"
	"github.com/HuguesBt/solstice/pkg/solstice"
	"time"
)

var (
	action        string
	year          int
	dateFormatStr = "2006-01-02"
)

func initFlags() {
	flag.IntVar(&year, "y", time.Now().Year(), "Year")
	flag.Parse()
}

func main() {
	initFlags()

	if year == 0 {
		fmt.Println("You need to specify a year")
	} else {
		fmt.Printf("For year %d\n", year)
		fmt.Printf("Spring equinox %s\n", solstice.Spring(year).Format(dateFormatStr))
		fmt.Printf("Summer solstice %s\n", solstice.Summer(year).Format(dateFormatStr))
		fmt.Printf("Fall equinox %s\n", solstice.Fall(year).Format(dateFormatStr))
		fmt.Printf("Winter solstice %s\n", solstice.Winter(year).Format(dateFormatStr))
	}
}
