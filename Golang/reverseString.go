package main
import (
	"fmt"
)
func main(){
	S:= []int{1,2,3,4,5,6}

	for i:= len(S) - 1;i >= 0;i--{
		fmt.Print(S[i])
		if i > 0{
			fmt.Print(",")
		}
	}
	fmt.Println()
}