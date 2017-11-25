//ChatBot project
//Author: Micheal Curley
//A program to return responses to user input
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
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
	reObjs := [11]string{
		//check for "my name is" and capturing data
		"(?i)(.*my name is)(.*)",
		//check for "I am", "I'm", "Iam", "i'm", "i am", "iam" in input string and capturing data
		"(?i)(.*I[' a]*m)([^.?!]*)[.?!]?",
		//check for "i think" and capturing data
		"(?i)(.*I think)([^.?!]*)[.?!]?",
		//check for you and ?
		"(?i)(you+)(.*)(\\?+)[.?!]?",
		//check for "i like" and capturing data
		"(?i)(.*I like)(\\s+)[.?!]?",
		//check for numbers
		"\\d[.?!]?",
		//check for how are you? from the user
		"(?i).*hello.*[.?!]?",
		//check for "i need" and capturing data
		"(?i)(^\\s*I need)([^.!?]*)[.!?]*\\s*$",
		//check for how are you? from the user
		"(?i).*where.*\\?+[.?!]?",
		//check for "why dont you" and capturing data
		"(?i)(^\\s*why can'?t i)([^.!?]*)[.!?]*\\s*$",
		//check for "why cant i" and capturing data
		"(?i)(^\\s*why don'?t you)([^.!?]*)[.!?]*\\s*$",
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
				//replacing "my name is" and adding captured data
				return re.ReplaceAllString(input, "Hi $1, how are you?")
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
			case 6:
				//hello response
				return "Hello there again!"
			case 7:
				//replacing "i need" and adding captured data
				return re.ReplaceAllString(input, "Why do you need $2?")
			case 8:
				//greeting response
				return "Somewhere over the rainbow!"
			case 9:
				//why dont you response
				return re.ReplaceAllString(input,  "Perhaps eventually I will $2")
			case 10:
				//why dont you response
				return re.ReplaceAllString(input,  "If you could $2, what would you do?")
			default:
				return "Error message"
			}
		}
	}
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(len(responses))
	return responses[randNum]
}

//Reflect Adapted from https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
//Pronouns reflection method
func Reflect(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile("\b")
	tokens := boundaries.Split(input, -1)
	// List the reflections.
	reflections := [][]string{
		{"your", "my"},
		{"my", "your"},
		{"you", "I"},
		{"me", "you"},
		{"am", "are"},
		{"are", "am"},
		{"yours", "mine"},
		{"mine", "yours"},
		{"myself", "yourself"},
		{"yourself", "myself"},
		{"i'd", "you would"},
		{"i've", "you have"},
		{"i'll", "you will"},
	}
	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}
	// Put the tokens back together.
	return strings.Join(tokens, ``)
}
func chatBotHandler(w http.ResponseWriter, r *http.Request) {
	//storing saved value from incoming request
	input := r.Header.Get("value")
	//printing it
	fmt.Fprintf(w, "You: %s\n", input)
	response := Reflect(input)
	refResponse := ChatResponse(response)
	fmt.Fprintf(w, "ChatterBot: %s\n", refResponse)
}

//redirect method
func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "127.0.0.1/chatBot.html", 301)
}

//main
func main() {
	//static file handler
	http.Handle("/", http.FileServer(http.Dir("./html")))
	//template file handler
	http.HandleFunc("/postSend", chatBotHandler)
	//Listen for requests on port 80
	http.ListenAndServe(":8080", nil)
}
