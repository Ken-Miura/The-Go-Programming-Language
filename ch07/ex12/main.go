// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{Mapping: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer, "usage")
		fmt.Fprintln(writer, "http://localhost:8000/list")
		fmt.Fprintln(writer, "http://localhost:8000/price?item='item name'")
		fmt.Fprintln(writer, "http://localhost:8000/create?item='item name'&price='price'")
		fmt.Fprintln(writer, "http://localhost:8000/update?item='item name'&price='price'")
		fmt.Fprintln(writer, "http://localhost:8000/delete?item='item name'")
	})
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	sync.Mutex
	Mapping map[string]dollars
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	db.Lock()
	defer db.Unlock()
	if err := itemList.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

var itemList = template.Must(template.New("item list").
	Funcs(template.FuncMap{"size": size}).
	Parse(`
	<h1>{{.Mapping | size}} items</h1>
	<table>
	<tr style='text-align: left'>
	  <th>item</th>
	  <th>price</th>
	</tr>
	{{range $key, $value := .Mapping}}
	<tr>
	  <td>{{$key}}</td>
	  <td>{{$value}}</td>
	</tr>
	{{end}}
	</table>
	`))

func size(mapping map[string]dollars) int {
	return len(mapping)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db.Lock()
	defer db.Unlock()
	if price, ok := db.Mapping[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStrings, ok := req.URL.Query()["price"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "query 'price' was not specified")
		return
	}
	price, err := strconv.ParseFloat(priceStrings[0], 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "parsing query (price) as dollors: %v\n", err)
		return
	}
	db.Lock()
	defer db.Unlock()
	if _, ok := db.Mapping[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item: %q has already existed\n", item)
	} else {
		db.Mapping[item] = dollars(price)
		fmt.Fprintf(w, "%s was created and price of that was set to %g\n", item, price)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	newPriceStrings, ok := req.URL.Query()["price"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "query 'price' was not specified")
		return
	}
	newPrice, err := strconv.ParseFloat(newPriceStrings[0], 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "parsing query (price) as dollors: %v\n", err)
		return
	}
	db.Lock()
	defer db.Unlock()
	if price, ok := db.Mapping[item]; ok {
		db.Mapping[item] = dollars(newPrice)
		fmt.Fprintf(w, "price of %s was updated from %g to %g\n", item, price, newPrice)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db.Lock()
	defer db.Unlock()
	if _, ok := db.Mapping[item]; ok {
		delete(db.Mapping, item)
		fmt.Fprintf(w, "%s was deleted\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
