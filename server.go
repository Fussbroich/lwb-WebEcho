// Zweck: Wie baue ich einen Webserver in go?
// Mögliche Anwendung für dynamische Views auf eine Datenbank über einen Browser.
//	Go ab 1.22 (!), aktueller Browser
//
//	verwendete Pakete: strconv, embed, fmt, io, log, net, os, os/signal, time, sowie:
//
//		HTTP:		https://pkg.go.dev/net/http
//		HTML:		https://pkg.go.dev/html/template
//
//	Datum: 23.05.2024

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func echoHandler(tag string, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Ich handle %s.\n"+
			"Du hast den Server mit der %s Methode unter %s erreicht.",
			tag, r.Method, r.URL.Path)
		logger.Println(msg) // -> logging für die Server-Konsole

		// response (eine Web-Page) für den Browser/Client -> Fehler sind zu behandeln.
		// Tipp: Benutze html/template zur dynamischen Erzeugung der HTML-Seiten
		htmlContent := `
<html>
	<head>
    	<title>WebEcho</title>
	</head>
	<body>
    	<div>` + msg + `</div>
	</body>
</html>`

		w.Header().Set("Content-Type", "text/html")
		if _, err := w.Write([]byte(htmlContent)); err != nil {
			http.Error(w, fmt.Sprintf("interner Serverfehler %s", err), http.StatusInternalServerError)
			return
		}
	})
}

func main() {
	// ein neuer Multiplexer für das Routing
	mux := http.NewServeMux()
	logger := log.New(os.Stdout, "**", 0)

	// registriere einige Methoden und Routen, die Deine WebApp beantworten soll
	// Hinweis: vor go 1.22 musste man die Methode händisch aus dem Request extrahieren.
	// Der neue ServeMux seit 1.22 nimmt einem diese und viele andere Arbeiten ab ...
	mux.Handle("GET /meinPfad", echoHandler("Get an /meinPfad", logger))
	mux.Handle("GET /{$}", echoHandler("Get an der Wurzel", logger))
	mux.Handle("POST /meinPfad", echoHandler("Posts an /meinPfad", logger))
	mux.Handle("/", echoHandler("Fallbacks", logger))

	// Baue einen Web-Server und starte ihn
	server := &http.Server{
		// Hier würde man die eigene IP-Adresse angeben und die Firewall für ankommendes TCP freigeben
		// Damit wird der Server von anderen erreichbar.
		Addr:    net.JoinHostPort("localhost", "8081"),
		Handler: http.Handler(mux),
	}

	fmt.Printf("starte Web-Server unter http://%s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Server-Fehler: %s\n", err)
	}
}
