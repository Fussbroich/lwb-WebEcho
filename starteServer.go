/*
	 Thema: Ich baue einen Webserver in go

	 Zweck: Erzeuge und veröffentliche dynamische Web-Seiten mit Deinen Daten aus einem go-Programm.

	 HINWEIS: benötigt go 1.22 im Modulmodus:

	 1. U-Variable setzen
		GO111MODULES=auto
	 2. altes go löschen:
		sudo rm -rf /usr/lib/golang
	 3. neue Version runterladen, dann entpacken
		sudo tar -C /usr/lib/ -xzf go1.22.4.linux-amd64.tar.gz
	 4. Verzeichnis umbenennen
		sudo mv /usr/lib/go /usr/lib/golang
	 5. fehlende Module laden:
		go mod tidy
	 6. Server compilieren und starten
		go run starteServer.go

	 Autor: T. Schrader

	 Datum: 02.06.2024
*/
package main

import (
	"fmt"
	"strings"

	http "github.com/Fussbroich/lwb-WebEcho/httpserver"
)

// Etwas HTML für die Anzeige im Browser. Regel: Strukturiere Deinen HTML-Inhalt
// mit div-Elementen. Die "class" bezieht sich auf einen Styling-Eintrag
// in der style.css-Datei im öffentlichen Verzeichnis.
// So hast Du eine gute Kontrolle über das Design der Darstellung
// und das Html bleibt übersichtlich. Mehr zu HTML und CSS lernst Du unter
//
// https://wiki.selfhtml.org/wiki/HTML
// https://www.w3schools.com/html/
const grussSeite = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
	    <link rel="stylesheet" href="/style.css">
		<title>mini Web-Projekt</title>
	</head>
	<body>
	<!-- ein div für die Navigation -->
	<div class="navigation">
		<a href="/gruss">grüße normal</a> |
		<a href="/gruss/berlin">grüße berlinerisch</a> |
		<a href="/stress">provoziere Stress</a>
	</div>
	<!-- ein div für den Inhalt -->
	<div class="inhalt">{{Gruss}}</div>
</body>
</html>`

// Drei beispielhafte Anfragen-Bediener, die den Inhalt
// für die Serverantwort erzeugen: Hier wird HTML produziert.
// Wir benutzen eine Vorlage zur Erzeugung von HTML-Seiten und
// ersetzen die Parameter mit Daten. Dabei können Fehler passieren,
// die natürlich vom Server behandelt werden müssen.
//
// Tipp: Du kannst das Html auch ohne Vorlage direkt erzeugen.
//
// Der Server liefert das hier erzeugte Html an den Browser.
func normalBediener() ([]byte, error) {
	var html string
	// Hier wird die Vorlage mit "Daten" gefüllt.
	html = strings.Replace(grussSeite, "{{Gruss}}", "Hallo", -1)
	// ... es hat beim Befüllen keinen Fehler gegeben:
	return []byte(html), nil
}

func berlinBediener() ([]byte, error) {
	var html string
	html = strings.Replace(grussSeite, "{{Gruss}}", "N'juuten", -1)
	return []byte(html), nil
}

// erzeugt absichtlich einen Server-Fehler
func stressBediener() ([]byte, error) {
	// ... hier passiert ein Fehler
	var err error
	err = fmt.Errorf("Wat willste?")
	return nil, err
}

func main() {
	// erzeuge ein Server-Objekt
	var srv http.HttpServer = http.NewHttpServer("127.0.0.1", 8080)

	// Wir veröffentlichen zunächst ein Verzeichnis für statische HTML-Seiten,
	// Bilder, Downloads und CSS. In diesem Beispielprojekt haben wir auch
	// eine "index.html" in dem Verzeichnis hinterlegt.
	// Diese Datei findet der Browser automatisch, wenn man diesen Url ("/") anfragt:
	srv.VeroeffentlicheVerzeichnis("/", "static")

	// Registriere jetzt einige Pfade, die Deine WebApp bedienen soll.
	// Hier wird jeder Kombination aus <Http-Methode> und <Url-Pfad> ein Bediener zugeordnet.

	// Wir bedienen die folgende sehr einfache Schnittstelle:

	// Bediene "GET /gruss"
	srv.SetzeHtmlBediener(http.MethodeGet, "/gruss", normalBediener)
	// Bediene "GET /gruss/berlin"
	srv.SetzeHtmlBediener(http.MethodeGet, "/gruss/berlin", berlinBediener)
	// Bediene "GET /stress"
	srv.SetzeHtmlBediener(http.MethodeGet, "/stress", stressBediener)

	// Starte den Web-Server - jetzt läuft er und kann über einen Browser oder
	// über andere Client-Programme (z.B. curl oder das eigene) erreicht werden.
	//
	// Man fährt den Server herunter mit der Tastenkombination Ctrl-C in der Konsole.
	srv.LauscheUndBediene()
	// Code, der danach kommt, wird erst ausgeführt, wenn der Server wieder heruntergefahren ist
	fmt.Println("Tschüss")
}
