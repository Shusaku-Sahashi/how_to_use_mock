package mock_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Mockの使用準備 ストラクトでmock.Mockを記述
type TriangleMock struct{
	mock.Mock
}
//mockからCalcAreaを呼び出し
func (mock *TriangleMock) CalcArea() (int, error) {
  args := mock.Called()
  return  args.Int(0), args.Error(1)
}
//mockからRandomAreaを呼び出し
func (mock *TriangleMock) RandomArea() (int, error) {
  args := mock.Called()
  return  args.Int(0), args.Error(1)
}

func TestCalcArea(t *testing.T) {
	assert := assert.New(t)
	mock := new(TriangleMock)
	service := &TriangleService {
		TInterface: mock,
	}
	expect := 10

	mock.On("CalcArea").Return(expect, nil)
	res := service.CalcArea()

	assert.Equal(res, expect)
}


func TestRandomArea(t *testing.T) {
	assert := assert.New(t)
	mock := new(TriangleMock)
	service := &TriangleService {
		TInterface: mock,
	}
	expect := 10

	mock.On("RandomArea").Return(expect, nil)
	res := service.RandomArea()

	assert.Equal(res, expect)
}