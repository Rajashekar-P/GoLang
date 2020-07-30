package main
import (
	"fmt"
)

func main(){
	var fname,lname string = "Rajashekar","Pallala"
	var a,b,c int = 1,2,3      // or we can use a,b,c := 1,2,3
	item,Price := "Mobile",9000 // := shorthand declaration if we dont know the exact var type then we can use :=
 
	// normal varibale declaration at one file and another type is there which constant
	//in constant we cannot change the variable name because it is conctant var
	var (
		name = "Rajshekar"
		email = "shekar.p4u1234@gmail.com"
		mobile = 9951880030
	)
	name  ="Raja" // normal variable can change when we update the var name but we cannot change in constant.
	
	//constant variables
	const(
		cname = "Raj"
		cemail = "shekar@gmail.com"
		cmobile = 99520080
	)
	//cname = "shekar" // error will occur when we re trying change constant varible.

	fmt.Println(fname,"",lname)
	fmt.Println(a+b+c)
	fmt.Println(item,"-",Price)

	fmt.Println(name,email,mobile)
	fmt.Println(cname,cemail,cmobile)
}