package utils

type Popup interface {
	Open(content string) error
	Enabled() bool
	Drawer
	Reacter
}
