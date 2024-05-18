package text

type FormattedText interface {
	Get() string
}

const (
	Esc byte = 27
)

type ColoredText interface {
	Get() string
}
