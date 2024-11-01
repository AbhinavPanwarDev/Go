package main

import "fmt"

func main () {

	user, err := getUser()

	if err!= nil {
		fmt.Println(err)
		return
	}

	user, err := getUserProfile(user.ID) {
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}