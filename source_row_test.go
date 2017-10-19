package dataset_sql

import (
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
		_, err := srg.Row()
		if err != nil {
			t.Errorf("row %d unexpected error: %s", count, err.Error())
			return
		}

		// i := srg.Indexes()
		// fmt.Println(i["t1"], i["t2"])
	}

	if count != 100 {
		t.Errorf("wrong number of iterations. expected %d, got %d", 100, count)
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
