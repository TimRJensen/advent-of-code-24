package util

func Atoi(buff []byte) int {
	r := 0
	for _, b := range buff {
		r = r<<1 + r<<3
		r += int(b) - 48
	}
	return r
}
