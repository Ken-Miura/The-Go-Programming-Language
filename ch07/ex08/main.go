// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byLength []*Track

func (x byLength) Len() int           { return len(x) }
func (x byLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type CustomSort struct {
	SortKeys []string
	t        []*Track
}

func (x CustomSort) Len() int           { return len(x.t) }
func (x CustomSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
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
func (x CustomSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	fmt.Println("Custom:")
	fmt.Println("Primary: Title, Secondary: Year, Tertiary: Length")
	sort.Sort(CustomSort{[]string{"Length", "Year", "Title"}, tracks})
	printTracks(tracks)

	// sort.Stableは、一つ前のソート結果を保ちつつ、新しく与えられたキーでソートするように見える。
	// 下記の出力は、一次キー: Year, 二次キー: Lengthの結果となった。
	fmt.Println("\nStable Sort:")
	fmt.Println("sort.Stable sort by Title at first, next by Year and finally by Length")
	sort.Stable(byTitle(tracks))
	sort.Stable(byYear(tracks))
	sort.Stable(byLength(tracks))
	printTracks(tracks)
}
