package httpclients

// Vor.: Der Server läuft unter der angegebenen Adresse mit dem angegebenen Verbindungsport.
//	zielIP ist die Server-IP-Adresse in der Form "192.168.1.100",
//	portnummer ist die Nummer des Ports, über den der Server kontaktiert werden soll.
// Erg.: Der Client ist initialisiert
// Eff.: In der Konsole des Client-Prozesses werden Status-Nachrichten ausgegeben.
// func New(zielIP string, portnummer uint16)

type HttpClient interface {
	// Sendet eine Anfrage an den Server
	//	Vor.: Der Server läuft unter der angegebenen Adresse mit dem angegebenen Verbindungsport.
	//	Erg.: Der Antwort-Body vom Server ist gemäß dem Server-Protokoll geliefert. Bei einem
	// Server-Fehler wird dieser statt dessen geliefert.
	Anfragen(http_methode, url string) ([]byte, error)
}
