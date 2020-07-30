package main
import (
	"fmt"
)
func main(){
	num  := 10
	if num % 2 == 0{
		fmt.Println("The number is EVEN")
	}else{                // its very important else use this line only otherwise error will come.
		fmt.Println("The number is ODD")
	}

	// another one 
	if num1 := 21; num1 % 2 == 0{
		fmt.Println(num1,"is EVEN")
	}else{
		fmt.Println(num1,"is ODD")
	} 

	// if else if example.
	num2 := 101
	if num2 <= 50{
		fmt.Println("number is less than equal to 50")
	}else if num2 >= 51 && num2 <= 100{
		fmt.Println("number is between 51 and 100")
	}else{
		fmt.Println("number is greater than 100")
	}
}