package main

import (
	"fmt"

	client "github.com/Fussbroich/lwb-WebEcho/httpclients"
	server "github.com/Fussbroich/lwb-WebEcho/httpserver"
)

// kleines Programm zum Testen Deines Servers
func main() {

	var client = client.NewHttpClient("127.0.0.1", 8080)
	var url1 = "/gruss"
	var url2 = "/stress"

	var ant []byte
	var fehler error

	fmt.Printf("Sende GET-Anfrage für %s\n", url1)
	ant, fehler = client.Anfragen(server.MethodeGet, url1)
	if fehler != nil {
		fmt.Printf("Fehler: %v\n", fehler)
	}
	fmt.Printf("Antwort vom Server: %s\n", string(ant))
	fmt.Println("*********************************************")

	fmt.Printf("Sende GET-Anfrage für %s\n", url2)
	ant, fehler = client.Anfragen(server.MethodeGet, url2)
	if fehler != nil {
		fmt.Printf("Fehler: %v\n", fehler)
	}
	fmt.Printf("Antwort vom Server: %s\n", string(ant))

}
