//ChatBot project
//Author: Micheal Curley
//A program to return responses to user input
package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

//ChatResponse method
func ChatResponse(input string) string {
	//saved responses array
	responses := [3]string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	//regexp strings to look for certain patterns
	reObjs := [6]string{
		//check for father regardless of lower/upper case
		"(?i).*\b[Ff]ather\b.*",
		//check for "I am", "I'm", "Iam", "i'm", "i am", "iam" in input string and capturing data
		"(?i)(I[' a]*m)([^.?!]*)[.?!]?",
		//check for "i think" and capturing data
		"(?i)(I think)([^.?!]*)[.?!]?",
		//check for you and ?
		"(?i)(you+)(.*)(\\?+)",
		//check for "i like" and capturing data
		"(?i)(I like)(\\s+)[.?!]?",
		//check for numbers
		"\\d",
	}
	//for loop iterating over regexp strings and compiling them
	for i := 0; i < len(reObjs); i++ {
		//test print to console
		//fmt.Println(reObjs[i])
		re := regexp.MustCompile(reObjs[i])
		match := re.MatchString(input)
		if match == true {
			switch i {
			case 0:
				//father response
				return "Why don’t you tell me more about your father?"
			case 1:
				//replacing start of response and adding captured data
				return re.ReplaceAllString(input, "How do you know you are $2?")
			case 2:
				return re.ReplaceAllString(input, "What makes you think that $2?")
			case 3:
				//talking about you not me response
				return "I thought we were talking about you, not me!"
			case 4:
				//I like response
				return re.ReplaceAllString(input, "Why is it that you like $2?")
			case 5:
				//numbers response
				return "I'm not a fan of numbers, can you simplify it in any way?"
			default:
				return "Error message"
			}
		}
	}
	return responses[random(0, len(responses))]
}

//random number generator method
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

//Chatdata struct
type Chatdata struct {
	UserInput string
	Response  string
}

func chatBotHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/chatBot.html")
	if err != nil {
		fmt.Println("template retrieval failed:", err)
	}
	t.Execute(w, Chatdata{UserInput: ""})
}

//main
func main() {
	//template file handler
	http.HandleFunc("/", chatBotHandler)
	//Listen for requests on port 80
	http.ListenAndServe(":8080", nil)
}
