package main

import "math"

//constants
const start = 97
const end = 122
const passwordmaxlength = 2
const hashToCrack = "$6$hpKW0d9oSGIXHJpn$ItrRNxxgJE.sYixsZtfrDJyYj9XMheFEcyD1ybc/4gJMrICcchyU/D1gYMN7gQKuA3ZDNuqRWbWm37k3zTyvG1"

//args [0] = minlength , [1] = max length
func get_possibilites() int {

	args := []int{1, passwordmaxlength}

	sum := 0.0

	for i := args[0]; i <= args[1]; i++ {
		sum += math.Pow(float64(start-end+2), float64(i))
	}

	return int(sum)
}
