package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
)

//cap 65-90
//small 97-122

const start = 97
const end = 122
const hashToCrack = "$6$hpKW0d9oSGIXHJpn$ItrRNxxgJE.sYixsZtfrDJyYj9XMheFEcyD1ybc/4gJMrICcchyU/D1gYMN7gQKuA3ZDNuqRWbWm37k3zTyvG1"

//const search = "qwert"

//var possibilities float64 = math.Pow(float64(len(search)), float64(122-97))

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
			crypter := sha512_crypt.New()
			hash, err := crypter.Generate([]byte(s), []byte("$6$hpKW0d9oSGIXHJpn$"))
			if err != nil {
				panic(err)
			}

			if strings.Compare(hashToCrack, hash) == 0 {
				fmt.Println("Found the value at ", i)
				fmt.Println("The Password is : ", s)
				c <- 1
				breaksearch = true
			}
		}

		i++
		if i > iters {
			breaksearch = true
		}

	}

	fmt.Println("ending search thread", i-1)
}

func main() {
	var s1 string = "aaaaa"
	var s2 string = "qweaa"

	recieve := make(chan int)

	wg.Add(1)
	go Crack(s2, 1000, recieve)
	wg.Add(1)
	go Crack(s1, 1000, recieve)

	fmt.Println("ending search")
	wg.Wait()

}
