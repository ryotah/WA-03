package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ryotah/WA-03/roman"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number")
		os.Exit(1)
	}

	s := os.Args[1]
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r, err := roman.ToRoman(i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(r)
}
