package mathutils

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	} else {
		return GCD(b, a%b)
	}
}
