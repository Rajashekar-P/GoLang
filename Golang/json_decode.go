package main
import (
	"fmt"
	"encoding/json"
)
// JSON Contains a sample String to unmarshal.
 
var JSON = `{
    "name":"Rajashekar P",
    "jobtitle":"Software Developer",
    "phone":{
        "home":"123-456-789",
        "mobile":"9951880030"
    },
    "email":"shekar@gmail.com"
}`
 
func main() {
    // Unmarshal the JSON string into info map variable.
    var info map[string]interface{}
    json.Unmarshal([]byte(JSON),&info)
 
    // Print the output from info map.
    fmt.Println(info["name"])
    fmt.Println(info["jobtitle"])
    fmt.Println(info["email"])
    fmt.Println(info["phone"].(map[string]interface{})["home"]) 
    fmt.Println(info["phone"].(map[string]interface{})["mobile"])
}