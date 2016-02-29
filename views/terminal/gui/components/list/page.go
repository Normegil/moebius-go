package list

import log "github.com/Sirupsen/logrus"

func (list List) Page() int {

	log.WithFields(log.Fields{
		"selected":  list.Selected,
		"displayed": list.Displayed(),
	}).Debug("Calculating PageNB")

	if list.Displayed() == 0 {
		return 0
	}
	return list.Selected / list.Displayed()
}
