package page

import log "github.com/Sirupsen/logrus"

// Open the Popup and set it's content t given string
func (goTo *GoTo) Open(initialContent string) {
	log.WithField("content", initialContent).Debug("Enabling GoTo Popup")
	goTo.enabled = true
	goTo.Content = initialContent
}
