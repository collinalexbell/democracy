package main

import (
	"fmt"
	"os"

	"github.com/kuberlog/democracy/construct"
	"github.com/sethvargo/go-password/password"
)

func generateRandomPassword() (string, error) {
	return password.Generate(31, 10, 10, true, false)
}

func main() {
	if os.Args[1] == "createBitLaunchAccount" {
		email := os.Args[2]
		password, err := generateRandomPassword()
		if err != nil {
			panic("password generation crashed")
		}
		err = construct.CreateBitLaunchAccount(email, password)
		if err != nil {
			fmt.Printf("error creating account\n")
		} else {
			fmt.Printf("Account created\nusername: %s\npassword: %s", email, password)
		}
	}
}
