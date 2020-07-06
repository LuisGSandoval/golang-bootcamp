package passme

import (
	"fmt"
	"os"
)

// Program is a test func
func Program() {

	if len(os.Args) < 2 {
		fmt.Println("[username] [password]")
		return
	}

	const (
		validUser = "Luis"
		validPwd  = "1234"
	)

	username, password := os.Args[1], os.Args[2]

	if validUser != username && validPwd != password {
		fmt.Printf("Access denied for %s \n", username)
	} else if validUser == username && validPwd != password {
		fmt.Printf("password incorrect for %s \n", username)
	} else if validUser == username && validPwd == password {
		fmt.Printf("Acces granted %s. Welcome!  \n", username)
	}
}
