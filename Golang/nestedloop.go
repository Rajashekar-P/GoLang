package main
import (
	"fmt"
)
// A for loop which has another for loop inside it is called a nested for loop.
func main(){
   n:= 5
  for i:=0; i<n;i++{
	   for j:=0;j<=i;j++{
		fmt.Printf("*")
	}
	fmt.Println()
  }
}