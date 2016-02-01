package models

// Manga represents manga general infos.
type Manga struct {
	ID     string
	Title  string
	Tags   []string
	Image  string // URL to main image representing the manga
	Alias  string
	Status int
}
