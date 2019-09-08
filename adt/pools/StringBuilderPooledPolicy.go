package pools

import "strings"

type StringBuilderPooledPolicy struct {
	initialCapacity         int
	maximumRetainedCapacity int
}

func NewStringBuilderPooledPolicyWithMaximumCapacity(initialCapacity int, maximumRetainedCapacity int) *StringBuilderPooledPolicy {

	if initialCapacity >= maximumRetainedCapacity {
		panic("The initial capacity cannot be the same size or larger than the maximum retained capacity")
	}

	return &StringBuilderPooledPolicy{
		initialCapacity: initialCapacity,
		maximumRetainedCapacity: maximumRetainedCapacity,
	}
}

func NewStringBuilderPooledPolicy() *StringBuilderPooledPolicy {
	return NewStringBuilderPooledPolicyWithCapacity(100)
}

func NewStringBuilderPooledPolicyWithCapacity(initialCapacity int) *StringBuilderPooledPolicy {
	return NewStringBuilderPooledPolicyWithMaximumCapacity(256, 4*1024)
}

func (s *StringBuilderPooledPolicy) Create() interface{} {
	builder := &strings.Builder{}
	builder.Grow(s.initialCapacity)
	return builder
}

func (s *StringBuilderPooledPolicy) Return(obj interface{}) bool {
	builder, ok := obj.(*strings.Builder)
	if ok {
		builder.Reset()
	}
	return true
}

