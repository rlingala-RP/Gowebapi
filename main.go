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

func whoAmI(response http.ResponseWriter, r *http.Request) {
	who := []whoami{
		whoami{Name: "Ram Lingala",
			Title: "DevOps Engineer",
			State: "LN",
		},
	}

	json.NewEncoder(response).Encode(who)

	fmt.Println("Endpoint Hit: whoami")
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the Go Web API!!!")
	fmt.Println("Endpoint Hit: homePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Hello from Ram Lingala !!!!")
	fmt.Println("Endpoint Hit: aboutMe")
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/whoami", whoAmI)

	log.Fatal(http.ListenAndServe(":8181", nil))
}
