package httpserver

//Vor.: hostIP ist die Server-IP-Adresse in der Form
//      "192.168.1.100", portnummer ist die Nummer des Ports, über den
//      der Server kontaktiert werden soll.
//Erg.: Der Server ist initialisiert
//Eff.: In der Konsole des Server-Prozesses werden fortlaufend
//      Status-Nachrichten ausgegeben.
// func New(hostIP string, portnummer uint16)

type HttpServer interface {

	// Http-Anfragen bedienen:
	//	anfrage ist "<Http-Methode> <Url-Muster>"
	//	bediener ist die "Bearbeitungsfunktion" des Servers. Sie generiert und
	//	liefert die Antwort in Form einer Bytefolge (Html). Falls dabei ein
	// Fehler passiert, liefert sie diesen.
	SetzeBediener(http_methode, url_muster string, bediener func() ([]byte, error))

	// Das komplette Server-Verzeichnis ist unter dem url zugreifbar.
	VeroeffentlicheVerzeichnis(url string, server_verzeichnis string)

	// Die Start-Methode des Web-Servers wird zuletzt aufgerufen.
	// Anweisungen im Quelltext hinter dem Aufruf der Methode werden
	// erst ausgeführt, nachdem der Server beendet wurde.
	LauscheUndBediene() error
}
