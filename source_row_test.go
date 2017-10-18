package dataset_sql

import (
	"testing"
)

func TestSourceRowGenerator(t *testing.T) {
	cases := []struct {
	}{}

	for i, c := range cases {

	}
}

func TestSourceRowFilter(t *testing.T) {
	ast, err := Parse("select")
	if err != nil {
		t.Errorf("statement parse error: %s", err.Error())
		return
	}

	srg := NewSourceRowFilter(ast)
	cases := []struct {
		row    SourceRow
		expect bool
	}{
		{SourceRow{}, false},
	}

	for i, c := range cases {
		got := srg.Filter(c.row)
		if got != c.expect {
			t.Errorf("case")
		}
	}
}
