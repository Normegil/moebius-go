package log

import (
	log "github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
)

var jsonFormatter = log.JSONFormatter{}

// CustomFieldJSONFormatter will add convenient field to all log messages
type CustomFieldJSONFormatter struct {
	Pid int
}

// Format will format a lo line and add custom fieds to it
func (f *CustomFieldJSONFormatter) Format(entry *log.Entry) ([]byte, error) {

	entry.Data["id"] = uuid.NewV4().String()
	entry.Data["pid"] = f.Pid

	return jsonFormatter.Format(entry)
}
