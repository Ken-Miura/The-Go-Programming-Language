// Copyright 2017 Ken Miura
// gopl.io/ch11/word2のテスト内容をコピーしてきてさらにテスト追加
package ex03

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unicode"

	"gopl.io/ch11/word2"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
		{" ,.　、。	\n\t", true}, // 句読点や空白が無視されることを確認
	}
	for _, test := range tests {
		if got := word.IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func randomNotPalindrome(rng *rand.Rand) string {
	n := rng.Intn(23)
	n += 2 // 2以上24以下の長さ. 長さ0と1は文字列に関係なく回文となるので除外。
	runes := make([]rune, n)
	for {
		for i := 0; i < n; i++ {
			r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
			runes[i] = r
		}
		// 全部非文字だったら回文になる。
		if isOnlyNonLetters(runes) {
			continue
		}
		// 前半分か後ろ半分が全部非文字だと回文になっている可能性がある
		if isOnlyNonLetters(runes[0:n/2]) || isOnlyNonLetters(runes[n/2:n]) {
			continue
		}

		palindrome := true
		for i := 0; i < (n+1)/2; i++ {
			if unicode.ToLower(runes[i]) != unicode.ToLower(runes[n-i-1]) {
				palindrome = false
			}
		}
		if palindrome {
			continue
		}
		return string(runes)
	}
}

func isOnlyNonLetters(runes []rune) bool {
	result := true
	for _, r := range runes {
		result = result && !unicode.IsLetter(r)
	}
	return result
}

func TestRandomNotPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)
		if word.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		word.IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func ExampleIsPalindrome() {
	fmt.Println(word.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(word.IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !word.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
