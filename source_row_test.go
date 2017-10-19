package dataset_sql

import (
	"fmt"
	"testing"
)

func TestSourceRowGenerator(t *testing.T) {
	store, resources, err := makeTestData()
	if err != nil {
		t.Errorf("error creating test data: %s", err.Error())
		return
	}

	srg, err := NewSourceRowGenerator(store, resources)
	if err != nil {
		t.Errorf("error creating generator: %s", err.Error())
		return
	}

	count := 0
	for srg.Next() {
		count++
		// TODO - check that rows are iterating the right values
		row, err := srg.Row()
		if err != nil {
			t.Errorf("row %d unexpected error: %s", count, err.Error())
			return
		}

		fmt.Println(len(row["t1"]), len(row["t2"]))
	}

	if count != 20 {
		t.Errorf("wrong number of iterations. expected %d, got %d", 20, count)
	}
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
