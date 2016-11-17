package dataset_sql

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"testing"

	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/dataset_generate"
	"github.com/qri-io/datatype"
)

type TestDomain struct {
	datasets []*dataset.Dataset
}

func (t *TestDomain) DatasetForAddress(adr dataset.Address) (*dataset.Dataset, error) {
	for _, ds := range t.datasets {
		if ds.Address.Equal(adr) {
			return ds, nil
		}
	}

	return nil, fmt.Errorf("Not Found")
}

func TestSelectFields(t *testing.T) {
	created := &dataset.Field{Name: "created", Type: datatype.Date}
	title := &dataset.Field{Name: "title", Type: datatype.String}
	views := &dataset.Field{Name: "views", Type: datatype.Integer}
	rating := &dataset.Field{Name: "rating", Type: datatype.Float}
	notes := &dataset.Field{Name: "notes", Type: datatype.String}

	ds := dataset_generate.RandomDataset(func(o *dataset_generate.RandomDatasetOpts) {
		o.Name = "select_test"
		o.Address = dataset.NewAddress("select_test")
		o.Fields = []*dataset.Field{created, title, views, rating, notes}
		o.NumRecords = 10
	})

	dsTwo := dataset_generate.RandomDataset(func(o *dataset_generate.RandomDatasetOpts) {
		o.Name = "select_test_two"
		o.Address = dataset.NewAddress("select_test_two")
		o.Fields = []*dataset.Field{created, title, views, rating, notes}
		o.NumRecords = 10
	})

	domain := &TestDomain{datasets: []*dataset.Dataset{ds, dsTwo}}

	cases := []struct {
		statement string
		expect    error
		fields    []*dataset.Field
		numRows   int
	}{
		// {"select * from select_test", nil, []*dataset.Field{created, title, views, rating, notes}, 10},
		// {"select created, title, views, rating, notes from select_test", nil, []*dataset.Field{created, title, views, rating, notes}, 10},
		// {"select select_test->created from select_test limit 5", nil, []*dataset.Field{created}, 5},
		// {"select created from select_test limit 1 offset 1", nil, []*dataset.Field{created}, 1},
		{"select * from select_test, select_test_two", nil, []*dataset.Field{created, title, views, rating, notes, created, title, views, rating, notes}, 20},
		// {"select 1 from select_test", nil, []*dataset.Field{&dataset.Field{Name: "result", Type: datatype.Integer}}, 1},
	}

	for i, c := range cases {
		stmt, err := Parse(c.statement)
		if err != nil {
			t.Errorf("case %d parse error: %s", i, err.Error())
			continue
		}

		results, err := stmt.Exec(domain)
		if err != c.expect {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.expect, err.Error())
			continue
		}

		if len(results.Fields) != len(c.fields) {
			t.Errorf("case %d field length mismatch. expected: %d, got: %d", i, len(c.fields), len(results.Fields))
			continue
		}

		for j, f := range c.fields {
			if results.Fields[j].Name != f.Name {
				t.Errorf("case %d field %d name mismatch. expected: %s, got: %s", i, j, f.Name, results.Fields[j].Name)
				continue
			}
			if results.Fields[j].Type != f.Type {
				t.Errorf("case %d field %d type mismatch. expected: %s, got: %s", i, j, f.Type, results.Fields[j].Type)
				continue
			}
		}

		r := csv.NewReader(bytes.NewBuffer(results.Data))
		records, err := r.ReadAll()
		if err != nil {
			t.Error(err.Error())
			continue
		}

		if len(records) != c.numRows {
			t.Errorf("case %d result count mismatch. expected: %d, got: %d", i, c.numRows, len(records))
			continue
		}

		// table := tablewriter.NewWriter(os.Stdout)
		// table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		// table.SetCenterSeparator("|")
		// table.SetHeader(results.FieldNames())
		// table.AppendBulk(records)
		// table.Render()
	}
}
