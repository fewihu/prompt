// Package main get all information for the prompt and prints them to STDOUT formatted
package main

import (
	"fmt"
	"os"
	"prompt/exit"
	"prompt/pwd"
	"prompt/username"
)

func main() {
	ret := exit.FormattExitCode(exitCode())
	user := username.GetUser()
	dir := pwd.GetPwd()
	fmt.Print(ret.Get() + " ")
	fmt.Print(user.Get())
	fmt.Print(":" + dir.Get())
}

func exitCode() string {
	args := os.Args
	if len(args) > 1 {
		return args[1]
	}
	return "?"
}
