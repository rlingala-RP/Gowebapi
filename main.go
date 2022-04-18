package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type whoami struct {
	Name  string
	Title string
	State string
}

func main() {
	request1()
}

func WhoAmI(response http.ResponseWriter, r *http.Request) {
	who := []whoami{
		whoami{Name: "Ram Lingala",
			Title: "DevOps Engineer",
			State: "LN",
		},
	}

	json.NewEncoder(response).Encode(who)

	fmt.Println("Endpoint Hit: whoami")
}

func HomePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to Publicis Assessment!!!")
	fmt.Println("Endpoint Hit: homePage")
}

func AboutMe(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Hello from Ram Lingala !!!!")
	fmt.Println("Endpoint Hit: aboutMe")
}

func request1() {
	http.HandleFunc("/homepage", HomePage)
	http.HandleFunc("/aboutme", AboutMe)
	http.HandleFunc("/whoami", WhoAmI)

	log.Fatal(http.ListenAndServe(":8181", nil))
}
