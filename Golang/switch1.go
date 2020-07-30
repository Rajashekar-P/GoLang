package main
import (
	"fmt"
)


func main(){
	num:= 150
	switch  {// expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println("number is greater than 0 and  less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("number is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("number is greater than 100")
	}
}