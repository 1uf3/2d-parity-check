package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

var ErrMultiMiss = errors.New("Multiple bit missed")
var ErrIsNotBit = errors.New("it is not 1bit")

func main() {
	b := [][]int{
		{1, 0, 0, 0, 1},
		{1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0},
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1},
	}

	jb, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}

	dx, dy := get2Ddiv(b)
	x, y, err := evenCheck(b, dx, dy)
	if err != nil {
		log.Fatal(err)
	}
	// 	if x != y {
	// 		log.Println("Miss Detected! but cannot correction this.")
	// 	}
	cb, err := correction(b, x, y)
	if err != nil {
		log.Fatal(err)
	}

	ja, err := json.Marshal(cb)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", jb)
	fmt.Printf("%s\n", ja)
}

func get2Ddiv(b [][]int) (dx, dy int) {
	dy = len(b)
	dx = len(b[0])
	return
}

func evenCheck(b [][]int, dx, dy int) (x, y int, err error) {
	wg := &sync.WaitGroup{}

	missX := 0
	missY := 0
	wg.Add(1)
	go func() {
		for i, v := range b {
			l := 0
			for _, vv := range v {
				l += vv
			}
			if l%2 != 0 {
				missX++
				x = i
			}
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < dx; i++ {
			c := 0
			for ii, vv := range b {
				if i == dx-1 && ii == dy-1 {
					break
				}
				c += vv[i]
			}
			if c%2 != 0 {
				missY++
				y = i
			}
		}
		wg.Done()
	}()
	wg.Wait()

	if missX > 1 && missY > 1 {
		return x, y, ErrMultiMiss
	}

	return
}

func oddCheck(b [][]int, dx, dy int) (x, y int, err error) {
	wg := &sync.WaitGroup{}

	missX := 0
	missY := 0
	wg.Add(1)
	go func() {
		for i, v := range b {
			l := 0
			for _, vv := range v {
				l += vv
			}
			if l%2 == 0 {
				missX++
				x = i
			}
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < dx; i++ {
			c := 0
			for ii, vv := range b {
				if i == dx-1 && ii == dy-1 {
					continue
				}
				c += vv[i]
			}
			if c%2 == 0 {
				missY++
				y = i
			}
		}
		wg.Done()
	}()
	wg.Wait()

	if missX > 1 && missY > 1 {
		return x, y, ErrMultiMiss
	}

	return
}

func correction(b [][]int, x, y int) ([][]int, error) {
	if b[x][y] == 1 {
		b[x][y] = 0
	} else {
		b[x][y] = 1
	}
	return b, nil
}
