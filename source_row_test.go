package dataset_sql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/datatypes"
)

type TestSourceRow map[string][]string

func TestSourceRowGenerator(t *testing.T) {
	// we do this test lots to ensure source row generation is determinstic
	for i := 0; i < 100; i++ {
		if err := testSourceRowGenerator(); err != nil {
			t.Errorf("error on test iteration: %d: %s", i, err.Error())
			return
		}
	}
}

func testSourceRowGenerator() error {
	store, resources, err := makeTestStore()
	if err != nil {
		return fmt.Errorf("error creating test data: %s", err.Error())
	}

	resultf, err := os.Open("testdata/test_source_row_generator.csv")
	if err != nil {
		return fmt.Errorf("error opening data file: %s", err.Error())
	}

	results := []TestSourceRow{}
	if err := json.NewDecoder(resultf).Decode(&results); err != nil {
		return fmt.Errorf("error decoding data json: %s", err.Error())
	}

	srg, err := NewSourceRowGenerator(store, map[string]*dataset.Dataset{
		"t1": resources["t1"],
		"t2": resources["t2"],
	})
	if err != nil {
		return fmt.Errorf("error creating generator: %s", err.Error())
	}

	count := 0
	for srg.Next() {
		sr, err := srg.Row()
		if err != nil {
			return fmt.Errorf("row %d unexpected error: %s", count, err.Error())
		}

		if err := CompareSourceRows(results[count], sr); err != nil {
			return fmt.Errorf("row %d unexpected source row mismatch: %s\nrow:\n%s", count, err.Error(), sr.String())
		}
		count++
	}

	if count != 100 {
		return fmt.Errorf("wrong number of iterations. expected %d, got %d", 100, count)
	}

	return nil
}

func CompareSourceRows(a TestSourceRow, b SourceRow) error {
	if a == nil && b != nil || a != nil && b == nil {
		return fmt.Errorf("nil mismatch: %s != %s", a, b)
	}

	if len(a) != len(b) {
		return fmt.Errorf("sourcerow length mismatch: %d != %d", len(a), len(b))
	}
	for name, rowa := range a {
		rowb := b[name]
		if len(rowa) != len(rowb) {
			return fmt.Errorf("row length mismatch: %d != %d", len(rowa), len(rowb))
		}
		for i, cell := range rowa {
			if !bytes.Equal([]byte(cell), rowb[i]) {
				return fmt.Errorf("byte mismatch on source %s, column: %d. expected: %s, got: %s", name, i, string(cell), string(rowb[i]))

			}
		}
	}
	return nil
}

func TestSourceRowFilter(t *testing.T) {
	// TODO - need to test limit / offset clauses
	stmt, err := Parse("select * from t1 where t1.a > 5")
	if err != nil {
		t.Errorf("statement parse error: %s", err.Error())
		return
	}

	srg, err := NewSourceRowFilter(stmt, nil)
	if err != nil {
		t.Errorf("errog creating source row filter: %s", err.Error())
		return
	}

	resources := map[string]*dataset.Dataset{
		"t1": {
			Structure: &dataset.Structure{
				Format: dataset.CSVDataFormat,
				Schema: &dataset.Schema{
					Fields: []*dataset.Field{
						{Name: "a", Type: datatypes.Integer},
					},
				},
			},
		},
	}

	if err := PrepareStatement(stmt, resources); err != nil {
		t.Errorf("error preparing statement: %s", err.Error())
		return
	}
	cols := CollectColNames(stmt)

	cases := []struct {
		sr     SourceRow
		expect bool
	}{
		{SourceRow{"t1": [][]byte{[]byte("0")}}, false},
		{SourceRow{"t1": [][]byte{[]byte("5")}}, false},
		{SourceRow{"t1": [][]byte{[]byte("6")}}, true},
		{SourceRow{"t1": [][]byte{[]byte("10")}}, true},
		{SourceRow{"t1": [][]byte{[]byte("200")}}, true},
		{SourceRow{"t1": [][]byte{[]byte("7")}}, true},
	}

	for i, c := range cases {
		if err := SetSourceRow(cols, c.sr); err != nil {
			t.Errorf("case %d error setting source row: %s", i, err.Error())
			return
		}

		got := srg.Match()
		if got != c.expect {
			t.Errorf("case %d fail %t != %t", i, c.expect, got)
		}
	}
}
