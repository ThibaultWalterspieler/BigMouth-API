package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var (
	ctx context.Context 
	sa option.ClientOption 
	app *firebase.App
	err error
	client *firestore.Client

)

func initFirestore()  {
	ctx = context.Background()
	sa = option.WithCredentialsFile("../firebase.json")
	app, err = firebase.NewApp(ctx, nil, sa)
	if err != nil {
  		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
  		log.Fatalln(err)
	}
}

func main() {
	initFirestore()
	fmt.Println("Big Mouth API v1.0")
	Characters = []Character{}
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the Big Mouth API 👄")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    // New instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc(apiURL, homePage)
	myRouter.HandleFunc(fmt.Sprintf("%s/characters", apiURL), returnAllCharacters)
	myRouter.HandleFunc(fmt.Sprintf("%s/character", apiURL), createNewCharacter).Methods("POST")
    myRouter.HandleFunc(fmt.Sprintf("%s/characters/{id}", apiURL), returnSingleCharacter)
    log.Fatal(http.ListenAndServe(":667", myRouter))
}

func returnAllCharacters(w http.ResponseWriter, r *http.Request){
	iter := client.Collection("characters").Documents(ctx)
	for {
        doc, err := iter.Next()
        if err == iterator.Done {
                break
        }
        if err != nil {
                log.Fatalf("Failed to iterate: %v", err)
        }
		fmt.Println("Endpoint Hit: returnAllArticles")
		json.NewEncoder(w).Encode(doc.Data())
	}
}

func returnSingleCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	
	dsnap,err := client.Collection("characters").Doc(key).Get(ctx)
	
	characterData := dsnap.Data()
	json.NewEncoder(w).Encode(characterData)

	if err != nil {
		return
	}
}

func createNewCharacter(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)


	_, err := client.Collection("characters").Doc("id").Set(ctx, reqBody )
if err != nil {
        // Handle any errors in an appropriate way, such as returning them.
        log.Printf("An error has occurred: %s", reqBody)
}

	fmt.Fprintf(w, "%+v", string(reqBody))
}

const apiURL string = "/api/v1"

// Date type
type Date struct {
	Year int
	Month time.Month
	Day time.Weekday
}

// HormoneMonster type
type HormoneMonster struct {
	current string
	formers []string
}

// Timestamp type
type Timestamp time.Time

// Character type 
type Character struct {
	ID string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age uint16 `json:"age"`
	Aliases []string `json:"aliases"`
	Status string `json:"status"`
	Birthdate *Timestamp `json:"Birthdate"`
	Gender string `json:"gender"`
	SexualOrientation string `json:"sexual_orienation"`
	Religion string `json:"religion"`
	Occupation string `json:"occupation"`
	HormoneMonster HormoneMonster `json:"hormone_monster"`
}

// Characters array 
var Characters []Character;