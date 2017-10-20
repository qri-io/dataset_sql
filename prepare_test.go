package dataset_sql

import (
	"testing"

	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/datatypes"
)

func TestPrepareStatement(t *testing.T) {
	stmt, err := Parse("select * from t1")
	if err != nil {
		t.Errorf("error parsing statement: %s", err.Error())
		return
	}

	resources := map[string]*dataset.Structure{
		"t1": &dataset.Structure{
			Format: dataset.CsvDataFormat,
			Schema: &dataset.Schema{
				Fields: []*dataset.Field{
					&dataset.Field{Name: "a", Type: datatypes.Integer},
					&dataset.Field{Name: "b", Type: datatypes.Any},
				},
			},
		},
	}

	if err := PrepareStatement(stmt, resources); err != nil {
		t.Errorf("error remapping statement: %s", err.Error())
		return
	}

	sel := stmt.(*Select)
	if len(resources["t1"].Schema.Fields) != len(sel.SelectExprs) {
		t.Errorf("select expressions length mismatch. expected %d, got %d", len(resources["t1"].Schema.Fields), len(sel.SelectExprs))
	}
}
