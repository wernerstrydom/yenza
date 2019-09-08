package pools

// Pool represents a pool of objects that can be reused
type Pool interface {
	// Get returns an available object from the pool, or creates a new one
	Get() interface{}

	// Return attempts to return obj to the pool
	Return(obj interface{})
}