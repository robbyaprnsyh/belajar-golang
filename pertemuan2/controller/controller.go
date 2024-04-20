package controller

import "fmt"

func main() {
	fmt.Println(controller.ReadUserList())
}

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ReadUserList() (user []user) {
	users = make([]user, 2)

	users[0] = user{
		ID:    1,
		Name:  "Robby Apriansyah",
		Email: "robbyapriansyah979@gmail.com",
	}

	users[1] = user{
		ID:    2,
		Name:  "Diva Dien Alhaq",
		Email: "diva@plabs.id",
	}

	return
}
