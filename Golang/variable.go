package main
import (
	"fmt"
)
func main(){
	var i int = 20
	var s string = "Shekar"
	var f float64 = 20.5
	fmt.Println(i,s,f)
	fmt.Printf("i is %d\ns is %s\nf is %.2f\n",i,s,f)
	fmt.Printf("Type of i is %T\nType of s is %T\nType of f is %T",i,s,f)
	
}