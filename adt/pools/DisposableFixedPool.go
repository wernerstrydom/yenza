package pools

import (
	"github.com/wernerstrydom/yenza/adt"
	"runtime"
)

type DisposableFixedPool struct {
	FixedPool
	disposed bool
}

func NewDisposableObjectPool(policy PooledObjectPolicy) *DisposableFixedPool {
	return NewDisposableObjectPoolWithPolicyAndSize(policy, runtime.NumCPU() * 2)
}

func NewDisposableObjectPoolWithPolicyAndSize(policy PooledObjectPolicy, maximumRetained int) *DisposableFixedPool {
	_, isDefaultPolicy := policy.(*DefaultPooledObjectPolicy)
	return &DisposableFixedPool{
		FixedPool: FixedPool{
			items:           make([]element, maximumRetained-1),
			policy:          policy,
			isDefaultPolicy: isDefaultPolicy,
			firstItem:       nil,
		},
		disposed:          false,
	}
}

func (d *DisposableFixedPool) Get() interface{} {
	if d.disposed {
		panic("object of type 'DisposableFixedPool' was disposed")
	}
	return d.FixedPool.Get()
}

func (d *DisposableFixedPool) Return(obj interface{}) {
	if d.disposed || !d.doReturn(obj) {
		d.disposeItem(obj)
	}
}

func (d *DisposableFixedPool) Dispose() {
	d.disposed = true
	d.disposeItem(d.firstItem)
	d.firstItem = nil

	for i := range  d.items {
		d.disposeItem(d.items[i].value)
		d.items[i].value = nil
	}
}

func (d *DisposableFixedPool) disposeItem(item interface{}) {
	disposable, ok := item.(adt.Disposable)
	if ok {
		disposable.Dispose()
	}
}


