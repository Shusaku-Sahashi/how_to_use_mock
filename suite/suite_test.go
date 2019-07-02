package suite_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CountTestSuite struct {
	suite.Suite
	count     int
	overCount int
}

// 各テストの実行前に呼び出されます。
func (suite *CountTestSuite) SetupTest() {
	suite.count = 0
}

// 各テスト終了時に呼び出されます。
func (suite *CountTestSuite) TearDownTest() {
	fmt.Println(suite.count)
	fmt.Println(suite.overCount)
}

func (suite *CountTestSuite) TestCountUp() {
	assert.Equal(suite.T(), 0, suite.count)
	suite.count += suite.count + 10
	assert.Equal(suite.T(), 10, suite.count)
}

func TestExpectedSuite(t *testing.T) {
	suite.Run(t, new(CountTestSuite))
}
