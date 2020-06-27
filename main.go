package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
)

type Article struct{
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles[] Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title:"Test Title", Desc: "Test Description1", Content: "Hallo Welt"},
	}

	fmt.Println("Endpoint Hi: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "HomePage Endpoint Hit <3 <3 <3")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":" + getPort(), nil))
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		return "8081"
	}
	return port
}


func main(){
	handleRequests()
}
