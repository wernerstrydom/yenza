package pools

// PooledObjectPolicy represents a policy for managing pooled objects.
type PooledObjectPolicy interface {
	// Create creates a new object
	Create() interface{}

	// Return runs some processing on the object before returning it to the pool.
	// For example it may reset the value
	Return(obj interface{}) bool
}
