// Package username formats the username
package username

import (
	"os/user"
	"prompt/text"
)

// GetPwd returns the username
func GetUser() text.FormattedText {
	currentUser, err := user.Current()
	if err != nil {
		return getDefault()
	}
	username := currentUser.Username
	return text.BoldColor(text.Blue(username))
}

func getDefault() text.FormattedText {
	return text.Normal("<?>")
}
