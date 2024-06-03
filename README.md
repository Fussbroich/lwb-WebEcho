 Thema: Wie baue ich einen Webserver in go?

Zweck: Mögliche Anwendung für dynamische Views auf eine Datenbank über einen Browser.
Go ab 1.22 (!), aktueller Browser (es geht auch mit go 1.18, ist aber etwas mehr Arbeit)

verwendete Pakete: strconv, embed, fmt, io, log, net, os, os/signal, time, sowie:

* HTTP:		https://pkg.go.dev/net/http
* HTML:		https://pkg.go.dev/html/template

Datum: 02.06.2024
