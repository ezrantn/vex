package formatter

type FormatHandler interface {
	Convert(sourceType string, toType string) error
}