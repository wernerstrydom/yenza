package templates

//ObjectPool defines an interface for an object pool
type ObjectPool interface {
	// Gets an object from the pool
	Get() *T

	// Returns an object to the pool
	Return(obj *T)
}

//PooledObjectPolicy represents a policy for managing pooled objects.
type PooledObjectPolicy interface {
	//Create a T.
	Create() *T

	//Return runs some processing when an object was returned to the pool.
	//Can be used to reset the state of an object and indicate if the
	//object should be returned to the pool.
	Return(obj *T) bool
}

//DefaultPooledObjectPolicy is a default implementation
type DefaultPooledObjectPolicy struct {
}

//Create creates a new T
func (policy *DefaultPooledObjectPolicy) Create() *T {
	return new(T)
}

//Return always returns true
func (policy *DefaultPooledObjectPolicy) Return(obj *T) bool {
	return true
}

//DefaultObjectPool is a default implmenentation of an object pool
type DefaultObjectPool struct {
}

//NewDefaultObjectPool creates a default object pool
func NewDefaultObjectPool() ObjectPool {
	return &DefaultObjectPool{}
}

//Get returns an object from the pool, or a new object
func (pool *DefaultObjectPool) Get() *T {
	return nil
}

//Return returns an object from the pool, or a new object
func (pool *DefaultObjectPool) Return(obj *T) {
}
