package main

import (
	"fmt"
	"math"
	"sync"
)

//cap 65-90
//small 97-122

const start = 97
const end = 122
const search = "baaaabc"

var possibilities float64 = math.Pow(float64(len(search)), float64(122-97))

var wg sync.WaitGroup

func increment(s *string) {

	b := []byte(*s)
	var carry int = len(b)
	for i := len(b) - 1; i > -1; i-- {
		if b[i] < end {
			b[i]++
			break
		} else {
			b[i] = start
			carry--
		}
	}
	if carry == 0 {
		b = append(b, start)
	}
	*s = string(b)
}

func Crack(s string, iters int, c chan int) {
	defer wg.Done()

	var breaksearch bool = false

	var i int = 0
	for !breaksearch {

		select {
		case <-c:
			fmt.Println("Ending Search at ", i)
			breaksearch = true
		default:
			increment(&s)
			if s == search {
				fmt.Println("Found the value at ", i)
				c <- 1
				breaksearch = true
			}
		}
		i++

	}

	fmt.Println("ending search thread", i-1)
}

func main() {
	var s1 string = "aaaaaab"
	var s2 string = "baaaaac"

	recieve := make(chan int)

	wg.Add(1)
	go Crack(s2, 100000, recieve)
	wg.Add(1)
	go Crack(s1, 100000, recieve)

	fmt.Println("ending search")
	wg.Wait()
}
