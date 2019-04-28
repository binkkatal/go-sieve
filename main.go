package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/binkkatal/go-sieve/seive"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Program requires one argument.", os.Args)
		os.Exit(0)
	}
	arg := os.Args[1]
	input, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatal(err)
	}
	s := seive.NewPrimeSeive(input)
	s.PrintTable()
}
