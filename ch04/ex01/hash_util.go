// Copyright 2017 Ken Mirua
package main

func CountDifferencesBetweenTwoSHA256Messages(x, y [32]byte) int {
	count := 0
	for i := range x {
		result := x[i] ^ y[i]
		count += bitCount(uint64(result))
	}
	return count
}

// サンプルコードより
func bitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}
