package main
import (
	"fmt"
)
func main(){
	var action int

	fmt.Println("Enter 1 fo Student and 2 for Professional:")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Println("I am Student")
	case 2:
		fmt.Println("I am Professional")
	default:
		panic(fmt.Sprintf("I am a %d",action))
		
	}
	fmt.Println("")
	fmt.Println("Enter 1 For US and 2 For UK:")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Println("US")
	case 2:
		fmt.Println("UK")
	default:
		panic(fmt.Sprintf("Iam A %d",action))
		
	}
}