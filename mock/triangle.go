package mock_test

import (
	"math/rand"
)

//トライアングルの構造体の作成
type Triangle struct {
	bottom int
	height int
}

func NewTriangle(height, bottom int) TriangleInterface {
	t := Triangle{height: height, bottom: bottom}
	var triangle TriangleInterface = &t
	return triangle
}

//インターフェース　CalcArea RandomArea
type TriangleInterface interface {
	CalcArea() (int, error)
	RandomArea() (int, error)
}

//トライアングルに関数を作成
func (triangle *Triangle) CalcArea() (int, error) {
	return triangle.bottom * triangle.height / 2, nil
}
func (triangle *Triangle) RandomArea() (int, error) {
	return rand.Int() * rand.Int() / 2, nil
}
