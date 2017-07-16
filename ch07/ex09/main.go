package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch07/ex08/track_sort"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		customSortTracks.Lock()
		defer customSortTracks.Unlock()
		printTracks(w, customSortTracks)
	})
	http.HandleFunc("/sort_by_title", func(w http.ResponseWriter, request *http.Request) {
		customSortTracks.Lock()
		defer customSortTracks.Unlock()
		sortBy("Title")
		printTracks(w, customSortTracks)
	})
	http.HandleFunc("/sort_by_artist", func(w http.ResponseWriter, request *http.Request) {
		customSortTracks.Lock()
		defer customSortTracks.Unlock()
		sortBy("Artist")
		printTracks(w, customSortTracks)
	})
	http.HandleFunc("/sort_by_album", func(w http.ResponseWriter, request *http.Request) {
		customSortTracks.Lock()
		defer customSortTracks.Unlock()
		sortBy("Album")
		printTracks(w, customSortTracks)
	})
	http.HandleFunc("/sort_by_year", func(w http.ResponseWriter, request *http.Request) {
		customSortTracks.Lock()
		defer customSortTracks.Unlock()
		sortBy("Year")
		printTracks(w, customSortTracks)
	})
	http.HandleFunc("/sort_by_length", func(w http.ResponseWriter, request *http.Request) {
		customSortTracks.Lock()
		defer customSortTracks.Unlock()
		sortBy("Length")
		printTracks(w, customSortTracks)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func sortBy(key string) {
	for i, sortKey := range customSortTracks.SortKeys {
		if sortKey == key {
			remove(customSortTracks.SortKeys, i)
		}
	}
	customSortTracks.SortKeys = append(customSortTracks.SortKeys, key)
	sort.Sort(customSortTracks)
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func printTracks(w io.Writer, customSortTracks track_sort.CustomSort) {
	if err := trackList.Execute(w, customSortTracks.T); err != nil {
		log.Fatal(err)
	}
}

var trackList = template.Must(template.New("track list").
	Parse(`
	<h1>Track List</h1>
	<table>
	<tr style='text-align: left'>
	  <th><a href="sort_by_title">Title</a></th>
	  <th><a href="sort_by_artist">Artist</a></th>
	  <th><a href="sort_by_album">Album</a></th>
	  <th><a href="sort_by_year">Year</a></th>
	  <th><a href="sort_by_length">Length</a></th>
	</tr>
	{{range $value := .}}
	<tr>
	  <td>{{$value.Title}}</td>
	  <td>{{$value.Artist}}</td>
	  <td>{{$value.Album}}</td>
	  <td>{{$value.Year}}</td>
	  <td>{{$value.Length}}</td>
	</tr>
	{{end}}
	</table>
	`))

var customSortTracks = track_sort.CustomSort{[]string{}, tracks, sync.Mutex{}}

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
