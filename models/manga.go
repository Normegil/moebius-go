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

func (manga Manga) String() string { return manga.Title }

// ByTitle compare mangas by their titles for sorting purpose
type ByTitle []Manga

func (mangas ByTitle) Len() int           { return len(mangas) }
func (mangas ByTitle) Swap(i, j int)      { mangas[i], mangas[j] = mangas[j], mangas[i] }
func (mangas ByTitle) Less(i, j int) bool { return mangas[i].Title < mangas[j].Title }
