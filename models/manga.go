package models

// MangaStatus represents status of a Manga
type MangaStatus int

const (
	// Suspended means a manga is OnHold for an undefined period
	Suspended MangaStatus = iota
	// Ongoing means a manga is still being published
	Ongoing
	// Completed means a manga is has reached its end.
	Completed
)

// Manga represents manga general infos.
type Manga struct {
	ID     string
	Title  string
	Tags   []string
	Image  string // URL to main image representing the manga
	Alias  string
	Status MangaStatus
}
