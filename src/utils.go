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

func ByteAddtion(b []byte, add int) string {

	diff := byte(end - start + 1)

	carry := byte(add)

	for i := len(b) - 1; i > -1; i-- {
		if carry == 0 {
			break
		} else if carry >= diff {
			carry -= diff
		} else {
			val := b[i] + carry

			if val > byte(end) {
				b[i] = byte(end)
				carry = val - b[i]
			}

		}
	}

	if carry == 0 {
		return string(b)
	}
	return string(b)

}
