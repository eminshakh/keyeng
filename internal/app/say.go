package app

import (
	"os/exec"
	"strings"
)

func Say(s <-chan string) {
	go func(s <-chan string) {
		for word := range s {
			split := strings.Split(word, " - ")
			if word[0] >= 97 && word[0] <= 122 {
				cmd := exec.Command("say", split[0])
				cmd.Run()
			} else {
				cmd := exec.Command("say", split[1])
				cmd.Run()
			}
		}
	}(s)
}
