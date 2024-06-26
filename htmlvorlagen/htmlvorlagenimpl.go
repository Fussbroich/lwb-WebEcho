package htmlvorlagen

import (
	"bytes"
	"html/template"
)

type data struct {
	vorlage *template.Template
	params  map[string]string
}

func NewVorlage(text string) (*data, error) {
	var err error
	v := new(data)
	if v.vorlage, err = template.New("tpl").Parse(text); err != nil {
		return nil, err
	}
	v.params = make(map[string]string)
	return v, nil
}

func (v *data) SetzeParameter(name, wert string) {
	v.params[name] = wert
}

func (v *data) ErzeugeHTML() ([]byte, error) {
	var body = &bytes.Buffer{}
	if err := v.vorlage.Execute(body, v.params); err != nil {
		return nil, err
	}
	return body.Bytes(), nil
}
