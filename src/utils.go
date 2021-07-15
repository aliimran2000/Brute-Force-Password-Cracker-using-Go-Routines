package main

import "math"

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

//returns the string at count i of base
func get_pass_at(i int) string {
	var str string
	var index = i
	const basn = base

	for index >= 0 {

		mod := byte(index%basn + start)
		index = (index / basn) - 1
		str = string(mod) + str
	}

	return str
}

//return starting integer for starting length from base
func get_starting_int(lens int) float64 {
	sum := 0.0
	for i := 0; i < lens; i++ {
		sum += math.Pow(float64(base), float64(i))
	}

	return sum

}

func get_possibilites() uint64 {

	args := []int{passwordminlength, passwordmaxlength}

	var sum uint64

	for i := args[0]; i <= args[1]; i++ {
		sum += uint64(math.Pow(float64(base+1), float64(i)))
	}

	return sum
}
