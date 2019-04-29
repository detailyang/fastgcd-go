// +build !amd64

package fastgcd

func ctz(x uint64) uint64 {
	// This uses a binary search algorithm from Hacker's Delight.
	n := uint64(1)
	if (x & 0x0000FFFF) == 0 {
		n = n + 16
		x = x >> 16
	}

	if (x & 0x000000FF) == 0 {
		n = n + 8
		x = x >> 8
	}

	if (x & 0x0000000F) == 0 {
		n = n + 4
		x = x >> 4
	}

	if (x & 0x00000003) == 0 {
		n = n + 2
		x = x >> 2
	}

	return n - (x & 1)
}
