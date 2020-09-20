package main
import (
	"fmt"
)
//returning multiple values.
func rectangle(l int,b int)(area int, perimeter int)  {
	area = l * b
	perimeter = 2 * (l + b)
	return
}
func main(){
    var a, p int
    a, p = rectangle(20,30)
 //fmt.Printf("Area:%d Perimeter:%d",a,p)
     fmt.Println("Area:",a)
     fmt.Println("Perimeter:",p)
}
/*
1. A name must begin with a letter, and can have any number of additional letters and numbers.
2. A function name cannot start with a number.
3. A function name cannot contain spaces.
4. If the functions with names that start with an uppercase letter will be exported to other packages. 
   If the function name starts with a lowercase letter, it won't be exported to other packages, 
   but you can call this function within the same package.
5. If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
6. function names are case-sensitive (car, Car and CAR are three different variables).
*/