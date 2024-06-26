// Thema: Ich baue einen Webserver in go
//
// Zweck: Erzeuge und veröffentliche dynamische Web-Seiten mit Deinen Daten aus einem go-Programm.
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
const grussSeite = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<!-- Trick: nach Änderungen zähle die Versionsnummer v=... beliebig hoch -->
	    <link rel="stylesheet" href="/style.css?v=1.2">
		<title>{{$.Title}}</title>
	</head>
	<body>
	<!-- ein div für die Navigation -->
	<div class="navigation">
		<a href="/gruss">grüße normal</a> |
		<a href="/gruss/berlin">grüße berlinerisch</a> |
		<a href="/stress">provoziere Stress</a> |
		<a href="/">zur Übersicht</a>
	</div>
	<!-- ein div für den Inhalt -->
	<div class="inhalt">{{$.Gruss}}</div>
</body>
</html>`

// Drei beispielhafte Anfragen-Bediener, die den "Body"
// für eine Http-Response erzeugen: Hier wird HTML produziert.
// Wir benutzen eine HtmlVorlage zur Erzeugung von HTML-Seiten und
// füllen die Parameter mit Daten.
// Du kannst das Html auch ohne Vorlage selbst erzeugen.
// Der Server liefert das vom Bediener erzeugte Html an den Browser.
func normalBediener() ([]byte, error) {
	var vorlage html.HtmlVorlage
	var err error
	vorlage, err = html.NewVorlage(grussSeite)
	if err != nil {
		return nil, err
	}
	// Fülle mit Daten ("Parameter")
	vorlage.SetzeParameter("Title", "mini Web-Projekt")
	vorlage.SetzeParameter("Gruss", "Hallo Welt")
	return vorlage.ErzeugeHTML()
}

func berlinBediener() ([]byte, error) {
	var vorlage html.HtmlVorlage
	var err error
	vorlage, err = html.NewVorlage(grussSeite)
	if err != nil {
		return nil, err
	}
	vorlage.SetzeParameter("Title", "mini Web-Projekt")
	vorlage.SetzeParameter("Gruss", "N'juuten")
	return vorlage.ErzeugeHTML()
}

// erzeugt absichtlich einen Fehler
func stressBediener() ([]byte, error) {
	// ... hier passiert ein Fehler
	return nil, fmt.Errorf("wat willste?")
}

func main() {
	// erzeuge ein Server-Objekt
	var srv http.HttpServer = http.New("127.0.0.1", 8080)

	// Registriere einige Pfade, die Deine WebApp bedienen soll.
	// Hier wird jeder Kombination aus <Http-Methode> und <Url-Pfad> ein Bediener zugeordnet.

	// HINWEIS: das funktioniert so erst mit go 1.22:
	// siehe https://eli.thegreenplace.net/2023/better-http-server-routing-in-go-122

	// Wir bedienen die folgende sehr einfache Schnittstelle.

	// Bediene "GET /gruss"
	srv.SetzeBediener(http.MethodeGet, "/gruss", normalBediener)
	// Bediene "GET /gruss/berlin"
	srv.SetzeBediener(http.MethodeGet, "/gruss/berlin", berlinBediener)
	// Bediene "GET /stress"
	srv.SetzeBediener(http.MethodeGet, "/stress", stressBediener)

	// Wir veröffentlichen auch ein Verzeichnis für statische HTML-Seiten,
	// Bilder, Downloads und CSS. Bediene "GET /" (die Wurzel) muss damit
	// nicht extra bedient werden, denn hier liegt eine "index.html" -
	// diese Datei findet der Browser automatisch.
	srv.VeroeffentlicheVerzeichnis("/", "static")

	// den Web-Server starten - jetzt läuft er und kann über einen Browser oder
	// über das Programm curl erreicht werden.
	// Man fährt den Server herunter mit der Tastenkombination Ctrl-C in der Konsole.
	srv.LauscheUndBediene()
}
