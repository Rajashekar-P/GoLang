package main
import (
	"fmt"
)
//strings are a collection of bytes in goalng .its alright if this defination doesnt make sense.
// for now string to be collection of charactors.
func main(){
	first := "Raja"
	last := "shekar"
	name := first+""+last// strings can be concatenated using + operator.
	fmt.Println("My name is",name)

	//type conversion
	i := 55     // integer type
	j := 45.56  // float type(float64)
	sum := i + int(j) // int + float not allowed ..so we fix that error by converting j into integer type
	fmt.Println("the sum is ",sum)

}