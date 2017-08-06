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

func randomNotPalindrome(rng *rand.Rand) string {
	n := rng.Intn(23)
	n += 2 // 2以上24以下の長さ. 長さ0と1は文字列に関係なく回文となるので除外。
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}

	// 回文の前半分を現在の文字とは異なる文字に変更
	for i := 0; i < (n+1)/2; i++ {
		for {
			r := rune(rng.Intn(0x1000))                                                   // random rune up to '\u0999'
			if unicode.IsLetter(r) && (unicode.ToLower(runes[i]) != unicode.ToLower(r)) { // 回文判定は、文字以外は無視する、かつ大文字小文字の違いも無視する
				runes[i] = r
				break
			}
		}
	}
	// 回文の後ろ半分がすべて非文字のとき、後ろ半分が無視され、前半分だけが考慮されて回文になりえる。なので後ろ半分がすべて非文字だったら修正
	allNonLetter := true
	for i := 0; i < (n+1)/2; i++ {
		allNonLetter = allNonLetter && !unicode.IsLetter(runes[n-1-i])
	}
	if allNonLetter {
		i := rng.Intn((n + 1) / 2)
		for {
			r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
			if unicode.IsLetter(r) {    // 文字以外は代入しても無視されるので文字を入れる。
				runes[n-1-i] = r
				break
			}
		}
	}
	return string(runes)
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
