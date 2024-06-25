// Thema: Wie baue ich einen Webserver in go?
// Zweck: Mögliche Anwendung für dynamische Views auf eine Datenbank über einen Browser.
//
//	Datum: 02.06.2024

package main

import (
	"fmt"

	http "github.com/Fussbroich/lwb-WebEcho/httpserver"
)

// Drei beispielhafte Anfragen-Bediener, die
// Response-Bodies (eine HTML-Seite) für den Browser erzeugen.
// Tipp: Benutze html/template zur dynamischen Erzeugung von HTML-Seiten und
// fülle die HTML-Seite dynamisch mit Daten.
func normalBediener() ([]byte, error) {
	var content = `<html>
		<body>
		   	<div>Hallo! <a href="/">home</a></div>
		</body>
	</html>`
	return []byte(content), nil
}

func berlinBediener() ([]byte, error) {
	var content = `<html>
		<body>
    		<div>N'juuten! <a href="/">home</a></div>
		</body>
	</html>`
	return []byte(content), nil
}

func stressBediener() ([]byte, error) {
	return nil, fmt.Errorf("wat willste?")
}

func main() {
	// erzeuge ein Server-Objekt
	var srv http.HttpServer = http.New("127.0.0.1", 8080)

	// Registriere einige Pfade, die Deine WebApp bedienen soll.
	// Hier wird jeder Kombination aus Http-Methode und Pfad ein Bediener zugeordnet.

	// HINWEIS: das funktioniert so erst mit go 1.22:
	// siehe https://eli.thegreenplace.net/2023/better-http-server-routing-in-go-122

	// Wir bedienen nur die folgende sehr einfache Benutzerschnittstelle.
	// Es wird eine HTML-Seite ausgeliefert,
	// die hallo oder n'juuten ausgibt, oder es gibt Stress.

	// Bediene "GET /gruss"
	srv.SetzeBediener("GET /gruss", normalBediener)
	// Bediene "GET /gruss/berlin"
	srv.SetzeBediener("GET /gruss/berlin", berlinBediener)
	// Bediene "GET /stress"
	srv.SetzeBediener("GET /stress", stressBediener)

	// Wir veröffentlichen auch ein Verzeichnis für statische HTML-Seiten,
	// Bilder, Downloads und CSS. Bediene "GET /" muss damit auch nicht sein, denn
	// hier liegt eine "index.html" - die findet der Browser automatisch.
	srv.VeroeffentlicheVerzeichnis("/", "static")

	// den Web-Server starten - beendet wird er mit Ctrl-C im Terminal
	srv.LauscheUndBediene() // Aufruf blockiert, bis Server beendet wird
}
