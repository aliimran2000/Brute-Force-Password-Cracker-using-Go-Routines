package main

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

func adder(a, b byte) (byte, byte) {
	ans := b%byte(end-start+1) + a
	if ans > byte(end) {
		return byte(a - (ans - byte(end))), (26 - (ans - byte(end)))
	}

	return byte(ans), (b % byte(end-start+1))

}

func ByteAddtion(b []byte, add int) string {

	carry := byte(add)

	for i := len(b) - 1; i > -1; i-- {
		by, c := adder(b[i], carry)
		b[i] = byte(by)
		fmt.Println(c)
		carry -= c
	}

	return string(b)

}
