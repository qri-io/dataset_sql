package dataset_sql

import (
	"testing"
)

func TestSourceRowGenerator(t *testing.T) {
	// cases := []struct {
	// }{}

	// for i, c := range cases {

	// }
}

func TestSourceRowFilter(t *testing.T) {
	ast, err := Parse("select")
	if err != nil {
		t.Errorf("statement parse error: %s", err.Error())
		return
	}

	srg, err := NewSourceRowFilter(ast)
	if err != nil {
		t.Errorf("errog creating source row filter: %s", err.Error())
		return
	}

	cases := []struct {
		row    SourceRow
		expect bool
	}{
		{SourceRow{}, false},
	}

	for i, c := range cases {
		got := srg.Filter(c.row)
		if got != c.expect {
			t.Errorf("case %d fail %t != %t", i, c.expect, got)
		}
	}
}
