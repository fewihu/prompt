// Package pwd formats the current working directory(pwd)
package pwd

import (
	"os"
	"os/user"
	"prompt/text"
	"strings"
)

// FormatPwd returns the normalized path of the current directory or a default value
func FormatPwd(ch chan<- text.FormattedText) {
	pwd, err := os.Getwd()
	if err != nil {
		ch <- getDefault()
		return
	}

	currentUser, err := user.Current()
	if err != nil {
		ch <- getDefault()
		return
	}
	username := currentUser.Username

	pwd = strings.Replace(pwd, "/home/"+username, "~", -1)
	ch <- text.BoldColor(text.Cyan(pwd))
}

// GetPwd return the pwd or nil
func GetPwd(pwdCh chan<- *string) {
	pwd, err := os.Getwd()
	if err != nil {
		pwdCh <- nil
		return
	}
	pwdCh <- &pwd
}

func getDefault() text.FormattedText {
	return text.Normal("<?>")
}
