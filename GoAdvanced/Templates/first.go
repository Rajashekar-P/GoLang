package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Person struct {
	Id   string
	Name string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{Id: "9951", Name: "Shekar"}
	parsedTemplate, _ := template.ParseFiles("first.html")
	err := parsedTemplate.Execute(w, person)

	if err != nil {
		log.Printf("Error Occured While exe the Template:", err)
		return
	}

}
func main() {
	fmt.Println("Start Here")
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT,nil)
	if err != nil {
		log.Fatal("error starting http server:", err)
		return
	}
}
