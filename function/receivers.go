package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:""`
	Likes   []Like
}

type Like struct {
	ID   string
	Date time.Time
}

func (u *User) Like() {
	if u.Likes == nil {
		u.Likes = make([]Like, 0)
	}
	like := Like{
		ID:   "id",
		Date: time.Time{},
	}
	u.Likes = append(u.Likes, like)
}

func (u User) LikeCount() int {
	return len(u.Likes)
}

func main() {
	user1 := User{
		Name:    "go",
		Surname: "turkiye",
	}
	user1.Like()
	user1.Like()
	user1.Like()
	user1.Like()

	arr, _:=json.Marshal(user1)
	fmt.Println(string(arr))
	fmt.Println(user1.LikeCount())
}
