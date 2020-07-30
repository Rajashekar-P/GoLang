package main
import (
	"fmt"
)
func main(){
	var a , b = 35,45

	 a = b // here we are changing value of a into b then a will become b(45)
	 fmt.Println("a=b is",a)

	 a = 35
	 a+=b //a+=b means a = a+b
	 fmt.Println("a+=b is",a)

	 a = 50
	a -= b // a-= means a = a-b
	fmt.Println("a-=b is",a)

	a = 2
	a*=b // a*= means a = a*b
	fmt.Println("a*=b is",a)

	a = 100
	a /= b // a/=b means a = a/b
	fmt.Println("a/=b is",a)

	a = 40
	a %= b // a%=b means a = a%b
	fmt.Println("a%= is",a)

}