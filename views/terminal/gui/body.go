package gui

type body interface {
	draw(start, end int) error
}
