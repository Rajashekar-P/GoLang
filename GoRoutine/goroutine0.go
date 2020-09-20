package main
import (
	"fmt"
	"time"
)
func getName(s string){
	for _,c := range s{
		fmt.Printf("%c",c)
	}
}
func getRollno(s[]int){
	for _,r := range s{
		fmt.Printf("%d",r)
	}
}
func main(){
	fmt.Println("Main Started")
	go getName("Shekar")
	go getRollno([]int{10,20,30})
	time.Sleep(time.Millisecond)
	fmt.Println("\n Main Completed")
}