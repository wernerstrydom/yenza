package adt

// Disposable provides a common way to dispose of objects. For example, closing a file, or freeing memory uses very
// different APIs. However, when those are stored in an abstract data type, like a cache or a pool, having a common
// interface simplifies the code considerably.
type Disposable interface {
	Dispose()
}
