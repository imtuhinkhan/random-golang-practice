package user

import (
	"os"
)

func User() {
	const (
		username = "Tuhin"
		password = "123456"
	)
	user := os.Args[1]
	pass := os.Args[2]
	if username == user && pass == password {
		println("Access Granted")
	} else if username == user && pass != password {
		println("Invalid Password")
	} else {
		println("No User  Found")

	}
}
