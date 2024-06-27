package httpclients

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
)

type data struct {
	hostIP string
	portNr uint16
}

func NewHttpClient(zielIP string, portnummer uint16) *data {
	s := new(data)
	s.hostIP = zielIP
	s.portNr = portnummer
	return s
}

func (s *data) Anfragen(http_methode, zielUrl string) ([]byte, error) {
	var anf string
	var ant *http.Response
	var body []byte
	var pfad string
	var fehler error

	pfad, fehler = url.JoinPath(zielUrl)
	if fehler != nil {
		return []byte{}, fmt.Errorf("kann Anfrage nicht erzeugen: %v", fehler)
	}
	anf = fmt.Sprintf("http://%s/%s", net.JoinHostPort(s.hostIP, fmt.Sprint(s.portNr)), pfad)

	switch http_methode {
	case "GET":
		ant, fehler = http.Get(anf)
	default:
		return []byte{}, fmt.Errorf("%s wird derzeit nicht unterstützt", http_methode)
	}

	if fehler != nil {
		return []byte{}, fmt.Errorf("kann Anfrage nicht senden: %v", fehler)
	}
	defer ant.Body.Close()

	body, fehler = io.ReadAll(ant.Body)
	if fehler != nil {
		return []byte{}, fmt.Errorf("kann Antwort nicht lesen: %v", fehler)
	}

	if ant.StatusCode != 200 {
		return body, fmt.Errorf("die Statusmeldung ist %s", ant.Status)
	}
	return body, nil
}
