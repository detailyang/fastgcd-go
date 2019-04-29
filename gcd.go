package fastgcd

// GCD get the greatest common divisor
func GCD(a, b uint64) uint64 {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

// BinaryGCD use Binary GCD Algo
// See https://en.wikipedia.org/wiki/Binary_GCD_algorithm
func BinaryGCD(u, v uint64) uint64 {
	shift := uint64(0)

	if u == 0 {
		return v
	}

	if v == 0 {
		return u
	}

	shift = ctz(u | v)
	u >>= ctz(u)

	for {
		v >>= ctz(v)
		if u > v {
			t := v
			v = u
			u = t
		}
		v = v - u

		if v == 0 {
			break
		}
	}

	return u << shift
}
