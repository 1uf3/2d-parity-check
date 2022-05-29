package parityc

import (
	"errors"
	"sync"
)

var ErrMultiMiss = errors.New("Multiple bit missed")
var ErrIsNotBit = errors.New("it is not 1bit")

func ParityCheck2D(b [][]int) ([][]int, error) {
	dx, dy := get2Ddiv(b)
	x, y, err := evenCheck(b, dx, dy)
	if err != nil {
		return nil, err
	}
	if x < 0 {
		return b, nil
	}
	// 	if x != y {
	// 		return nil, errors.New("Miss Detected! but cannot correction this.")
	// 	}
	cb, err := correction(b, x, y)
	if err != nil {
		return nil, err
	}
	return cb, nil
}

func get2Ddiv(b [][]int) (dx, dy int) {
	dy = len(b)
	dx = len(b[0])
	return
}

func evenCheck(b [][]int, dx, dy int) (x, y int, err error) {
	wg := &sync.WaitGroup{}

	missX, missY := 0, 0
	x, y = 0, 0
	wg.Add(1)
	go func() {
		for i := 0; i < dy-1; i++ {
			l := 0
			for _, vv := range b[i] {
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
		for i := 0; i < dx-1; i++ {
			c := 0
			for _, vv := range b {
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

	if missX == 0 && missY == 0 {
		return -1, 0, nil
	}

	return
}

func oddCheck(b [][]int, dx, dy int) (x, y int, err error) {
	wg := &sync.WaitGroup{}

	missX := 0
	missY := 0
	wg.Add(1)
	go func() {
		for i := 0; i < dy-1; i++ {
			l := 0
			for _, vv := range b[i] {
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
		for i := 0; i < dx-1; i++ {
			c := 0
			for _, vv := range b {
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

	if missX == 0 && missY == 0 {
		return -1, 0, nil
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
