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
	http.HandleFunc("/sort_by_title", sortByTitle)
	http.HandleFunc("/sort_by_artist", sortByArtist)
	http.HandleFunc("/sort_by_album", sortByAlbum)
	http.HandleFunc("/sort_by_year", sortByYear)
	http.HandleFunc("/sort_by_length", sortByLength)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func sortByTitle(w http.ResponseWriter, r *http.Request) {
	sortBy("Title", w, r)
}

func sortByArtist(w http.ResponseWriter, r *http.Request) {
	sortBy("Artist", w, r)
}

func sortByAlbum(w http.ResponseWriter, r *http.Request) {
	sortBy("Album", w, r)
}

func sortByYear(w http.ResponseWriter, r *http.Request) {
	sortBy("Year", w, r)
}

func sortByLength(w http.ResponseWriter, r *http.Request) {
	sortBy("Length", w, r)
}

func sortBy(key string, w http.ResponseWriter, _ *http.Request) {
	customSortTracks.Lock()
	defer customSortTracks.Unlock()
	for i, sortKey := range customSortTracks.SortKeys {
		if sortKey == key {
			remove(customSortTracks.SortKeys, i)
		}
	}
	customSortTracks.SortKeys = append(customSortTracks.SortKeys, key)
	sort.Sort(customSortTracks)
	printTracks(w, customSortTracks)
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
