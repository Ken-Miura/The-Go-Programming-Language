// Copyright 2017 Ken Mirua
package ex07

// TODO 追加のメモリ割り当てなしのバージョン。UTF-8が可変長エンコーディングなせいで多分無理そうな気がする。
func Reverse(s []byte) []byte {
	runes := []rune(string(s))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return []byte(string(runes))
}
