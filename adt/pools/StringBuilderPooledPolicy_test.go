package pools

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type StringBuilderPooledPolicyTestSuite struct {
	suite.Suite
}

func TestStringBuilderPooledPolicy(t *testing.T) {
	suite.Run(t, new(StringBuilderPooledPolicyTestSuite))
}

func (suite *StringBuilderPooledPolicyTestSuite) TestCreateAndReturn() {
	// Arrange
	target := &StringBuilderPooledPolicy{}

	// Act
	actual := target.Create().(*strings.Builder)
	actual.WriteString("Hello World")
	target.Return(actual)

	// Assert
	assert.Equal(suite.T(), 0, actual.Len())
}

func ExampleStringBuilderPooledPolicy() {
	var policy PooledObjectPolicy
	var pool Pool

	policy = NewStringBuilderPooledPolicyWithCapacity(1024)
	pool = NewFixedPoolWithPolicy(policy)

	builder := pool.Get().(*strings.Builder)
	builder.WriteString("Hello World")
	pool.Return(builder)

	builder2 := pool.Get().(*strings.Builder)
	if builder2 == builder {
		fmt.Println("Same Builder")
	}
	builder2.WriteString("Bye World")
	pool.Return(builder2)

	//Output:Same Builder
}