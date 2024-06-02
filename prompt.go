// Package main gets all information for the prompt and prints them to STDOUT formatted
package main

import (
	"fmt"
	"os"
	"prompt/exit"
	"prompt/git"
	"prompt/pwd"
	"prompt/text"
	"prompt/username"
)

func main() {
	ret := exit.FormattExitCode(exitCode())
	user := username.GetUser()
	dir := pwd.FormatPwd()

	pwdRawCh := make(chan *string)
	go pwd.GetPwd(pwdRawCh)

	gitState := make(chan text.FormattedText)
	go git.GetStateOrBranch(pwdRawCh, gitState)

	fmt.Print(ret.Get() + " ")
	fmt.Print(user.Get())
	fmt.Print(":" + dir.Get())
	fmt.Print("\n" + (<-gitState).Get() + "\n")
}

func exitCode() string {
	args := os.Args
	if len(args) > 1 {
		return args[1]
	}
	return "?"
}
