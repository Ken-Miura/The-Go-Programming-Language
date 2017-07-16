// Copyright 2017 Ken Miura
package track_sort

import (
	"sync"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type ByTitle []*Track

func (x ByTitle) Len() int           { return len(x) }
func (x ByTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x ByTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ByYear []*Track

func (x ByYear) Len() int           { return len(x) }
func (x ByYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x ByYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ByLength []*Track

func (x ByLength) Len() int           { return len(x) }
func (x ByLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x ByLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// SortKeysは、添え字の大きい要素ほど優先度の高いキーを意味する。
type CustomSort struct {
	SortKeys []string
	T        []*Track
	sync.Mutex
}

func (x CustomSort) Len() int           { return len(x.T) }
func (x CustomSort) Less(i, j int) bool { return x.less(x.T[i], x.T[j]) }
func (x CustomSort) less(a, b *Track) bool {
	for i := len(x.SortKeys) - 1; i > -1; i-- {
		switch x.SortKeys[i] {
		case "Title":
			if a.Title != b.Title {
				return a.Title < b.Title
			}
		case "Artist":
			if a.Artist != b.Artist {
				return a.Artist < b.Artist
			}
		case "Album":
			if a.Album != b.Album {
				return a.Album < b.Album
			}
		case "Year":
			if a.Year != b.Year {
				return a.Year < b.Year
			}
		case "Length":
			if a.Length != b.Length {
				return a.Length < b.Length
			}
		default:
			panic("This line must not be reached.")
		}
	}
	return false
}
func (x CustomSort) Swap(i, j int) { x.T[i], x.T[j] = x.T[j], x.T[i] }
