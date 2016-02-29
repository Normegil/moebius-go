package list

import (
	"fmt"
	"strings"
)

func (list *List) Filter(newFilter string) {
	list.Displayable = make([]fmt.Stringer, 0)
	for _, object := range list.OriginalContent {
		if strings.Contains(strings.ToLower(object.String()), strings.ToLower(newFilter)) {
			list.Displayable = append(list.Displayable, object)
		}
	}
}
