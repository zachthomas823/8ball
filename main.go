package main

// Imports
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)                 // Reader for input from user
	fmt.Println("Are you sure you want to know? [y/n]") // Prompt user
	fmt.Println()
	fmt.Print("-->")
	in, _ := reader.ReadString('\n') // Read user's input

	var absolute bool                            // boolean for whether or not the user wants absolute answers
	if strings.Replace(in, "\n", "", 1) == "y" { // If user answers "y"
		absolute = true // Make absolute equal to true
	} else if strings.Replace(in, "\n", "", 1) == "n" { // If the user answers "n"
		absolute = false // Make absolute equal to false
	} else { // If the user answers anything else
		absolute = false // Set absolute equal to false
		fmt.Println()
		fmt.Println("I don't know what that means, so I guess you dont care that much.") // Chastise the user
		fmt.Println()
	}

	var responses [6]string // All of the built in responses
	responses[0] = "Yes"
	responses[1] = "No"
	responses[2] = "Maybe"
	responses[3] = "Ask me after another beer"
	responses[4] = "Terrible idea ... do it"
	responses[5] = "Hahahahahahahaha ... no"

	numResponses := len(responses)
	if absolute { // If the user wants absolute answers set the number of responses to 2 (yes and no)
		numResponses = 2
	}
	var t int64                    // variable to store current time (seconds)
	t = int64(time.Now().Second()) // get current time (seconds)
	rand.Seed(t)                   // Set the seed equal to t in order to achieve pseudorandomness
	n := rand.Intn(numResponses)   // get a random int based on the random number t
	fmt.Println()
	fmt.Println(responses[n]) // print the random response from the selection
}
