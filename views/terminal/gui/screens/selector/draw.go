package selector

import log "github.com/Sirupsen/logrus"

// Draw selector screen
func (selector *Selector) Draw(active bool) error {
	err := selector.search.Draw(active && selector.active == searchID)
	if nil != err {
		return err
	}
	log.Debug("Search drawn successfully")

	err = selector.list.Draw(active && selector.active == listID)
	if nil != err {
		return err
	}
	log.Debug("List drawn successfully")

	return selector.Footer()
}
