package main
import (
	"fmt"
)
func sampleExample(s ...string){
	fmt.Println(s)
}
func main(){
	sampleExample()
	sampleExample("Green","Red")
	sampleExample("Red","Blue","Shekar","Raja")
	sampleExample("Red","White","Blue")

}