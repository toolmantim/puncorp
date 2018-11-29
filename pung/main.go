package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomPun generates a random pun
func RandomPun() string {
	puns := []string{
		"My scripting is rubbash.",
		"I'd switch to zsh, but I'm too bashful.",
		"You use Bash? Go fish!",
		"I went to see a Doctor about a bash. It was terminal",
		"Just bash it out.",
	}
	rand.Seed(time.Now().Unix())
	return puns[rand.Int()%len(puns)]
}

func main() {
	fmt.Println(RandomPun())
}
