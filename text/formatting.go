package text

import (
	"strings"
)

//  	set bold mode.
// ESC[2m 	ESC[22m 	set dim/faint mode.
// ESC[3m 	ESC[23m 	set italic mode.
// ESC[4m 	ESC[24m 	set underline mode.
// ESC[5m 	ESC[25m 	set blinking mode
// ESC[7m 	ESC[27m 	set inverse/reverse mode
// ESC[8m 	ESC[28m 	set hidden/invisible mode
// ESC[9m 	ESC[29m 	set strikethrough mode.

type formattedText struct {
	formatCode string
	text       string
	resetCode  string
}

func (f *formattedText) Get() string {
	return f.formatCode + f.text + f.resetCode
}

func JoinSpace(texts ...FormattedText) FormattedText {
	return join(" ", texts...)
}

func JoinNarrow(texts ...FormattedText) FormattedText {
	return join("", texts...)
}

func join(del string, texts ...FormattedText) FormattedText {
	var sb strings.Builder
	for _, t := range texts {
		sb.WriteString(t.Get() + del)
	}
	return &formattedText{
		formatCode: "",
		text:       sb.String(),
		resetCode:  "",
	}
}

func Bold(text string) FormattedText {
	// ESC[1m
	bold := []byte{Esc, '[', '1', 'm'}
	// ESC[22m
	//	end := []byte{Esc, '[', '2', '9', 'm'}
	end := []byte{Esc, '[', '0', 'm'}
	return &formattedText{string(bold), text, string(end)}
}

func BoldColor(text ColoredText) FormattedText {
	// ESC[1m
	bold := []byte{Esc, '[', '1', 'm'}
	// ESC[22m
	end := []byte{Esc, '[', '0', 'm'}
	return &formattedText{string(bold), text.Get(), string(end)}
}

func Normal(text string) FormattedText {
	end := []byte{Esc, '[', '0', 'm'}
	return &formattedText{string(end), text, ""}
}

func Newline() FormattedText {
	return Normal("\n")
}
