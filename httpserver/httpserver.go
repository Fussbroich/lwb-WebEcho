package httpserver

// Vor.: mindestens go 1.22
//	siehe https://eli.thegreenplace.net/2023/better-http-server-routing-in-go-122
// Vor.:
//	hostIP ist die Server-IP-Adresse in der Form "192.168.1.100",
//	portnummer ist die Nummer des Ports, über den der Server kontaktiert werden soll.
// Erg.: Der Server ist initialisiert
// Eff.: In der Konsole des Server-Prozesses werden fortlaufend
//	 Status-Nachrichten ausgegeben.
// func New(hostIP string, portnummer uint16)
type HttpServer interface {

	// Http-Anfragen bedienen:
	// Eine Http-Anfrage lautet "<Http-Methode> <Url>".
	// Ein bediener ist die "Bearbeitungsfunktion" des Servers. Sie generiert und
	// liefert die Antwort in Form einer Bytefolge (Html). Falls dabei ein
	// Fehler passiert, liefert sie diesen zurück.
	//
	// Erg.:
	//	Registriert den Bediener auf diese Methode und alle URLs, die zu dem Muster passen.
	//	Http-Anfragen, die diesem Muster entsprechen, werden fortan durch diesen Bediener
	//	beantwortet.
	//	Das funktioniert so: Der Bediener produziert eine Ausgabe, und diese Ausgabe
	//	wird als Body der Http-Antwort gesendet. Wenn bei der Ausführung des Bedieners
	//	ein Fehler passiert, so verursacht das es einen sog. "internen Server-Fehler",
	//	der im Browser angezeigt wird.
	SetzeHtmlBediener(http_methode, url_muster string, bediener func() ([]byte, error))

	// Vorsicht: Das komplette Server-Verzeichnis ist unter dem url zugreifbar.
	VeroeffentlicheVerzeichnis(url string, server_verzeichnis string)

	// Die Start-Methode des Web-Servers wird zuletzt aufgerufen.
	// Anweisungen im Quelltext hinter dem Aufruf der Methode werden
	// erst ausgeführt, nachdem der Server beendet wurde.
	LauscheUndBediene() error
}
