package gui

import "github.com/nsf/termbox-go"

func drawHeader() {
	w, _ := termbox.Size()
	fill(coordinates{0, 0}, sizes{w, 1}, attributes{
		foreground: termbox.ColorBlack,
		background: termbox.ColorCyan,
	})
	print(coordinates{headerOpts.titlePt, 0}, attributes{
		foreground: termbox.ColorBlack | termbox.AttrBold,
		background: termbox.ColorCyan,
	}, "Moebius-go")
}
