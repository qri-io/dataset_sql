package dataset_sql

import (
	"testing"

	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/datatypes"
	// "github.com/qri-io/dataset/dsfs"
)

// func TestPrepare(t *testing.T) {
// 	ds := &dataset.Transform{
// 		Data: "select * from t1",
// 		Resources: map[string]*dataset.Dataset{
// 			"t1": {
// 				Data: "t1/data/path",
// 				Structure: &dataset.Structure{
// 					Format: dataset.CSVDataFormat,
// 					Schema: &dataset.Schema{
// 						Fields: []*dataset.Field{
// 							{Name: "one", Type: datatypes.String},
// 							{Name: "two", Type: datatypes.Boolean},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	prep, err := Prepare(ds, &ExecOpt{Format: dataset.CSVDataFormat})
// 	if err != nil {
// 		t.Errorf("unexpected error from Prepare: %s", err.Error())
// 		return
// 	}

// 	str := String(prep.stmt)
// 	expect := "select t1.a as a, t1.b as b from t1"
// 	if expect != str {
// 		t.Errorf("statement error, expected: '%s', got: '%s'", expect, str)
// 		return
// 	}

// 	if prep.paths["t1"].String() != "/t1/data/path" {
// 		t.Errorf("data path error, expected %s, got %s", "/t1/data/path", prep.paths["a"].String())
// 	}
// }

// func TestPreparedQueryPath(t *testing.T) {
// 	store, ds, err := makeTestStore()
// 	if err != nil {
// 		t.Errorf("error making test store: %s", err.Error())
// 		return
// 	}

// 	opts := &ExecOpt{
// 		Format: dataset.CSVDataFormat,
// 	}

// 	q := &dataset.Transform{
// 		Syntax: "sql",
// 		Data:   "select * from t1",
// 		Resources: map[string]*dataset.Dataset{
// 			"t1": ds["t1"],
// 		},
// 	}

// 	path, err := PreparedQueryPath(store, q, opts)
// 	if err != nil {
// 		t.Errorf("error preparing query path: %s", err.Error())
// 		return
// 	}

// 	_, _, err = Exec(store, q, func(o *ExecOpt) { o.Format = dataset.CSVDataFormat })
// 	if err != nil {
// 		t.Errorf("error executing query: %s", err.Error())
// 		return
// 	}

// 	r := &dataset.Dataset{}
// 	r.Transform = q
// 	rpath, err := dsfs.SaveDataset(store, r, false)
// 	if err != nil {
// 		t.Errorf("error saving dataset: %s", err.Error())
// 		return
// 	}

// 	r2, err := dsfs.LoadDatasetRefs(store, rpath)
// 	if err != nil {
// 		t.Errorf("error loading saved dataset: %s", err.Error())
// 		return
// 	}

// 	if !r2.Transform.Path().Equal(path) {
// 		t.Errorf("path mistmatch. expected: %s, got: %s", r2.Transform.Path().String(), path.String())
// 		return
// 	}
// }

func TestPrepareStatement(t *testing.T) {
	stmt, err := Parse("select * from t1")
	if err != nil {
		t.Errorf("error parsing statement: %s", err.Error())
		return
	}

	resources := map[string]*dataset.Dataset{
		"t1": {
			Structure: &dataset.Structure{
				Format: dataset.CSVDataFormat,
				Schema: &dataset.Schema{
					Fields: []*dataset.Field{
						{Name: "a", Type: datatypes.Integer},
						{Name: "b", Type: datatypes.Float},
					},
				},
			},
		},
	}

	if err := PrepareStatement(stmt, resources); err != nil {
		t.Errorf("error remapping statement: %s", err.Error())
		return
	}

	sel := stmt.(*Select)
	if len(resources["t1"].Structure.Schema.Fields) != len(sel.SelectExprs) {
		t.Errorf("select expressions length mismatch. expected %d, got %d", len(resources["t1"].Structure.Schema.Fields), len(sel.SelectExprs))
	}
}
