package main
import (
	"fmt"
)
func main(){
	var action int
	
	fmt.Println("Enter 1 For Student 2 For Professional:")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Println("I am A Student")
	case 2 :
		fmt.Println("I am A Professional")
	default:
		panic(fmt.Sprintf("I am A %d",action))
		
	}
	defer func(){
		action := recover()
		fmt.Println(action)
	}()
}