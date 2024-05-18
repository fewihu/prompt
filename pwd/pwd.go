// Package pwd formats the current working directory(pwd)
package pwd

import (
	"os"
	"os/user"
	"prompt/text"
	"strings"
)

// GetPwd returns the normalized path of the current directory or a default value
func GetPwd() text.FormattedText {
	pwd, err := os.Getwd()
	if err != nil {
		return getDefault()
	}

	currentUser, err := user.Current()
	if err != nil {
		return getDefault()
	}
	username := currentUser.Username

	pwd = strings.Replace(pwd, "/home/"+username, "~", -1)
	return text.BoldColor(text.Cyan(pwd))
}

func getDefault() text.FormattedText {
	return text.Normal("<?>")
}
