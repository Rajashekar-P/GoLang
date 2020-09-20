package main
import (
	"fmt"
	"time"
)
func numbers(){
	for i:= 0; i < 5;i++{
		time.Sleep(450 * time.Millisecond)
		fmt.Printf("%d",i)
	}
	
}
func alphabets()  {
	for c:= 'a';c < 'e';c++{
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%c",c)
	}
	
}
func main(){
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("\n Main Completed")
}