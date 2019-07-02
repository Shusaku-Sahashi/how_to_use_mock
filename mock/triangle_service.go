package mock_test

type TriangleService struct {
	TInterface TriangleInterface
}

func NewGetTriangleService(height,bottom int) TriangleService {
  return TriangleService {
     TInterface : NewTriangle(height,bottom,),
  }
}
 
//triangleInterfaceのCalcAreaを返却
func (service *TriangleService) CalcArea() int{
  area, err := service.TInterface.CalcArea()
  if err != nil {
    panic("error")
  }
  return area
}
//triangleInterfaceのRandomAreaを返却
func (service *TriangleService) RandomArea() int{
  area, err := service.TInterface.RandomArea()
  if err != nil {
    panic("error")
  }
  return area
}