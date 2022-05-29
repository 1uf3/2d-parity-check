package main

import (
	"fmt"
	"log"

	"main/parityc"
)

func main() {
	bb := [][]int{
		{1, 0, 0, 0, 1},
		{1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1},
	}
	fmt.Println(bb)

	ba, err := parityc.ParityCheck2D(bb)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ba)
}
