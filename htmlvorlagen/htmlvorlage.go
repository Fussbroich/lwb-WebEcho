package htmlvorlagen

// Zweck: Erzeuge dynamische Html-Seiten.
//
// Autor: T. Schrader
//
// Datum: 02.06.2024
//
// Eine HtmlVorlage wird aus einem HTML-Text erzeugt. Der HTML-Text muss bereits syntaktisch
// richtiges HTML sein, aber er kann zusätzlich Parameter enthalten, die erst zur Laufzeit
// mit Daten gefüllt werden. Das ist die Aufgabe dieser Vorlage.
// Ein Parameter wird im HTML-Text durch {{.name}} eingefügt.
//
// Zur Laufzeit muss der HtmlVorlage dann mit der Methode SetzeParameter ein Wert für jeden
// verwendeten Parameter übergeben werden. Der Wert kann irgendein Objekt sein. Daher wird hier
// ein Wert vom Typ "any" erwartet.
// Konstruktor: NewVorlage(html_text string) *data
type HtmlVorlage interface {

	// Für jeden Parameter, der in der Vorlage mit {{.name}} verwendet wird,
	// muss hier ein Wert gesetzt werden.
	SetzeParameter(name string, wert any)

	// Erst nachdem alle Parameter der Vorlage einen Wert haben,
	// wird das Html für den Browser erzeugt. Falls ein Wert weggelassen wird,
	// wird hierfür im erzeugten HTML nichts eingesetzt.
	ErzeugeHTML() ([]byte, error)
}
