// Copyright 2017 Ken Mirua
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{mapping: map[string]dollars{"shoes": 50, "socks": 5}}
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
	mapping map[string]dollars
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	db.Lock()
	defer db.Unlock()
	for item, price := range db.mapping {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db.Lock()
	defer db.Unlock()
	if price, ok := db.mapping[item]; ok {
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
	if _, ok := db.mapping[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item: %q has already existed\n", item)
	} else {
		db.mapping[item] = dollars(price)
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
	if price, ok := db.mapping[item]; ok {
		db.mapping[item] = dollars(newPrice)
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
	if _, ok := db.mapping[item]; ok {
		delete(db.mapping, item)
		fmt.Fprintf(w, "%s was deleted\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
