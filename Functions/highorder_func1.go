package main
import (
	"fmt"
)
func squareSum(x int) func(int) func(int) func(int)int{

	return func(y int) func(int) func(int) int{

		return func (z int) func(int) int{

		    return func (w int) int{

			    return x * x + y * y + z * z + w * w
		}
	}
  }
}
func main(){
	fmt.Println(squareSum(5)(6)(7)(8))
}