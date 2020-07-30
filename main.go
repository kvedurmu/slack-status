package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	status := "away"
	dnd := false

	if len(os.Args) > 1 {
		if os.Args[1] == "active" {
			status = "auto"
		} else if os.Args[1] == "dnd" {
			dnd = true
		}
	}

	home, _ := os.UserHomeDir()
	file, err := os.Open(home + "/.slack_api/keys")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		api := slack.New(scanner.Text())
		err = api.SetUserPresence(status)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		if dnd {
			_, _ = api.SetSnooze(60)
		} else {
			_, _ = api.EndSnooze()
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
