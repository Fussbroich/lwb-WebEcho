// Thema: Ich baue einen Webserver in go
//
// Zweck: Erzeuge und veröffentliche dynamische Web-Seiten mit Deinen Daten aus einem go-Programm.
//
// HINWEIS: benötigt go 1.22 im Modulmodus
//
// Autor: T. Schrader
//
// Datum: 02.06.2024

package main

import (
	"fmt"

	html "github.com/Fussbroich/lwb-WebEcho/htmlvorlagen"
	http "github.com/Fussbroich/lwb-WebEcho/httpserver"
)

// Etwas HTML für die Anzeige im Browser. Regel: Strukturiere Deinen HTML-Inhalt
// mit div-Elementen. Die "class" bezieht sich auf einen Styling-Eintrag
// in der style.css-Datei im öffentlichen Verzeichnis.
// So hast Du eine gute Kontrolle über das Design der Darstellung.
// Mehr zu HTML und CSS lernst Du unter
// https://wiki.selfhtml.org/wiki/HTML
// https://www.w3schools.com/html/
var grussSeite = html.NewVorlage(`
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<!-- Trick: nach Änderungen zähle die Versionsnummer v=... beliebig hoch -->
	    <link rel="stylesheet" href="/style.css?v=1.2">
		<title>{{$.Titel}}</title>
	</head>
	<body>
	<!-- ein div für die Navigation -->
	<div class="navigation">
		<a href="/gruss">grüße normal</a> |
		<a href="/gruss/berlin">grüße berlinerisch</a> |
		<a href="/stress">provoziere Stress</a>
	</div>
	<!-- ein div für den Inhalt -->
	<div class="inhalt">{{$.Gruss}}</div>
</body>
</html>`)

// Drei beispielhafte Anfragen-Bediener, die den Inhalt
// für die Serverantwort erzeugen: Hier wird HTML produziert.
// Wir benutzen eine HtmlVorlage zur Erzeugung von HTML-Seiten und
// füllen die Parameter mit Daten.
// (Tipp: Du kannst das Html auch ohne Vorlage selbst erzeugen.)
//
// Der Server liefert das hier erzeugte Html an den Browser.
func normalBediener() ([]byte, error) {
	grussSeite.SetzeParameter("Titel", "mini Web-Projekt")
	grussSeite.SetzeParameter("Gruss", "Hallo Welt")
	return grussSeite.ErzeugeHTML()
}

func berlinBediener() ([]byte, error) {
	grussSeite.SetzeParameter("Titel", "mini Web-Projekt")
	grussSeite.SetzeParameter("Gruss", "N'juuten")
	return grussSeite.ErzeugeHTML()
}

// erzeugt absichtlich einen Server-Fehler
func stressBediener() ([]byte, error) {
	// ... hier passiert ein Fehler
	return nil, fmt.Errorf("wat willste?")
}

func main() {
	// erzeuge ein Server-Objekt
	var srv http.HttpServer = http.New("127.0.0.1", 8080)

	// Registriere einige Pfade, die Deine WebApp bedienen soll.
	// Hier wird jeder Kombination aus <Http-Methode> und <Url-Pfad> ein Bediener zugeordnet.
	//
	// Wir bedienen die folgende sehr einfache Schnittstelle:

	// Bediene "GET /gruss"
	srv.SetzeHtmlBediener(http.MethodeGet, "/gruss", normalBediener)
	// Bediene "GET /gruss/berlin"
	srv.SetzeHtmlBediener(http.MethodeGet, "/gruss/berlin", berlinBediener)
	// Bediene "GET /stress"
	srv.SetzeHtmlBediener(http.MethodeGet, "/stress", stressBediener)

	// Wir veröffentlichen auch ein Verzeichnis für statische HTML-Seiten,
	// Bilder, Downloads und CSS. Bediene "GET /" (die Wurzel) muss damit
	// in diesem Beispielprojekt nicht extra bedient werden, denn wir haben
	// eine "index.html" in dem Verzeichnis hinterlegt.
	// Diese Datei findet der Browser automatisch.
	srv.VeroeffentlicheVerzeichnis("/", "static")

	// den Web-Server starten - jetzt läuft er und kann über einen Browser oder
	// über andere Client-Programme (z.B. curl oder das eigene ) erreicht werden.
	// Man fährt den Server herunter mit der Tastenkombination Ctrl-C in der Konsole.
	srv.LauscheUndBediene()
}
