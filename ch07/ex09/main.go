package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch07/ex08/track_sort"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// リクエストをもとにtracksをソートする処理
		printTracks(w, tracks)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

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

func printTracks(w io.Writer, tracks []*track_sort.Track) {
	if err := trackList.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}

var trackList = template.Must(template.New("track list").
	Parse(`
	<h1>Track List</h1>
	<table>
	<tr style='text-align: left'>
	  <th>Title</th>
	  <th>Artist</th>
	  <th>Album</th>
	  <th>Year</th>
	  <th>Length</th>
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
