package main
import (
	"fmt"
)
func main(){
	finger := 6
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Wrong Index")
		
	}

	//Multiple Expressions in CASE:

	letter := "b"
	switch letter {
	case "a","e","i","o","u":
		fmt.Println("vowel")
	default:
		fmt.Println("Not A Vowel")
		
	}
}