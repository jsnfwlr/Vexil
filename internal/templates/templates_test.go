package templates_test

import (
	"testing"

	"github.com/jsnfwlr/vexil/internal/templates"
)

func TestLoadFiles(t *testing.T) {
	table, err := templates.Files.ReadFile("table.html")
	if err != nil {
		t.Error(err)
	}

	t.Log(string(table))
}
