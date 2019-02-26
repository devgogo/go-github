package main

import (
	"context"
	"fmt"
	"github.com/wenmingtang/go-github/github"
)

func main() {
	client := github.NewClient(nil)
	user, _, err := client.Users.Get(context.Background(), "wenmingtang")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user.ID, *user.Login, *user.URL, *user.AvatarURL)

	opt := &github.ListOptions{Page: 1}
	users, _, err := client.Users.ListFollowers(context.Background(), "wenmingtang", opt)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, u := range users {
		fmt.Println(*u.Login, *u.URL)
	}
}
