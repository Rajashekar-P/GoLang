package main
import (
	"fmt"
	"time"
)
func main() {
	start := time.Now()
	go func()  {
	  for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
		
	}()
	go func()  {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
		
	}()
	elaspedTime := time.Since(start)
	fmt.Println("Total time For execution:" + elaspedTime.String())
	time.Sleep(time.Second)
	
}