package cache

// Cache interface contains both Saver and Loader methods
type Cache interface {
	Loader
	Saver
}
