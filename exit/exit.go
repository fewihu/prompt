// Package exit handles $?
package exit

import (
	"prompt/text"
)

const (
	ok = 0
)

func FormatExitCode(code string, ch chan<- text.FormattedText) {
	if code != "0" {
		ch <- text.BoldColor(text.Red(code))
		return
	}
	ch <- text.Ok()

}








