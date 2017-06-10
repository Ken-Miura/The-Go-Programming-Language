// Copyright 2017 Ken Miura
package ex03

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
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

func NewIntSet(integers ...int) *IntSet {
	pIntSet := new(IntSet)
	for _, integer := range integers {
		pIntSet.Add(integer)
	}
	return pIntSet
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

// ch2のサンプルコードより
func popCountByClearing(x uint64) int {
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
	word, bit := x/64, uint(x%64)
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
	copied := &IntSet{words: make([]uint64, len(s.words), cap(s.words))}
	for i := range s.words {
		copied.words[i] = s.words[i]
	}
	return copied
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
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
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
