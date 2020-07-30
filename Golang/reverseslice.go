package main
import (
	"fmt"
)
func main(){

s := []int{5, 2, 6, 3, 1, 4}

for i := len(s) - 1; i >= 0; i-- {
    fmt.Print(s[i])
    if i > 0 {
        fmt.Print(", ")
    }
}
fmt.Println()
}