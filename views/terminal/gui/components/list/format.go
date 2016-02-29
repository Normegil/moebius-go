package list

func format(toDisplay string, maxSize int) string {
	toPrint := toDisplay
	if len(toPrint) > maxSize {
		toPrint = toPrint[0:maxSize-3] + "..."
	}
	return toPrint
}
