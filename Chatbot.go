//ChatBot project
//Author: Micheal Curley
//A program to return responses to user input
package main

import (
	"net/http"
	"text/template"
	"fmt"
)
type Chatdata struct {
	InitMsg string
}
func chatBotHandler(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("template/chatBot.html")
	if(err != nil){
		fmt.Println("Error message: ", err)
	}
	t.Execute(w, Chatdata{InitMsg: "Say something to Chatterbot"})
}
//main
func main() {
	http.HandleFunc("/", chatBotHandler)
	//Listen for requests on port 80
	http.ListenAndServe(":8080", nil)
}