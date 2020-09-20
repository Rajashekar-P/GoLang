package main
import (
	"fmt"
	"reflect"
)
func SampleExample(i ...interface{}){
	for _,v := range i {
		fmt.Println(v,"--",reflect.ValueOf(v).Kind())
	}
}
func main(){
	SampleExample(1,"Red",true,10.5, []string{"Foo","Bar","Baz"},map[string]int{"Apple":24,"Tomato":10})
}