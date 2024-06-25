package httpserver

type HttpServer interface {

	// Http-Methoden und URL-Pfade bedienen
	BedieneGET(pfad string, bediener func() []byte)

	BedienePOST(pfad string, bediener func() []byte)

	BedienePUT(pfad string, bediener func() []byte)

	BedieneDELETE(pfad string, bediener func() []byte)

	BedieneVerzeichnis(pfad string, server_verzeichnis string)

	// die Start-Methode des Web-Servers
	LauscheUndBediene() error
}
