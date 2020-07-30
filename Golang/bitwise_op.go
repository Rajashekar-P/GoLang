package main
import (
	"fmt"
)
func main(){
	var a uint = 9  // 0000 1001
	var b uint = 65 // 0100 0001
	// 1 << 1 ---> 0010 --> 2
	// 1 << 2 ---> 0100 --> 4
	// 1 << 3 ---> 1000 --> 8

   //Right shift
   // 1011 >> 1 ---> 0101
   // 1011 >> 3 ---> 0001
  
   var z uint

   z = a & b 
   fmt.Println("a & b is",z)

   z = a | b 
   fmt.Println("a | b is",z)

   z = a ^ b 
   fmt.Println("a ^ b is",z)
   
   z = a >> 1 
   fmt.Println("a >> b is",z)

   z = a << 1 
   fmt.Println("a << b is",z)





}