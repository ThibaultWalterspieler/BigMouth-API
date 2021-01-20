package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
    log.Fatal(http.ListenAndServe(":9302", nil))
}

func main() {
	Articles = []Article{
		{Name:"Pimiento", Species:"Capsicum annuum", ScovilleScale:[2]int{100,500}},
		{Name:"Ghost pepper", Species:"Unknown", ScovilleScale:[2]int{100}},
	}
	handleRequests()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

// Article type 
type Article struct {
	Name string `json:"Name"`
	Species string `json:"Species"`
	ScovilleScale [2]int `json:"ScovilleScale"`
}

// Articles array 
var Articles []Article;