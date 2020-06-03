package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type River struct {
	Name        string `json:"Name"`
	Desc        string `json:"desc"`
	CurrentFlow int    `json:"currentFlow"`
}

type Rivers []River

func allRivers(w http.ResponseWriter, r *http.Request) {
	rivers := Rivers{
		River{Name: "Colorado River", Desc: "The Colorado River is 2000 Miles in length", CurrentFlow: 800},
	}
	fmt.Println("Endpoint Hit: All Rivers Endpoint")
	json.NewEncoder(w).Encode(rivers)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/allrivers", allRivers)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
