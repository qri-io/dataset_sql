package sqlparser

import (
	"testing"

	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/dataset_generate"
	"github.com/qri-io/datatype"
)

func TestSelectFields(t *testing.T) {
	created := &dataset.Field{Name: "created", Type: datatype.Date}
	title := &dataset.Field{Name: "title", Type: datatype.String}
	views := &dataset.Field{Name: "views", Type: datatype.Integer}
	rating := &dataset.Field{Name: "rating", Type: datatype.Float}
	notes := &dataset.Field{Name: "notes", Type: datatype.String}

	ds, err := dataset_generate.RandomDataset(func(o *dataset_generate.RandomDatasetOpts) {
		o.Name = "select_test"
		o.Fields = []*dataset.Field{created, title, views, rating, notes}
		o.NumRecords = 10
	})
	if err != nil {
		t.Error(err.Error())
		return
	}

	cases := []struct {
		statement string
		expect    error
		fields    []*dataset.Field
		numRows   int
	}{
		{"select * from select_test", nil, []*dataset.Field{created, title, views, rating, notes}, 10},
		{"select created from select_test", nil, []*dataset.Field{created}, 10},
	}

	for i, c := range cases {
		results, err := repo.Select(c.statement)
		if err != c.expect {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.expect, err.Error())
			continue
		}

		if len(results.Fields) != len(c.fields) {
			t.Errorf("case %d field length mismatch. expected: %d, got: %d", i, len(c.fields), len(results.Fields))
			continue
		}

		for i, f := range c.fields {
			if results.Fields[i].Name != f.Name {
				t.Errorf("case %d field name mismatch. expected: %s, got: %s", i, f.Name, results.Fields[i].Name)
				continue
			}
			if results.Fields[i].Type != f.Type {
				t.Errorf("case %d field type mismatch. expected: %s, got: %s", i, f.Type, results.Fields[i].Type)
				continue
			}
		}
	}

	// r := csv.NewReader(bytes.NewBuffer(result.Data))
	// records, err := r.ReadAll()
	// if err != nil {
	// 	t.Error(err.Error())
	//
	// fmt.Println(records)	return
	// }
}
