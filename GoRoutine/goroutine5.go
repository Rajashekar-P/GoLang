package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)
// waitgroup
var wg sync.WaitGroup

func responseSize(url string,nums chan int){
	defer wg.Done()

	response,err := http.Get(url)
	if err != nil{
		log.Fatal(err)
	}
	defer response.Body.Close()
	body,err := ioutil.ReadAll(response.Body)

	if err != nil{
		log.Fatal(err)
	}
	nums <- len(body)  // Send values to unbuffered channel
}
func main(){
	nums:= make(chan int)  // Declare unbuffered channel
	wg.Add(1)
	go responseSize("https://facebook.com",nums)
	fmt.Println(<-nums) // read value from unbuffered channel
	wg.Wait()
	close(nums)  // closes the channel;
}