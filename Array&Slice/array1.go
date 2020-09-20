package main
import (
	"fmt"
)
func main(){
	var a[4]int
	a[0] = 4
	a[1] = 5
	a[2] = 6
	a[3] = 7

	b:=[...]int{1,2,3,4,5,6}// if give[...] compiler will automarically understand array length.
	
	fmt.Println(a)
	fmt.Println(b)
}