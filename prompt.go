// Package main get all information for the prompt and prints them to STDOUT formatted
package main

import (
	"fmt"
	"prompt/pwd"
	"prompt/username"
)

func main() {
	user := username.GetUser()
	dir := pwd.GetPwd()
	fmt.Print(user.Get())
	fmt.Print(":" + dir.Get())
}
