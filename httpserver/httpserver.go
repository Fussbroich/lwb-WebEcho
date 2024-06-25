package httpserver

type HttpServer interface {

	// Http-Anfragen bedienen
	Bediene(anfrage_muster string, bediener func() ([]byte, error))

	VeroeffentlicheVerzeichnis(pfad string, server_verzeichnis string)

	// die Start-Methode des Web-Servers
	LauscheUndBediene() error
}
