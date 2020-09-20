package main

import (
	"log"
	"net/smtp"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started ")
}
func main() {
	// connect remote SMTP server
	client, err := smtp.Dial("smtp.smail.cpm:45")
	if err != nil {
		log.Fatalln(err)
	}
	client.Data()
}
