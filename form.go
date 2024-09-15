package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"github.com/ousszizou/go_backend/08_forms/message"

	glandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	Port  = ":8080"
)
//html files processing 
func render(w http.ResponseWriter, filename string,data interface{}){
	t,err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal("Error econcounterd while parsing the template: ",err)
	}
	t.Execute(w, data)
}

func contanctHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		render(w, "./templates/contact.html",nil)
	}
	if r.Method == "POST"{
		msg := &message.Message {
			Email: r.PostFormValue("email"),
			Content: r.PostFormValue("content"),
		}
		if msg.Validate() == false {
			render(w, "./templates/contact.html",msg)
			return
		}
	}
}
func confirmHandler(w http.ResponseWriter, r*http.Request){
	render(w,"./templates/confirm.html",nil)
}

func main () {
	r:=mux.NewRouter() 
	r.HandelFunc("/contact", contanctHandler).Method("GET","POST")
	r.HandelFunc("/confirm", confirmHandler)

	loggedRouter := ghandlers.LoggingHandler(os.Stdout,r)
	http.ListenAndServe(Port, loggedRouter)
}