package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var t int64
	t = int64(time.Now().Second())
	rand.Seed(t)
	n := rand.Intn(6)
	var responses [6]string
	responses[0] = "Yes"
	responses[1] = "No"
	responses[2] = "Maybe"
	responses[3] = "Ask me after another beer"
	responses[4] = "Terrible idea ... do it"
	responses[5] = "Hahahahahahahaha"
	fmt.Println(responses[n])

	// if n == 0 {
	// 	println("No")
	// } else if n == 1 {
	// 	println("Yes")
	// } else if n == 2 {
	// 	println("Maybe")
	// }
}
