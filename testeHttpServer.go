// Thema: Wie baue ich einen Webserver in go?
// Zweck: Mögliche Anwendung für dynamische Views auf eine Datenbank über einen Browser.
//
//	Datum: 02.06.2024

package main

import (
	http "github.com/Fussbroich/lwb-WebEcho/httpserver"
)

// Zwei beispielhafte Anfragen-Bediener
func halloBediener() []byte {
	// Die Response (eine HTML-Page) für den Browser erzeugen.
	// Tipp: Benutze html/template zur dynamischen Erzeugung von HTML-Seiten und
	// fülle die HTML-Seite mit Daten.
	return []byte(`
<html>
	<body>
    	<div>Hallo!</div> <a href="/njuten">gehe zu njuten</a>
	</body>
</html>`)
}

func njutenBediener() []byte {
	return []byte(`
<html>
	<body>
    	<div>N'juuten!</div> <a href="/hallo">gehe zu hallo</a>
	</body>
</html>`)
}

func main() {
	// erzeuge ein Server-Objekt
	var srv http.HttpServer = http.New("127.0.0.1", 8080)

	// Registriere einige Methoden und Routen, die Deine WebApp bedienen soll.
	// Hier wird jeder Methode/Route ein Bediener zugeordnet.

	// HINWEIS: das funktioniert so erst mit go 1.22:
	// siehe https://eli.thegreenplace.net/2023/better-http-server-routing-in-go-122

	// Wir bedienen nur die folgende sehr einfache Benutzerschnittstelle.
	// Es wird derzeit in jedem Fall lediglich eine Webpage ausgeliefert,
	// die hallo oder n'juuten ausgibt.
	// Andere Bediener würden echte Web-Seiten erzeugen ...

	// GET / - die Pfad-Syntax, z.B. {$} wird im obigen Dokument zum neuen Routing erklärt.
	srv.BedieneGET("/{$}", halloBediener) // bediene Wurzel wie /ping
	// GET /hallo
	srv.BedieneGET("/hallo", halloBediener)
	// GET /njuten
	srv.BedieneGET("/njuten", njutenBediener)

	// Wir bedienen auch ein öffentliches Verzeichnis für Bilder und Downloads -
	// beispielsweise für das Tab-Icon und das Styling.
	srv.BedieneVerzeichnis("/", "static")

	// den Web-Server starten - beendet wird er mit Ctrl-C im Terminal
	srv.LauscheUndBediene() // Aufruf blockiert
}
