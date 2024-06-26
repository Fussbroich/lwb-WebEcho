package htmlvorlagen

// Zweck: Erzeuge dynamische Html-Seiten.
//
// Autor: T. Schrader
//
// Datum: 02.06.2024
//
// Konstruktor: NewVorlage(text string) (*data, error)
// Schon beim Erzeugen einer Vorlage kann es einen Fehler
// geben, der behandelt werden muss. In dem Fall wird auch
// keine Vorlage erzeugt.
type HtmlVorlage interface {

	// Für jeden Parameter, der in der Vorlage verwendet wird,
	// muss hier ein Wert gesetzt werden.
	SetzeParameter(name, wert string)

	// Erst nachdem alle Parameter der Vorlage einen Wert haben,
	// wird das Html für den Browser erzeugt.
	ErzeugeHTML() ([]byte, error)
}
