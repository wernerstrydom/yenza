package pools

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/wernerstrydom/yenza/adt/pools/mocks"
	"runtime"
	"testing"
)

type poolable struct {
	x int
}


type LeakTrackingPoolTestSuite struct {
	suite.Suite
}

func TestLeakTrackingPoolTestSuite(t *testing.T) {
	suite.Run(t, new(LeakTrackingPoolTestSuite))
}

func (suite *LeakTrackingPoolTestSuite) TestGet() {
	// Arrange
	inner := new(mocks.Pool)
	inner.On("Get").Return(&poolable{})
	target := NewLeakTrackingPool(inner)

	// Act
	actual := target.Get()

	// Assert
	assert.NotNil(suite.T(), actual)
	actual = nil
	runtime.GC()

	inner.AssertExpectations(suite.T())
}

func (suite *LeakTrackingPoolTestSuite) TestLeakTrackingPool_Return() {
	// Arrange
	element := &poolable{}
	inner := new(mocks.Pool)
	inner.On("Get").Return(&poolable{})
	inner.On("Return", element)
	target := NewLeakTrackingPool(inner)
	item := target.Get()

	// Act
	target.Return(item)

	// Assert
	assert.NotNil(suite.T(), item)
	inner.AssertExpectations(suite.T())
}

func (suite *LeakTrackingPoolTestSuite) TestNewLeakTrackingPool() {
	// Arrange
	inner := new(mocks.Pool)

	// Act
	actual := NewLeakTrackingPool(inner)

	// Assert
	assert.NotNil(suite.T(), actual)
	assert.NotNil(suite.T(), actual.trackers)
	assert.Same(suite.T(), inner, actual.inner)
}
