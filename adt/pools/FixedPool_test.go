package pools

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/wernerstrydom/yenza/adt/pools/mocks"
	"runtime"
	"testing"
)

type FixedPoolTestSuite struct {
	suite.Suite
}

func (suite *FixedPoolTestSuite) TestGet_WithNothingPooled() {
	// Arrange
	type stub struct {
		value int
	}

	expected := &stub{5}
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Create").Return(expected).Once()
	items := make([]element, 3)
	target := FixedPool{
		items:           items,
		policy:          policy,
		isDefaultPolicy: false,
		firstItem:       nil,
	}

	// Act
	actual := target.Get()

	// Assert
	assert.NotNil(suite.T(), actual, "Expected a valid item")
	assert.Same(suite.T(), expected, actual)
}

func (suite *FixedPoolTestSuite) TestGet_WithSingleItem() {
	// Arrange
	type stub struct {
		value int
	}

	expected := &stub{5}
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Create").Return(nil).Times(0)

	items := make([]element, 3)
	target := FixedPool{
		items:           items,
		policy:          policy,
		isDefaultPolicy: false,
		firstItem:       expected,
	}

	// Act
	actual := target.Get()

	// Assert
	assert.NotNil(suite.T(), actual, "Expected a valid item")
	assert.Same(suite.T(), expected, actual)
}

func (suite *FixedPoolTestSuite) TestGet_WithFragmentation1() {
	// Arrange
	type stub struct {
		value int
	}

	expected := &stub{5}
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Create").Return(nil).Times(0)

	items := make([]element, 3)
	items[0] = element{ value: expected }
	target := FixedPool{
		items:           items,
		policy:          policy,
		isDefaultPolicy: false,
		firstItem:       nil,
	}

	// Act
	actual := target.Get()

	// Assert
	assert.NotNil(suite.T(), actual, "Expected a valid item")
	assert.Same(suite.T(), expected, actual)
}

func (suite *FixedPoolTestSuite) TestGet_WithFragmentation2() {
	// Arrange
	type stub struct {
		value int
	}

	expected := &stub{5}
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Create").Return(nil).Times(0)

	items := make([]element, 3)
	items[1] = element{ value: expected }
	target := FixedPool{
		items:           items,
		policy:          policy,
		isDefaultPolicy: false,
		firstItem:       nil,
	}

	// Act
	actual := target.Get()

	// Assert
	assert.NotNil(suite.T(), actual, "Expected a valid item")
	assert.Same(suite.T(), expected, actual)
}

func (suite *FixedPoolTestSuite) TestReturn_WithPolicyReturnsFalse() {
	// Arrange
	type stub struct {
		value int
	}

	var item interface{} = &stub{5}
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Return", item).Return(false).Once()
	items := make([]element, 3)
	target := FixedPool{
		items:           items,
		policy:          policy,
		isDefaultPolicy: false,
		firstItem:       nil,
	}

	// Act
	target.Return(item)

	// Assert
	assert.Nil(suite.T(), target.firstItem)
	assert.Nil(suite.T(), target.items[0].value)
	assert.Nil(suite.T(), target.items[1].value)
	assert.Nil(suite.T(), target.items[2].value)
}

func (suite *FixedPoolTestSuite) TestReturn_WithEmptyPool() {
	// Arrange
	type stub struct {
		value int
	}

	var item interface{} = &stub{5}
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Return", item).Return(true).Once()
	items := make([]element, 3)
	target := FixedPool{
		items:           items,
		policy:          policy,
		isDefaultPolicy: false,
		firstItem:       nil,
	}

	// Act
	target.Return(item)

	// Assert
	assert.Same(suite.T(), item, target.firstItem, "expected first item to the item that was returned")
	assert.Nil(suite.T(), target.items[0].value)
	assert.Nil(suite.T(), target.items[1].value)
	assert.Nil(suite.T(), target.items[2].value)
}

func (suite *FixedPoolTestSuite) TestReturn_WithOnePooledObject() {
	// Arrange
	type stub struct {
		value int
	}
	s := &stub{4}
	var item interface{} = &stub{5}
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Return", item).Return(true).Once()
	items := make([]element, 3)
	target := FixedPool{
		items:           items,
		policy:          policy,
		isDefaultPolicy: false,
		firstItem:       s,
	}

	// Act
	target.Return(item)

	// Assert
	assert.Same(suite.T(), s, target.firstItem, "expected first item to the item that was returned")
	assert.Same(suite.T(), item, target.items[0].value)
	assert.Nil(suite.T(), target.items[1].value)
	assert.Nil(suite.T(), target.items[2].value)
}

func (suite *FixedPoolTestSuite) TestReturn_WithTwoPooledObjects() {
	// Arrange
	type stub struct {
		value int
	}
	s3 := &stub{3}
	s4 := &stub{4}
	var item interface{} = &stub{5}
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Return", item).Return(true).Once()
	items := make([]element, 3)
	items[0] = element{s4}
	target := FixedPool{
		items:           items,
		policy:          policy,
		isDefaultPolicy: false,
		firstItem:       s3,
	}

	// Act
	target.Return(item)

	// Assert
	assert.Same(suite.T(), s3, target.firstItem, "expected first item to the item that was returned")
	assert.Same(suite.T(), s4, target.items[0].value)
	assert.Same(suite.T(), item, target.items[1].value)
	assert.Nil(suite.T(), target.items[2].value)
}

func (suite *FixedPoolTestSuite) TestReturn_WithFullPool_DiscardsReturnedObject() {
	// Arrange
	type stub struct {
		value int
	}
	s1 := &stub{1}
	s2 := &stub{2}
	s3 := &stub{3}
	s4 := &stub{4}
	var item interface{} = &stub{5}
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Return", item).Return(true).Once()
	items := make([]element, 3)
	items[2] = element{s4}
	items[1] = element{s2}
	items[0] = element{s1}
	target := FixedPool{
		items:           items,
		policy:          policy,
		isDefaultPolicy: false,
		firstItem:       s3,
	}

	// Act
	target.Return(item)

	// Assert
	assert.Same(suite.T(), s3, target.firstItem)
	assert.Same(suite.T(), s1, target.items[0].value)
	assert.Same(suite.T(), s2, target.items[1].value)
	assert.Same(suite.T(), s4, target.items[2].value)
}

func (suite *FixedPoolTestSuite) TestCompareExchange_WithOriginalValueAndComparandMatching() {
	// Arrange
	type stub struct {
		value int
	}

	s1 := &stub{1}
	s2 := &stub{2}
	//s3 := &stub{3}

	var location1 interface{} = s1
	var value interface{} = s2
	var comparand interface{} = s1
	expected := location1

	// Act
	actual := compareExchange(&location1, value, comparand)

	// Assert
	assert.Same(suite.T(), expected, actual)
	assert.Same(suite.T(), location1, s2)
}

func (suite *FixedPoolTestSuite)  TestNewObjectPoolWithPolicy() {
	// Arrange
	policy := new(mocks.PooledObjectPolicy)

	// Act
	actual := NewFixedPoolWithPolicy(policy)

	// Assert
	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), runtime.NumCPU()*2-1, len(actual.items))
	assert.Nil(suite.T(), actual.firstItem)
	assert.Same(suite.T(), policy, actual.policy)
	assert.False(suite.T(), actual.isDefaultPolicy)
}

func (suite *FixedPoolTestSuite)  TestNewObjectPoolWithPolicyAndSize() {
	// Arrange
	policy := new(mocks.PooledObjectPolicy)

	// Act
	actual := NewFixedPoolWithPolicyAndSize(policy, 16)

	// Assert
	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), 15, len(actual.items))
	assert.Nil(suite.T(), actual.firstItem)
	assert.Same(suite.T(), policy, actual.policy)
	assert.False(suite.T(), actual.isDefaultPolicy)
}

func (suite *FixedPoolTestSuite) TestCompareExchange_WithOriginalValueAndComparandNotMatching() {
	// Arrange
	type stub struct {
		value int
	}

	s1 := &stub{1}
	s2 := &stub{2}
	s3 := &stub{3}

	var location1 interface{} = s1
	var value interface{} = s2
	var comparand interface{} = s3
	expected := location1

	// Act
	actual := compareExchange(&location1, value, comparand)

	// Assert
	assert.Same(suite.T(), expected, actual)
	assert.Same(suite.T(), location1, s1)
}


func TestFixedPool(t *testing.T) {
	suite.Run(t, new(FixedPoolTestSuite))
}

func ExampleFixedPool() {

	type StringTable struct {
		// details hidden
	}

	// A policy is in control of how objects are created, and what should happen when they are returned
	policy := NewDefaultPooledObjectPolicy(func() interface{} {
		return &StringTable{}
	})

	// Create a pool
	pool := NewFixedPoolWithPolicy(policy)

	// get a string table for use
	stringTable := pool.Get().(*StringTable)

	// Do something with the string table

	// return the string table to the pool where it can be used again
	pool.Return(stringTable)
}

const poolSize = 16

// This causes a new object to be created
func BenchmarkFixedPool_WithEmptyPool(b *testing.B) {
	type item struct {
		value int
	}
	policy := NewDefaultPooledObjectPolicy(func() interface{} {
		return &item{1}
	})
	pool := NewFixedPoolWithPolicyAndSize(policy, poolSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		elements := make([]interface{}, poolSize)
		for j := 0; j < poolSize; j++ {
			elements[j] = pool.Get()
		}
	}
}

// This Highlights any benefits of having the first item
func BenchmarkFixedPool_WithFullPool(b *testing.B) {
	type stub struct {
		value int
	}
	policy := NewDefaultPooledObjectPolicy(func() interface{} {
		return &stub{1}
	})

	pool := NewFixedPoolWithPolicyAndSize(policy, poolSize)
	items := make([]interface{}, poolSize)
	for i := 0; i < poolSize; i++ {
		items[i] = pool.Get()
	}
	for i := 0; i < poolSize; i++ {
		pool.Return(items[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// This forces all the objects to be retrieved and added back into the pool
		for j := 0; j < poolSize; j++ {
			items[j] = pool.Get()
		}
		for j := 0; j < poolSize; j++ {
			pool.Return(items[j])
		}
		for j := 0; j < poolSize; j++ {
			items[j] = pool.Get()
		}
	}
}

func BenchmarkFixedPool_WithCommonScenario(b *testing.B) {
	type stub struct {
		value int
	}
	policy := NewDefaultPooledObjectPolicy(func() interface{} {
		return &stub{1}
	})

	pool := NewFixedPoolWithPolicyAndSize(policy, poolSize)
	items := make([]interface{}, poolSize)
	for i := 0; i < 1; i++ {
		items[i] = pool.Get()
	}
	for i := 0; i < 1; i++ {
		pool.Return(items[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// This is a typical scenario where we are reusing the same object from the pool
		element := pool.Get()
		pool.Return(element)

	}
}

func BenchmarkFixedPool_WithPoolHalfFull(b *testing.B) {
	type stub struct {
		value int
	}
	policy := NewDefaultPooledObjectPolicy(func() interface{} {
		return &stub{1}
	})
	pool := NewFixedPoolWithPolicyAndSize(policy, poolSize)
	items := make([]interface{}, poolSize)
	for i := 0; i < poolSize; i++ {
		items[i] = pool.Get()
	}
	for i := 0; i < poolSize/2; i++ {
		pool.Return(items[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		items := make([]interface{}, poolSize)
		for j := 0; j < poolSize; j++ {
			items[j] = pool.Get()
		}
		for j := 0; j < poolSize; j++ {
			pool.Return(items[j])
		}
		for j := 0; j < poolSize; j++ {
			items[j] = pool.Get()
		}
	}
}