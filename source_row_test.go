package dataset_sql

import (
	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/datatypes"
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
	}

	if count != 100 {
		t.Errorf("wrong number of iterations. expected %d, got %d", 100, count)
	}
}

func TestSourceRowFilter(t *testing.T) {
	stmt, err := Parse("select * from t1 where t1.a > 5")
	if err != nil {
		t.Errorf("statement parse error: %s", err.Error())
		return
	}

	srg, err := NewSourceRowFilter(stmt)
	if err != nil {
		t.Errorf("errog creating source row filter: %s", err.Error())
		return
	}

	cases := []struct {
		sr     SourceRow
		expect bool
	}{
		{SourceRow{"t1": [][]byte{[]byte("0")}}, false},
		{SourceRow{"t1": [][]byte{[]byte("5")}}, false},
		{SourceRow{"t1": [][]byte{[]byte("6")}}, true},
	}

	st := &dataset.Structure{
		Format: dataset.CsvDataFormat,
		Schema: &dataset.Schema{
			Fields: []*dataset.Field{
				&dataset.Field{Name: "a", Type: datatypes.Integer},
			},
		},
	}

	resources := map[string]*dataset.Structure{
		"t1": st,
	}

	if err := PopulateTableInfo(stmt, resources); err != nil {
		t.Errorf("error populating table info: %s", err.Error())
		return
	}
	cols := CollectColNames(stmt)

	for i, c := range cases {
		if err := SetSourceRow(cols, c.sr); err != nil {
			t.Errorf("case %d error setting source row: %s", i, err.Error())
			return
		}

		got := srg.Filter()
		if got != c.expect {
			t.Errorf("case %d fail %t != %t", i, c.expect, got)
		}
	}
}
