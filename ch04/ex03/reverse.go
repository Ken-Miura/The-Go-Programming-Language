// Copyright 2017 Ken Mirua
package ex03

// TODO func Reverse(s *[...]int) では駄目？？？関数の定義のコンパイルは通るが、実際に関数を使うコードは、実引数を与える部分でコンパイルエラーになる？？？なんで？？？
func Reverse(s *[10]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
