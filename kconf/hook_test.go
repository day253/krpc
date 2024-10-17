package kconf

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	os.Setenv("FOO", "bar")
	tpl := `{{env "FOO"}}`
	if err := runt(tpl, "bar"); err != nil {
		t.Error(err)
	}
}

func TestExpandEnv(t *testing.T) {
	os.Setenv("FOO", "bar")
	tpl := `{{expandenv "Hello $FOO"}}`
	if err := runt(tpl, "Hello bar"); err != nil {
		t.Error(err)
	}
}

func TestExpandConfig(t *testing.T) {
	os.Setenv("ENV_ROLE", "test")
	tpl := `Address: ":8181"
ServiceName: '/{{env "ENV_ROLE"}}/product-interface/re-voice'`
	if err := runt(tpl, `Address: ":8181"
ServiceName: '/test/product-interface/re-voice'`); err != nil {
		t.Error(err)
	}
}

func TestServiceNameWithEnvRole(t *testing.T) {
	os.Setenv("ENV_ROLE", "test")
	tpls := []string{
		`{{expandenv "/$ENV_ROLE/product-interface/re-voice"}}`,
		`/{{env "ENV_ROLE"}}/product-interface/re-voice`,
	}
	for _, tpl := range tpls {
		if err := runt(tpl, "/test/product-interface/re-voice"); err != nil {
			t.Error(err)
		}
	}
}
