package iteration

func Repeat(s string, t int) string {
	if t <= 0 {
		return s
	}
	var r string
	for i := 0; i < t; i++ {
		r += s
	}
	return r
}
