package kconf

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/mitchellh/mapstructure"
)

const (
	defaultTemplateName = "config"
)

func StringRenderTextTemplateHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t.Kind() != reflect.String {
			return data, nil
		}

		return runRaw(data.(string), map[string]string{})
	}
}

// runt runs a template and checks that the output exactly matches the expected string.
func runt(tpl, expect string) error {
	return runtv(tpl, expect, map[string]string{})
}

// runtv takes a template, and expected return, and values for substitution.
//
// It runs the template and verifies that the output is an exact match.
func runtv(tpl, expect string, vars interface{}) error {
	t := template.Must(template.New(defaultTemplateName).Funcs(sprig.TxtFuncMap()).Parse(tpl))
	var b bytes.Buffer
	err := t.Execute(&b, vars)
	if err != nil {
		return err
	}
	if expect != b.String() {
		return fmt.Errorf("expected '%s', got '%s'", expect, b.String())
	}
	return nil
}

// runRaw runs a template with the given variables and returns the result.
func runRaw(tpl string, vars interface{}) (string, error) {
	t := template.Must(template.New(defaultTemplateName).Funcs(sprig.TxtFuncMap()).Parse(tpl))
	var b bytes.Buffer
	err := t.Execute(&b, vars)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
