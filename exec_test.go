package dataset_sql

import (
	"bytes"
	"encoding/csv"
	"github.com/ipfs/go-datastore"
	"github.com/qri-io/cafs"
	"github.com/qri-io/cafs/memfs"
	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/datatypes"
	"github.com/qri-io/dataset/dsfs"
	"github.com/qri-io/dataset/generate"
	"testing"
)

type execTestCase struct {
	statement string
	expect    error
	fields    []*dataset.Field
	numRows   int
}

func TestSelectFields(t *testing.T) {
	created := &dataset.Field{Name: "created", Type: datatypes.Date}
	title := &dataset.Field{Name: "title", Type: datatypes.String}
	views := &dataset.Field{Name: "views", Type: datatypes.Integer}
	rating := &dataset.Field{Name: "rating", Type: datatypes.Float}
	notes := &dataset.Field{Name: "notes", Type: datatypes.String}

	t1Struct := generate.RandomStructure(func(o *generate.RandomStructureOpts) {
		o.Format = dataset.CsvDataFormat
		o.Fields = []*dataset.Field{created, title, views, rating, notes}
	})

	t1Data := generate.RandomData(t1Struct, func(o *generate.RandomDataOpts) {
		o.Data = []byte("Sun Dec 25 09:25:46 2016,test_title,68882,0.6893978118896484,no notes\n")
		o.NumRandRecords = 9
	})

	t1 := &dataset.Dataset{
		Data:      datastore.NewKey("t1Data"),
		Structure: t1Struct,
	}

	t2Struct := generate.RandomStructure(func(o *generate.RandomStructureOpts) {
		o.Format = dataset.CsvDataFormat
		o.Fields = []*dataset.Field{created, title, views, rating, notes}
	})

	t2Data := generate.RandomData(t2Struct, func(o *generate.RandomDataOpts) {
		o.Data = []byte("Sun Dec 25 09:25:46 2016,test_title_two,68882,0.6893978118896484,no notes\n")
		o.NumRandRecords = 9
	})

	t2 := &dataset.Dataset{
		Data:      datastore.NewKey("t2"),
		Structure: t2Struct,
	}

	// store := datastore.NewMapDatastore()
	store := memfs.NewMapstore()
	t1DataPath, err := store.Put(memfs.NewMemfileBytes("t1data", t1Data), true)
	if err != nil {
		t.Error(err)
		return
	}
	t1.Data = t1DataPath
	t1path, err := dsfs.SaveDataset(store, t1, true)
	if err != nil {
		t.Error(err)
		return
	}

	t2DataPath, err := store.Put(memfs.NewMemfileBytes("t2Data", t2Data), true)
	if err != nil {
		t.Error(err)
		return
	}
	t2.Data = t2DataPath
	t2path, err := dsfs.SaveDataset(store, t2, true)
	if err != nil {
		t.Error(err)
		return
	}

	cases := []execTestCase{
		// {"select * from t1", nil, []*dataset.Field{created, title, views, rating, notes}, 10},
		{"select created, title, views, rating, notes from t1", nil, []*dataset.Field{created, title, views, rating, notes}, 10},
		{"select created from t1 limit 5", nil, []*dataset.Field{created}, 5},
		// {"select t1.created from t1 limit 1 offset 1", nil, []*dataset.Field{created}, 1},
		// {"select * from t1 where title = 'test_title'", nil, []*dataset.Field{created, title, views, rating, notes}, 1},
		// {"select * from t2 where title = 'test_title'", nil, []*dataset.Field{created, title, views, rating, notes}, 0},
		// {"select * from t2 where title = 'test_title_two'", nil, []*dataset.Field{created, title, views, rating, notes}, 1},
		// {"select * from t1, t2", nil, []*dataset.Field{created, title, views, rating, notes, created, title, views, rating, notes}, 100},
		// {"select * from t1, t2 where t1.notes = t2.notes", nil, []*dataset.Field{created, title, views, rating, notes, created, title, views, rating, notes}, 1},
		{"select sum(5) from t1", nil, []*dataset.Field{
			&dataset.Field{Name: "sum", Type: datatypes.Float},
		}, 1},
		// {"select sum(views), avg(views), count(views), max(views), min(views) from t1", nil, []*dataset.Field{
		// 	&dataset.Field{Name: "sum", Type: datatypes.Float},
		// 	&dataset.Field{Name: "avg", Type: datatypes.Float},
		// 	&dataset.Field{Name: "count", Type: datatypes.Float},
		// 	&dataset.Field{Name: "max", Type: datatypes.Float},
		// 	&dataset.Field{Name: "min", Type: datatypes.Float},
		// }, 1},
		// TODO - need to check result structure name on this one:
		// {"select * from a as aa, b as bb where a.created = b.created", nil, []*dataset.Field{created, title, views, rating, notes, created, title, views, rating, notes}, 2},
		// {"select 1 from a", nil, []*dataset.Field{&dataset.Field{Name: "result", Type: datatypes.Integer}}, 1},
	}

	t1ds, err := dsfs.LoadDataset(store, t1path)
	if err != nil {
		t.Error(err)
		return
	}
	t2ds, err := dsfs.LoadDataset(store, t2path)
	if err != nil {
		t.Error(err)
		return
	}

	ns := map[string]*dataset.Dataset{
		"t1": t1ds,
		"t2": t2ds,
	}

	runCases(store, ns, cases, t)
	// for i, c := range cases {
	// 	stmt, err := Parse(c.statement)
	// 	if err != nil {
	// 		t.Errorf("case %d parse error: %s", i, err.Error())
	// 		continue
	// 	}

	// 	results, data, err := stmt.Exec(ns)
	// 	if err != c.expect {
	// 		t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.expect, err.Error())
	// 		continue
	// 	}

	// 	if len(results.Fields) != len(c.fields) {
	// 		t.Errorf("case %d field length mismatch. expected: %d, got: %d", i, len(c.fields), len(results.Fields))
	// 		continue
	// 	}

	// 	for j, f := range c.fields {
	// 		if results.Fields[j].Name != f.Name {
	// 			t.Errorf("case %d field %d name mismatch. expected: %s, got: %s", i, j, f.Name, results.Fields[j].Name)
	// 			continue
	// 		}
	// 		if results.Fields[j].Type != f.Type {
	// 			t.Errorf("case %d field %d type mismatch. expected: %s, got: %s", i, j, f.Type, results.Fields[j].Type)
	// 			continue
	// 		}
	// 	}

	// 	r := csv.NewReader(bytes.NewBuffer(data))
	// 	records, err := r.ReadAll()
	// 	if err != nil {
	// 		t.Error(err.Error())
	// 		continue
	// 	}

	// 	if len(records) != c.numRows {
	// 		t.Errorf("case %d result count mismatch. expected: %d, got: %d", i, c.numRows, len(records))
	// 		continue
	// 	}

	// 	// table := tablewriter.NewWriter(os.Stdout)
	// 	// table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	// 	// table.SetCenterSeparator("|")
	// 	// table.SetHeader(results.FieldNames())
	// 	// table.AppendBulk(records)
	// 	// table.Render()
	// }
}

// func TestNullValues(t *testing.T) {
// 	created := &dataset.Field{Name: "created", Type: datatypes.Date}
// 	title := &dataset.Field{Name: "title", Type: datatypes.String}
// 	views := &dataset.Field{Name: "views", Type: datatypes.Integer}
// 	rating := &dataset.Field{Name: "rating", Type: datatypes.Float}
// 	notes := &dataset.Field{Name: "notes", Type: datatypes.String}

// 	ds := generate.RandomStructure(func(o *generate.RandomStructureOpts) {
// 		// o.Name = "null_values_test"
// 		// o.Address = dataset.NewAddress("test.null_values_test")
// 		o.Fields = []*dataset.Field{created, title, views, rating, notes}
// 		o.Data = []byte(",,,,\n")
// 		o.NumRandRecords = 0
// 	})

// 	airportCodes, err := loadTestData("airport_codes")
// 	if err != nil {
// 		t.Errorf("error loading test data '%s': %s", "airport_codes", err.Error())
// 		return
// 	}

// 	ns := mem.NewNamespace(dataset.NewAddress("test"), []*dataset.Resource{ds, airportCodes}, nil)

// 	runCases([]execTestCase{
// 		{"select * from test.null_values_test", nil, []*dataset.Field{created, title, views, rating, notes}, 1},
// 		{"select * from okfn.airport_codes limit 500", nil, airportCodes.Fields, 500},
// 	}, ns, t)

// 	// _, data, err := stmt.Exec(ns)
// 	// if err != nil {
// 	// 	t.Errorf("unexpected error executing statement: %s", err.Error())
// 	// 	return
// 	// }

// 	// // for j, f := range c.fields {
// 	// // 	if results.Fields[j].Name != f.Name {
// 	// // 		t.Errorf("case %d field %d name mismatch. expected: %s, got: %s", i, j, f.Name, results.Fields[j].Name)
// 	// // 		return
// 	// // 	}
// 	// // 	if results.Fields[j].Type != f.Type {
// 	// // 		t.Errorf("case %d field %d type mismatch. expected: %s, got: %s", i, j, f.Type, results.Fields[j].Type)
// 	// // 		return
// 	// // 	}
// 	// // }

// 	// r := csv.NewReader(bytes.NewBuffer(data))
// 	// records, err := r.ReadAll()
// 	// if err != nil {
// 	// 	t.Error(err.Error())
// 	// 	return
// 	// }

// 	// if len(records) != 1 {
// 	// 	t.Errorf("case result count mismatch. expected: %d, got: %d", 1, len(records))
// 	// 	return
// 	// }

// }

// func loadTestData(dir string) (*dataset.Resource, error) {
// 	dsData, err := ioutil.ReadFile(filepath.Join("test_data", dir, dataset.Filename))
// 	if err != nil {
// 		return nil, err
// 	}
// 	ds := &dataset.Resource{}
// 	if err := ds.UnmarshalJSON(dsData); err != nil {
// 		return nil, err
// 	}

// 	data, err := ioutil.ReadFile(filepath.Join("test_data", dir, ds.File))
// 	if err != nil {
// 		return nil, err
// 	}

// 	ds.Data = data
// 	return ds, nil
// }

func runCases(store cafs.Filestore, ns map[string]*dataset.Dataset, cases []execTestCase, t *testing.T) {
	for i, c := range cases {

		ds := &dataset.Dataset{
			QueryString: c.statement,
			QuerySyntax: "sql",
			Resources:   ns,
		}
		// q := &dataset.Query{
		// 	Syntax:    "sql",
		// 	Resources: ns,
		// 	Statement: c.statement,
		// }

		// stmt, err := Parse(c.statement)
		// if err != nil {
		// 	t.Errorf("case %d parse error: %s", i, err.Error())
		// 	continue
		// }

		results, data, err := Exec(store, ds, func(o *ExecOpt) {
			o.Format = dataset.CsvDataFormat
		})
		if err != c.expect {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.expect, err.Error())
			continue
		}

		if len(results.Schema.Fields) != len(c.fields) {
			t.Errorf("case %d field length mismatch. expected: %d, got: %d", i, len(c.fields), len(results.Schema.Fields))
			// fmt.Println(c.fields)
			// fmt.Println(results.Schema.FieldNames())
			continue
		}

		for j, f := range c.fields {
			if results.Schema.Fields[j].Name != f.Name {
				t.Errorf("case %d field %d name mismatch. expected: %s, got: %s", i, j, f.Name, results.Schema.Fields[j].Name)
				break
			}
			if results.Schema.Fields[j].Type != f.Type {
				t.Errorf("case %d field %d type mismatch. expected: %s, got: %s", i, j, f.Type, results.Schema.Fields[j].Type)
				break
			}
		}

		r := csv.NewReader(bytes.NewBuffer(data))
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
