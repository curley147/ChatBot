Data Representation and Querying Project 2017
ChatBot

This repository contains a web application and server written in the programming language Go. The author is Micheal Curley.

The objective of this years assignment to build a chatbot web application posed many challenges. I spent a large portion of my time on the regular expressions
and trying to figure out the syntax. I'm not fully satisfied with my attempt as I feel the work didnt come off on the final product.
I got the basics working but I feel I should have done better. 

Refrences used: 
https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
https://developer.mozilla.org/en-US/docs/Learn/HTML/Forms/Sending_and_retrieving_form_data
http://regexlib.com/CheatSheet.aspx?AspxAutoDetectCookieSupport=1
https://github.com/data-representation/eliza

Running the code

To run the code in this repository, the files must first be compiled. The Go compiler must first be installed on your machine. Once that is installed, the code can be compiled and run by following these steps. We assume you are using the command line.

Clone this repository using Git.
> git clone https://github.com/curley147/ChatBot.git
Change into the folder.
> cd ChatBot
Compile the first file with the following command.
> go build Chatbot.go
Run the executable produced.
> Chatbot.exe
After you get the server running open a browser and go to localhost:8080/chatBot.html
