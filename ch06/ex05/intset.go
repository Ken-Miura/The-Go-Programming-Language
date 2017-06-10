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
		num += popCountByClearing(s.words[i])
	}
	return num
}

// ch2より引数の型を変更して利用
func popCountByClearing(x uint) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
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
