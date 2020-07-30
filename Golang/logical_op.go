package main
import (
	"fmt"
)
func main(){
	var a,b,c = 10,20,30
	fmt.Println(a<b && a>c) // if one is true then result is false.if both are true then result will true.
	fmt.Println(a<b || a>c) // if one true then result true.if both false the result will false.
	fmt.Println(a==b && a>c) // if both true result will true only.
}
/* 
AND LOGICAL OPERATOR                                 OR LOGICAL OPERATOR
	 A         B         RESULT                    A            B          RESULT
   FALSE     FALSE       FALSE                    FALSE       FALSE        FALSE
   FALSE     TRUE        FALSE                    FALSE       TRUE         TRUE
   TRUE      FALSE       FALSE                    TRUE        FALSE        TRUE
   TRUE      TRUE        TRUE                     TRUE        TRUE         TRUE

*/