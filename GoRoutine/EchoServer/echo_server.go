package main
import (
	"fmt"
	"log"
	"net"
	"time"
)
// in 
func echo(c net.Conn,shout string,delay time.Duration){
	fmt.Fprintln(c,"\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c,"\t",shout)
	time.Sleep(delay)
	fmt.Fprintln(c,"\t",strings.ToLower(shout))

}
func handleconn(c net.Conn)  {
	input := bufio.NewScanner(c)
	for input.Scan(){
		echo(c,input.Text(),1 *time.Second)
	}
	c.Close()
}
func main() {
    conn, err := net.Dial("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    go mustCopy(os.Stdout, conn)
    mustCopy(conn, os.Stdin)
}