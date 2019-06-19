package main

import (
	"O/O"
)

func main() {

	O.Test()

	user := O.User{
		ID: 1,
		Name: "one",
	}
	user.Info()

	item := O.Item{
		ID: 10,
		Name: "ten",
		User: user,
	}
	item.Info()

	O.EntityInfo(&user)
	O.EntityInfo(&item)

}
