// Copyright 2017 Ken Miura
package ex05

import (
	"bytes"
	"fmt"
)

const numOfBits = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < numOfBits; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, numOfBits*i+j)
			}
		}
	}
	return elems
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) AddAll(integers ...int) {
	for _, v := range integers {
		s.Add(v)
	}
}

func (s *IntSet) Len() int {
	num := 0
	if s == nil {
		return num
	}
	for i := range s.words {
		num += bitCount(s.words[i])
	}
	return num
}

// ch2のサンプルコードより
func bitCount(x uint) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func (s *IntSet) Remove(x int) {
	if s == nil {
		return
	}
	word, bit := x/numOfBits, uint(x%numOfBits)
	for word >= len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	if s == nil {
		return
	}
	for i := range s.words {
		s.words[i] &^= 63
	}
}

func (s *IntSet) Copy() *IntSet {
	if s == nil {
		return &IntSet{}
	}
	copied := &IntSet{words: make([]uint, len(s.words), cap(s.words))}
	for i := range s.words {
		copied.words[i] = s.words[i]
	}
	return copied
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/numOfBits, uint(x%numOfBits)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/numOfBits, uint(x%numOfBits)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < numOfBits; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", numOfBits*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
