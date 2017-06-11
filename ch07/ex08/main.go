// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch07/ex08/track_sort"
)

var tracks = []*track_sort.Track{
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

func printTracks(tracks []*track_sort.Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	fmt.Println("Custom:")
	fmt.Println("Primary: Title, Secondary: Year, Tertiary: Length")
	sort.Sort(track_sort.CustomSort{[]string{"Length", "Year", "Title"}, tracks})
	printTracks(tracks)

	// sort.Stableは、一つ前のソート結果を保ちつつ、新しく与えられたキーでソートするように見える。
	// 下記の出力は、一次キー: Year, 二次キー: Lengthの結果となった。
	fmt.Println("\nStable Sort:")
	fmt.Println("sort.Stable sort by Title at first, next by Year and finally by Length")
	sort.Stable(track_sort.ByTitle(tracks))
	sort.Stable(track_sort.ByYear(tracks))
	sort.Stable(track_sort.ByLength(tracks))
	printTracks(tracks)
}
