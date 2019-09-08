package pools

// LeakTrackingPool tracks whether objects obtained from the pool is returned
type LeakTrackingPool struct {
	inner Pool
	trackers map[interface{}]*Tracker
}

// Get returns an item from the pool if it is available, or create a new one
func (l *LeakTrackingPool) Get() interface{} {
	value := l.inner.Get()
	l.trackers[value] = NewTracker(value)
	return value
}

// Return attempts to return an object to a pool, and stops tracking. If the pool is full, Return results in a no-op
func (l *LeakTrackingPool) Return(obj interface{}) {
	tracker, ok := l.trackers[obj]
	if ok {
		tracker.Dispose()
		delete(l.trackers, tracker)
	}
	l.inner.Return(obj)
}

// Creates a new tracking pool
func NewLeakTrackingPool(inner Pool) *LeakTrackingPool {
	return &LeakTrackingPool{
		inner:    inner,
		trackers: make(map[interface{}]*Tracker),
	}
}