package main
import (
	"fmt"
)
func main(){
	country:=[...]string{"India","United States","United Kingdom","France","Germany"}

	country1:=country // copying one array to another array.

	country1[0]="Singapore"

	fmt.Println("Country is:",country)
	fmt.Println("Country1 is:",country1)
}