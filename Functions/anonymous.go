package main
import (
	"fmt"
)
//An Anonymous function is that was declared without any named identifier to refer to it.
var (
	area = func(l int, b int)int  {
	 return l * b
}
)
func main(){
	fmt.Println(area(20,30))
}