package main

//cap 65-90
//small 97-122

const start = 97
const end = 122


import "fmt"


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

func main() {

	var s string = "zzy"

	for {
		increment(&s)
		fmt.Println(s)
		if s == "zzzz" {
			break
		}

	}
}
