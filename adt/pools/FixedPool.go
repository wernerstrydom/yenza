package pools

import (
	"runtime"
)

type element struct {
	value interface{}
}

//FixedPool represents an default implementation of Pool
type FixedPool struct {
	items []element
	policy PooledObjectPolicy
	isDefaultPolicy bool
	firstItem interface{}
}

func NewFixedPoolWithPolicy(policy PooledObjectPolicy) *FixedPool {
	return NewFixedPoolWithPolicyAndSize(policy, runtime.NumCPU() * 2)
}

func NewFilledFixedPoolWithPolicyAndSize(policy PooledObjectPolicy, maximumRetained int) *FixedPool {
	_, isDefaultPolicy := policy.(*DefaultPooledObjectPolicy)
	elements := make([]element, maximumRetained-1)
	for key := range elements {
		elements[key].value = policy.Create()
	}

	return &FixedPool{
		items:           elements,
		policy:          policy,
		isDefaultPolicy: isDefaultPolicy,
		firstItem:       nil,
	}
}

func NewFixedPoolWithPolicyAndSize(policy PooledObjectPolicy, maximumRetained int) *FixedPool {
	_, isDefaultPolicy := policy.(*DefaultPooledObjectPolicy)
	return &FixedPool{
		items:           make([]element, maximumRetained-1),
		policy:          policy,
		isDefaultPolicy: isDefaultPolicy,
		firstItem:       nil,
	}
}

func (d *FixedPool) Get() interface{} {
	var item = d.firstItem
	if item == nil || compareExchange(&d.firstItem, nil, item) != item {
		var items= d.items
		for i := 0; i < len(items); i++ {
			item = items[i].value
			if item != nil && compareExchange(&items[i].value, nil, item) == item {
				return item
			}
		}
		return d.Create()
	}
	return item
}

func compareExchange(location1 *interface{}, value interface{}, comparand interface{}) interface{} {
	// TODO: Make this thread safe
	original := *location1
	if *location1 == comparand {
		*location1 = value
	}
	return original
}

// Return attempts to return the object to the pool, except when the policy returns false, or the pool is full
func (d *FixedPool) Return(obj interface{}) {
	d.doReturn(obj)
}

// doReturn attempts to return the object to the pool, except when the policy returns false, or the pool is full
func (d *FixedPool) doReturn(obj interface{}) bool {
	if !d.isDefaultPolicy && !d.policy.Return(obj) {
		return false
	}

	if d.firstItem == nil && compareExchange(&d.firstItem, obj, nil) == nil {
		return true
	}

	items := d.items
	for i := 0; i < len(items); i++ {
		if compareExchange(&items[i].value, obj, nil) == nil {
			return true
		}
	}

	return false
}

func (d *FixedPool) Create() interface{} {
	return d.policy.Create()
}
