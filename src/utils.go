package main

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
