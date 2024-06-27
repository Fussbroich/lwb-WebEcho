package htmlvorlagen

import (
	"bytes"
	"html/template"
)

type data struct {
	vorlage *template.Template
	params  map[string]any
}

func NewHtmlVorlage(html_text string) *data {
	v := new(data)
	v.vorlage = template.Must(template.New("tpl").Parse(html_text))
	v.params = make(map[string]any)
	return v
}

func (v *data) SetzeParameter(name string, wert any) {
	v.params[name] = wert
}

func (v *data) ErzeugeHTML() ([]byte, error) {
	var body = &bytes.Buffer{}
	if err := v.vorlage.Execute(body, v.params); err != nil {
		return nil, err
	}
	return body.Bytes(), nil
}
