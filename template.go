package main

import (
	"encoding/json"
	"strings"
	"text/template"
)

type TemplateType struct {
	Name    string
	Comment string
}

type Template struct {
	*template.Template
	Config TemplateConfig
}

type TemplateConfig struct {
	FlagName string   `json:"flag"`
	Usage    string   `json:"usage"`
	Imports  []string `json:"imports"`
}

func NewTemplate(name, input string) (Template, error) {
	cfg := TemplateConfig{}
	if endOfJson, has := findJsonObjectEnd(input); has {
		offset := endOfJson + 1
		reader := strings.NewReader(input[:offset])
		if err := json.NewDecoder(reader).Decode(&cfg); err != nil {
			return Template{}, err
		}
		input = input[offset:]
		// input = strings.TrimLeftFunc(input[offset:], unicode.IsSpace)
	}

	tmpl, err := template.New(name).Parse(input)
	if err != nil {
		return Template{}, err
	}

	return Template{
		Template: tmpl,
		Config:   cfg,
	}, nil
}

func findJsonObjectEnd(input string) (int, bool) {
	depth := 0
	for i, c := range input {
		switch c {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return i, true
			}
		}
	}
	return 0, false
}
