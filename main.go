package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Are you sure you want to know? [y/n]")
	fmt.Println()
	fmt.Print("-->")
	in, _ := reader.ReadString('\n')

	var absolute bool
	if strings.Replace(in, "\n", "", 1) == "y" {
		absolute = true
	} else if strings.Replace(in, "\n", "", 1) == "n" {
		absolute = false
	} else {
		absolute = false
		fmt.Println()
		fmt.Println("I don't know what that means, so I guess you dont care that much.")
		fmt.Println()
	}

	var responses [6]string
	responses[0] = "Yes"
	responses[1] = "No"
	responses[2] = "Maybe"
	responses[3] = "Ask me after another beer"
	responses[4] = "Terrible idea ... do it"
	responses[5] = "Hahahahahahahaha ... no"

	numResponses := len(responses)
	if absolute {
		numResponses = 2
	}
	var t int64
	t = int64(time.Now().Second())
	rand.Seed(t)
	n := rand.Intn(numResponses)
	fmt.Println()
	fmt.Println(responses[n])
}
