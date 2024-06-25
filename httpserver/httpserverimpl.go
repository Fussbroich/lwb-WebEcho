package httpserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

type data struct {
	hostIP string
	portNr uint16
	mux    *http.ServeMux
	logger *log.Logger
}

func New(hostIP string, portnummer uint16) *data {
	s := new(data)
	s.hostIP = hostIP
	s.portNr = portnummer
	// setze logging-Ausgang
	s.logger = log.New(os.Stdout, "** ", 0)
	// der neue (go 1.22) http-Multiplexer macht das Routing leichter
	s.mux = http.NewServeMux()
	return s
}

func (s *data) VeroeffentlicheVerzeichnis(pfad, server_verzeichnis string) {
	s.mux.Handle(pfad, http.FileServer(http.Dir(server_verzeichnis)))
}

func (s *data) Bediene(anfrage_muster string, bediener func() ([]byte, error)) {
	var handler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var err error
			var content []byte
			s.logger.Printf("Server hat %s unter %s erhalten.", r.Method, r.URL.Path)
			//Im Header sagt man dem Browser, was man da anliefert ...
			w.Header().Set("Content-Type", "text/html")
			if content, err = bediener(); err != nil {
				http.Error(w, fmt.Sprintf("interner Serverfehler: %s", err), http.StatusInternalServerError)
				return
			}
			// ...und schreibt die Response -> Fehler sind zu behandeln.
			if _, err := w.Write(content); err != nil {
				http.Error(w, fmt.Sprintf("interner Serverfehler %s", err), http.StatusInternalServerError)
				return
			}
		})
	s.logger.Printf("setze Bediener für %s", anfrage_muster)
	// HINWEIS: das funktioniert so erst mit go 1.22 - sonst kurz umbauen:
	// Vor go 1.22 musste man die HTTP-Methode händisch aus dem Request extrahieren.
	// (fehleranfällig und unschön)
	s.mux.Handle(anfrage_muster, handler)
}

// ListenAndServe ;-)
func (s *data) LauscheUndBediene() error {
	var ctx context.Context
	var cancel context.CancelFunc

	// ein go-Server braucht einen context, der Timeouts speichert und Signale behandelt
	ctx = context.Background() // ein zunächst leerer Context ;-)

	// fange os Interrupt ab
	ctx, cancel = signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	var err error

	// baue Server-Objekt
	server := &http.Server{
		Addr:    net.JoinHostPort(s.hostIP, fmt.Sprint(s.portNr)),
		Handler: http.Handler(s.mux),
	}

	// starte Server nebenläufig
	go func() {
		s.logger.Println("Starte HTTP-Server an der Adresse:")
		s.logger.Printf("http://%s:%d\n", s.hostIP, s.portNr)
		s.logger.Println("Kann mit Ctrl-C gestoppt werden.")
		// lausche unter der Adresse und leite eingehende Anfragen an den Multiplexer weiter
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Printf("Server-Fehler: %s\n", err)
			fmt.Fprintf(os.Stderr, "Server-Fehler: %s\n", err)
		}
	}()

	// warte nebenläufig auf den Context
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done() // Zeit zum gutmütigen Herunterfahren

		s.logger.Println("Stoppe HTTP-Server")
		// ein Timeout; wir warten nicht ewig
		var sdCtx context.Context
		sdCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		s.logger.Println("Warte auf Beenden der Anfragen ...")
		if err := server.Shutdown(sdCtx); err != nil {
			fmt.Fprintf(os.Stderr, "Fehler beim Stoppen des Server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
