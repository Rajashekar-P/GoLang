package main
import (
	"fmt"
)
//arrays are the collection of elements that belong to same type.ex integers 3,4,5,.. form an array.
//mixng values of different type ,for ex an array that conatins ints and strings is not allowed in Go.
// array declaration [n]T. Here T is type of elements,n is number of elements
func main(){
	var a[3]int // int array with length of 3
	//fmt.Println(a) we didnot give elements so array values assigned to 0

	a[0] = 12 // index starts at 0
	a[1] = 78
	a[2] = 50
	
	var b = [3]int{12,14,16}//another way to declare an array.
	
	fmt.Println(a)
	fmt.Println(b)
	
	fmt.Println(a[1])// accessing the 2nd index value.
	

	// shorthand declaration in arrays
	var b =[4]int{12,78,5,100} 
	for i,v:= range b{
     
		 fmt.Println(i,v)
     }
}
