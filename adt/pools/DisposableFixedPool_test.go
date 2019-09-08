package pools

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	mocks2 "github.com/wernerstrydom/yenza/adt/mocks"
	"github.com/wernerstrydom/yenza/adt/pools/mocks"
	"runtime"
	"testing"
)

type DisposableFixedPoolTestSuite struct {
	suite.Suite
}

func TestDisposableFixedPool(t *testing.T) {
	suite.Run(t, new(DisposableFixedPoolTestSuite))
}

func (suite *DisposableFixedPoolTestSuite) TestGet_WithNothingPooled() {
	// Arrange
	type stub struct {
		value int
	}

	expected := &stub{5}
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Create").Return(expected).Once()
	items := make([]element, 3)
	target := DisposableFixedPool{
		FixedPool: FixedPool{
			items:           items,
			policy:          policy,
			isDefaultPolicy: false,
			firstItem:       nil,
		},
	}

	// Act
	actual := target.Get()

	// Assert
	assert.NotNil(suite.T(), actual, "Expected a valid item")
	assert.Same(suite.T(), expected, actual)

	policy.AssertExpectations(suite.T())
}

func (suite *DisposableFixedPoolTestSuite) TestGet_WithSingleItem() {
	// Arrange
	var expected = new(mocks2.Disposable)
	policy := new(mocks.PooledObjectPolicy)

	items := make([]element, 3)
	target := DisposableFixedPool{
		FixedPool: FixedPool{
			items:           items,
			policy:          policy,
			isDefaultPolicy: false,
			firstItem:       expected,
		},
	}

	// Act
	actual := target.Get()

	// Assert
	assert.NotNil(suite.T(), actual, "Expected a valid item")
	assert.Same(suite.T(), expected, actual)

	policy.AssertExpectations(suite.T())
}

func (suite *DisposableFixedPoolTestSuite) TestGet_WithFragmentation1() {
	// Arrange
	type stub struct {
		value int
	}

	var expected = new(mocks2.Disposable)
	policy := new(mocks.PooledObjectPolicy)

	items := make([]element, 3)
	items[0] = element{ value: expected }
	target := DisposableFixedPool{
		FixedPool: FixedPool{
			items:           items,
			policy:          policy,
			isDefaultPolicy: false,
			firstItem:       nil,
		},
	}

	// Act
	actual := target.Get()

	// Assert
	assert.NotNil(suite.T(), actual, "Expected a valid item")
	assert.Same(suite.T(), expected, actual)

	policy.AssertExpectations(suite.T())
	expected.AssertExpectations(suite.T())
}

func (suite *DisposableFixedPoolTestSuite) TestGet_WithFragmentation2() {
	// Arrange
	type stub struct {
		value int
	}

	var expected = new(mocks2.Disposable)
	policy := new(mocks.PooledObjectPolicy)

	items := make([]element, 3)
	items[1] = element{ value: expected }
	target := DisposableFixedPool{
		FixedPool: FixedPool{
			items:           items,
			policy:          policy,
			isDefaultPolicy: false,
			firstItem:       nil,
		},
	}

	// Act
	actual := target.Get()

	// Assert
	assert.NotNil(suite.T(), actual, "Expected a valid item")
	assert.Same(suite.T(), expected, actual)

	policy.AssertExpectations(suite.T())
	expected.AssertExpectations(suite.T())
}

func (suite *DisposableFixedPoolTestSuite) TestReturn_WithPolicyReturnsFalse() {
	// Arrange
	type stub struct {
		value int
	}

	var item = new(mocks2.Disposable)
	item.On("Dispose").Once()
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Return", item).Return(false).Once()
	items := make([]element, 3)
	target := DisposableFixedPool{
		FixedPool: FixedPool{
			items:           items,
			policy:          policy,
			isDefaultPolicy: false,
			firstItem:       nil,
		},
	}

	// Act
	target.Return(item)

	// Assert
	assert.Nil(suite.T(), target.firstItem)
	assert.Nil(suite.T(), target.items[0].value)
	assert.Nil(suite.T(), target.items[1].value)
	assert.Nil(suite.T(), target.items[2].value)

	policy.AssertExpectations(suite.T())
	item.AssertExpectations(suite.T())
}

func (suite *DisposableFixedPoolTestSuite) TestReturn_WithEmptyPool() {
	// Arrange
	type stub struct {
		value int
	}

	s1 := new(mocks2.Disposable)
	s2 := new(mocks2.Disposable)
	s3 := new(mocks2.Disposable)
	s4 := new(mocks2.Disposable)
	item := new(mocks2.Disposable)

	policy := new(mocks.PooledObjectPolicy)
	policy.On("Return", item).Return(true).Once()
	items := make([]element, 3)
	target := DisposableFixedPool{
		FixedPool: FixedPool{
			items:           items,
			policy:          policy,
			isDefaultPolicy: false,
			firstItem:       nil,
		},
	}

	// Act
	target.Return(item)

	// Assert
	//var intf interface{} = item
	//assert.Same(suite.T(), intf, target.firstItem)
	assert.Nil(suite.T(), target.items[0].value)
	assert.Nil(suite.T(), target.items[1].value)
	assert.Nil(suite.T(), target.items[2].value)

	policy.AssertExpectations(suite.T())
	s1.AssertExpectations(suite.T())
	s2.AssertExpectations(suite.T())
	s3.AssertExpectations(suite.T())
	s4.AssertExpectations(suite.T())
	item.AssertExpectations(suite.T())
}

func (suite *DisposableFixedPoolTestSuite) TestReturn_WithOnePooledObject() {
	// Arrange
	s1 := new(mocks2.Disposable)
	s2 := new(mocks2.Disposable)
	s3 := new(mocks2.Disposable)
	s4 := new(mocks2.Disposable)
	var item = new(mocks2.Disposable)
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Return", item).Return(true).Once()
	items := make([]element, 3)
	target := DisposableFixedPool{
		FixedPool: FixedPool{
			items:           items,
			policy:          policy,
			isDefaultPolicy: false,
			firstItem:       item,
		},
	}

	// Act
	target.Return(item)

	// Assert
	var intf interface{} = item
	assert.Same(suite.T(), intf, target.firstItem, "expected first item to the item that was returned")
	assert.Same(suite.T(), intf, target.items[0].value)
	assert.Nil(suite.T(), target.items[1].value)
	assert.Nil(suite.T(), target.items[2].value)

	policy.AssertExpectations(suite.T())
	s1.AssertExpectations(suite.T())
	s2.AssertExpectations(suite.T())
	s3.AssertExpectations(suite.T())
	s4.AssertExpectations(suite.T())
	item.AssertExpectations(suite.T())
}

func (suite *DisposableFixedPoolTestSuite) TestReturn_WithTwoPooledObjects() {
	// Arrange
	type stub struct {
		value int
	}
	s1 := new(mocks2.Disposable)
	s2 := new(mocks2.Disposable)
	s3 := new(mocks2.Disposable)
	s4 := new(mocks2.Disposable)
	var item = new(mocks2.Disposable)
	policy := new(mocks.PooledObjectPolicy)
	policy.On("Return", item).Return(true).Once()
	items := make([]element, 3)
	items[0] = element{s4}
	target := DisposableFixedPool{
		FixedPool: FixedPool{
			items:           items,
			policy:          policy,
			isDefaultPolicy: false,
			firstItem:       s3,
		},
	}

	// Act
	target.Return(item)

	// Assert
	assert.Same(suite.T(), s3, target.firstItem, "expected first item to the item that was returned")
	assert.Same(suite.T(), s4, target.items[0].value)
	assert.Same(suite.T(), item, target.items[1].value)
	assert.Nil(suite.T(), target.items[2].value)

	policy.AssertExpectations(suite.T())
	s1.AssertExpectations(suite.T())
	s2.AssertExpectations(suite.T())
	s3.AssertExpectations(suite.T())
	s4.AssertExpectations(suite.T())
	item.AssertExpectations(suite.T())
}

func (suite *DisposableFixedPoolTestSuite) TestReturn_WithFullPool_DiscardsReturnedObject() {
	// Arrange
	s1 := new(mocks2.Disposable)
	s2 := new(mocks2.Disposable)
	s3 := new(mocks2.Disposable)
	s4 := new(mocks2.Disposable)

	var item = new(mocks2.Disposable)
	item.On("Dispose").Once()

	policy := new(mocks.PooledObjectPolicy)
	policy.On("Return", item).Return(true).Once()
	items := make([]element, 3)
	items[2] = element{s4}
	items[1] = element{s2}
	items[0] = element{s1}
	target := DisposableFixedPool{
		FixedPool: FixedPool{
			items:           items,
			policy:          policy,
			isDefaultPolicy: false,
			firstItem:       s3,
		},
	}

	// Act
	target.Return(item)

	// Assert
	assert.Same(suite.T(), s3, target.firstItem)
	assert.Same(suite.T(), s1, target.items[0].value)
	assert.Same(suite.T(), s2, target.items[1].value)
	assert.Same(suite.T(), s4, target.items[2].value)

	policy.AssertExpectations(suite.T())
	s1.AssertExpectations(suite.T())
	s2.AssertExpectations(suite.T())
	s3.AssertExpectations(suite.T())
	s4.AssertExpectations(suite.T())
	item.AssertExpectations(suite.T())
}



func (suite *DisposableFixedPoolTestSuite)  TestNewObjectPoolWithPolicy() {
	// Arrange
	policy := new(mocks.PooledObjectPolicy)

	// Act
	actual := NewDisposableObjectPool(policy)

	// Assert
	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), runtime.NumCPU()*2-1, len(actual.items))
	assert.Nil(suite.T(), actual.firstItem)
	assert.Same(suite.T(), policy, actual.policy)
	assert.False(suite.T(), actual.isDefaultPolicy)

	policy.AssertExpectations(suite.T())
}

func (suite *DisposableFixedPoolTestSuite)  TestNewObjectPoolWithPolicyAndSize() {
	// Arrange
	policy := new(mocks.PooledObjectPolicy)

	// Act
	actual := NewDisposableObjectPoolWithPolicyAndSize(policy, 16)

	// Assert
	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), 15, len(actual.items))
	assert.Nil(suite.T(), actual.firstItem)
	assert.Same(suite.T(), policy, actual.policy)
	assert.False(suite.T(), actual.isDefaultPolicy)

	policy.AssertExpectations(suite.T())
}

func (suite *DisposableFixedPoolTestSuite)   TestDispose() {
	// Arrange
	item0 := new(mocks2.Disposable)
	item0.On("Dispose").Once()
	item1 := new(mocks2.Disposable)
	item1.On("Dispose").Once()
	item2 := new(mocks2.Disposable)
	item2.On("Dispose").Once()
	item3 := new(mocks2.Disposable)
	item3.On("Dispose").Once()

	items := make([]element, 3)
	items[2] = element{item1}
	items[1] = element{item2}
	items[0] = element{item3}
	policy := new(mocks.PooledObjectPolicy)
	//policy.On("Return", item).Return(true).Once()
	target := DisposableFixedPool{
		FixedPool: FixedPool{
			items:           items,
			policy:          policy,
			isDefaultPolicy: false,
			firstItem:       item0,
		},
	}

	// Act
	target.Dispose()

	// Assert
	policy.AssertExpectations(suite.T())
	item0.AssertExpectations(suite.T())
	item1.AssertExpectations(suite.T())
	item2.AssertExpectations(suite.T())
	item3.AssertExpectations(suite.T())
}



