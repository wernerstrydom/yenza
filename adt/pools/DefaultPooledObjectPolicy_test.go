package pools

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DefaultPooledObjectPolicyTestSuite struct {
	suite.Suite
}

func TestDefaultPooledObjectPolicy(t *testing.T) {
	suite.Run(t, new(DefaultPooledObjectPolicyTestSuite))
}

func (suite *DefaultPooledObjectPolicyTestSuite) TestCreate() {
	type stub struct {
		n int
	}

	// Arrange
	var called bool
	var expected interface{} = &stub{3}
	target := &DefaultPooledObjectPolicy{
		factory: func() interface{} {
			called = true
			return expected
		},
	}

	// Act
	actual := target.Create()

	// Assert
	assert.True(suite.T(), called)
	assert.Same(suite.T(), expected, actual)
}

func (suite *DefaultPooledObjectPolicyTestSuite) TestReturn() {
	type stub struct {
		n int
	}

	// Arrange
	var item interface{} = &stub{5}
	target := &DefaultPooledObjectPolicy{
		factory: func() interface{} {
			return nil
		},
	}

	// Act
	actual := target.Return(item)

	// Assert
	assert.Equal(suite.T(), true, actual)
}

func (suite *DefaultPooledObjectPolicyTestSuite)  TestNewDefaultPooledObjectPolicy() {

	// Arrange
	f := func() interface{} {
		return nil
	}

	// Act
	actual := NewDefaultPooledObjectPolicy(f)

	// Assert
	assert.NotNil(suite.T(), actual)
}