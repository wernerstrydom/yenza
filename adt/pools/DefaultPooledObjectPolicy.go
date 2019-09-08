package pools

// DefaultPooledObjectPolicy represents a default implementation that does nothing in particular
type DefaultPooledObjectPolicy struct {
	factory func() interface{}
}

// Create calls the factory method
func (d *DefaultPooledObjectPolicy) Create() interface{} {
	return d.factory()
}

// Return does nothing when an object is returned to the pool
func (d *DefaultPooledObjectPolicy) Return(obj interface{}) bool {
	return true
}

func NewDefaultPooledObjectPolicy(factory func() interface{}) *DefaultPooledObjectPolicy {
	return &DefaultPooledObjectPolicy{factory: factory}
}

