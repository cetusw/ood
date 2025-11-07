package document

type Item interface {
	ToHTML() string
	ToString() string
}
