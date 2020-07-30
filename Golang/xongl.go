package main 
import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/labstack/echo"
)

type User struct {
	Name string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Clientid int `json:"clientid" form:"clientid`
	Responsetype `json:"responsetype" form:"responsetype"`
	Action string `json:"action" form:"action"`
}
type Users []User
var users Users

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK,users)
}

func main()  {
	fmt.Println("Server Starting.....")
	e := echo.New()
	e.Get("/users",getUsers)
	e.Start(":8080")
	
}