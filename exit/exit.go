// Package exit handles $?
package exit

import (
	"prompt/text"
)

const (
	ok = 0
)

func FormattExitCode(code string) text.FormattedText {
	if code != "0" {
		return text.BoldColor(text.Red(code))
	}
	return text.Ok()
}
