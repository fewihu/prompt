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
	returnCodeCh := make(chan string)
	go exitCode(returnCodeCh)
	formattedReturnCodeCh := make(chan text.FormattedText)
	go exit.FormatExitCode(<-returnCodeCh, formattedReturnCodeCh)

	userCh := make(chan text.FormattedText)
	go username.GetUser(userCh)
	pwdCh := make(chan text.FormattedText)
	go pwd.FormatPwd(pwdCh)

	pwdRawCh := make(chan *string)
	go pwd.GetPwd(pwdRawCh)

	gitState := make(chan text.FormattedText)
	go git.GetStateOrBranch(pwdRawCh, gitState)

	fmt.Print(<-formattedReturnCodeCh)
	fmt.Print(" ")
	fmt.Print(<-userCh)
	fmt.Print(":")
	fmt.Print(<-pwdCh)
	fmt.Print("\n")
	fmt.Print(<-gitState)
	fmt.Print("\n")
}

func exitCode(ch chan<- string) {
	args := os.Args
	if len(args) > 1 {
		ch <- args[1]
		return
	}
	ch <- "?"
}
