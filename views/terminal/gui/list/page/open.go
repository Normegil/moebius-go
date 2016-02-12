package page

func (goTo *GoTo) Open(initialContent string) {
	goTo.enabled = true
	goTo.Content = initialContent
}
