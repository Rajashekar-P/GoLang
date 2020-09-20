package main
import (
	"fmt"
)
func first(){
	fmt.Println("first")
}
func second(){
	fmt.Println("second")
}
func third(){
	fmt.Println("third")
}
func fourth(){
	fmt.Println("fourth")
}
func main(){
	defer fourth()
	third()
	first()
	second()
}