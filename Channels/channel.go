//if goroutine is 
// ch := make(chan int)
// ch  <- X  ---- send statement
//X =  <- ch  --- a recive expression in an assignment statement

// close(ch)
//type of channels are buffered and unbuffered.
//Unbuffered := make(chan int)
// buffered := make(chan int, 10)

//goroutine1 := make(chan string, 5)
//goroutine <- "Australia"       receiving

// dat := <- goroutine1          sending
package main
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var goRoutine sync.WaitGroup

func main(){
rand.Seed(time.Now().Unix())

projects := make(chan string, 10)
goRoutine.Add(5)   // Launch 5 goroutine

for i :=1; i <=5; i++ {
	go employee(projects, i)

}
for j:=1; j<=10; j++ {
	projects <- fmt.Sprintf("Project :%d", j)
} 
close(projects)
goRoutine.Wait()
}

func employee(projects chan string, employee int) {
	defer goRoutine.Done()
	for {
		project, result := <- projects
		if result == false {
			fmt.Printf("Employee : %d : Exit\n", employee)
			return
		}
		fmt.Printf("Employee : %d :  Started: %s \n", employee, project)

		sleep := rand.Int63n(50) // Randomly wait to simulate wor time

		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Println("\n Time to sleep ", sleep, "ms\n")

	    fmt.Printf("Employee : %d: completed %s\n",employee,project)
	}
}