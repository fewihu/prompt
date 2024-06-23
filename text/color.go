// Package text allow to set formatting and colors on text
package text

const (
	Esc byte = 27
)

type ColoredText interface {
	String() string
}

type coloredText struct {
	text      string
	colorCode string
}

var (
	red     = colorCode('1', '9', '6')
	blue    = colorCode('2', '6')
	cyan    = colorCode('3', '7')
	green   = colorCode('2', '8')
	black   = colorCode('2', '3', '3')
	violett = colorCode('6', '9')
)

func colorCode(code ...byte) string {
	sequence := []byte{Esc, '[', '3', '8', ';', '5', ';'}
	return string(append(append(sequence, code...), 'm'))
}

func (t *coloredText) String() string {
	return t.colorCode + t.text
}

func Red(text string) ColoredText {
	return &coloredText{text, red}
}

func Blue(text string) ColoredText {
	return &coloredText{text, blue}
}

func Cyan(text string) ColoredText {
	return &coloredText{text, cyan}
}

func Green(text string) ColoredText {
	return &coloredText{text, green}
}

func Black(text string) ColoredText {
	return &coloredText{text, black}
}

func Violett(text string) ColoredText {
	return &coloredText{text, violett}
}
