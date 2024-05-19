package text

func Check(color func(string) ColoredText) FormattedText {
	return color(" ✔ ")
}

func Ok() FormattedText {
	return Check(Green)
}

func Ballot(color func(string) ColoredText) FormattedText {
	return color(" ✘ ")
}

func NotOk() FormattedText {
	return Ballot(Red)
}

func Branch(color func(string) ColoredText) FormattedText {
	return color("⎇")
}
