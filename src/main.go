package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
)

var wg sync.WaitGroup

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
	args := os.Args[1:]
	count, _ := strconv.Atoi(args[0])
	threads := get_possibilites() / count
	rem := get_possibilites() % count

	fmt.Println(count, threads, rem) //num of threads to spawn

	// var s1 string = "aaaaa"
	// var s2 string = "qweaa"

	// recieve := make(chan int)

	// wg.Add(1)
	// go Crack(s2, 1000, recieve)
	// wg.Add(1)
	// go Crack(s1, 1000, recieve)

	// fmt.Println("ending search")
	// wg.Wait()

	fmt.Println(get_possibilites())
}
