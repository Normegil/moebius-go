package cache

import "github.com/normegil/moebius-go/models"

// MemoryCache is a cache that will register every object in memory and send them back when needed
type MemoryCache struct {
	memory map[string][]models.Manga
}

// NewMemoryCache create an initialized memory cache
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		memory: make(map[string][]models.Manga),
	}
}

// Save will put mangas in memory for later usage
func (cache *MemoryCache) Save(mangas []models.Manga, ID string) error {
	cache.memory[ID] = mangas
	return nil
}

// Load will get mangas from memory
func (cache *MemoryCache) Load(ID string) ([]models.Manga, error) {
	mangas, present := cache.memory[ID]
	if !present {
		return nil, DataNotFoundError{ID}
	}
	return mangas, nil
}
