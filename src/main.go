package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
)

var wg sync.WaitGroup

func Crack(s string, iters int, c chan int, r int) {
	defer wg.Done()

	fmt.Println("routine started")
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

		if i%RoutineNotify == 0 {
			log.Println("Routine : ", r, "Active", s)
		}

		i++
		if i >= iters {
			breaksearch = true
		}

	}

	fmt.Println("ending search thread", i-1)
}

func main() {

	c := make(chan int, 1)
	pos := get_possibilites()
	fmt.Println("Strat Cracking through", pos, "Combinations?")

	divide := int(pos / MaxNumofRoutines)
	remain := pos % MaxNumofRoutines
	fmt.Println(divide, remain)

	index := int(get_starting_int(passwordminlength))

	for i := 0; i < MaxNumofRoutines; i++ {
		str := get_pass_at(index)
		if len(str) > passwordmaxlength {
			break
		}
		fmt.Println(str)
		go Crack(str, int(divide), c, i)
		wg.Add(1)
		index += divide
	}
	log.Println("Passwords have been distributed")

	wg.Wait()
	log.Println("Search Concluded :) ")

}
