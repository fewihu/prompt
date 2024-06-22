// Package username formats the username
package username

import (
	"os/user"
	"prompt/text"
)

// GetPwd returns the username
func GetUser(ch chan<- text.FormattedText) {
	currentUser, err := user.Current()
	if err != nil {
		ch <- getDefault()
		return
	}
	username := currentUser.Username
	ch <- text.BoldColor(text.Blue(username))
}

func getDefault() text.FormattedText {
	return text.Normal("<?>")
}
