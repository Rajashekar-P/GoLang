package main
import (
	"fmt"
	"unsafe"
)
func main(){
	var a int = 89
	b:= 95
	fmt.Println("value of a is",a,"and b is",b)
	fmt.Printf("type of a is %T,size of a is %d",a,unsafe.Sizeof(a)) // type and size of a
	fmt.Printf("\n type of b is %T,size of b is %d\n",b,unsafe.Sizeof(b))// type and size of b

	//complex numbers
	c1:= complex(5,7)
	c2:= 8+27i
	cadd := c1+c2
	fmt.Println("sum:",cadd)
	cmul := c1 * c2
	fmt.Println("mul:",cmul)
}