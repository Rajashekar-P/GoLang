package main
import (
	"fmt"
)
func main(){

	a:=[...]float64{10.5,12.5,20.5,15.4}
	
	for i:=0;i < len(a);i++{
        fmt.Printf("%d the element of a is %.2f\n",i,a[i])
	}

}