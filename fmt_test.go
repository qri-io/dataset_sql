package dataset_sql

import (
	"github.com/ipfs/go-datastore"
	"testing"

	"github.com/qri-io/dataset"
)

func TestQueryRecordPath(t *testing.T) {
	store, resources, err := makeTestStore()
	if err != nil {
		t.Errorf("error creating test data: %s", err.Error())
		return
	}

	cases := []struct {
		query *dataset.Transform
		hash  datastore.Key
		err   string
	}{
		{&dataset.Transform{
			Syntax: "sql",
			Data:   "select * from foo",
			Resources: map[string]*dataset.Dataset{
				"foo": resources["t1"],
			}}, datastore.NewKey("/map/QmXHyTpaiStbVAcxMnqyJNMGhaeZacpcKf4GLYqVqBYVoP"), ""},
	}

	for i, c := range cases {
		path, err := QueryRecordPath(store, c.query, func(o *ExecOpt) {
			o.Format = dataset.CSVDataFormat
		})
		if !(err == nil && c.err == "" || err != nil && err.Error() == c.err) {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, err)
			continue
		}

		if !c.hash.Equal(path) {
			t.Errorf("case %d hash mismatch. expected: %s, got: %s", i, c.hash.String(), path.String())
			continue
		}
	}
}

func TestFormat(t *testing.T) {
	_, resources, err := makeTestStore()
	if err != nil {
		t.Errorf("error creating test data: %s", err.Error())
		return
	}

	cases := []struct {
		inStmt        string
		stmtStr       string
		resourceNames map[string]string
		remap         map[string]string
		err           string
	}{
		{
			"select city, amount, date from precip",
			"select city, amount, `date` from precip",
			map[string]string{"precip": "nonexistent"},
			nil,
			"invalid resource reference: precip",
		},
		// TODO - this is wrong and kinda passes
		// {
		// 	"select * from one, two where one.title = two.title order by one.views desc",
		// 	"select t1.a as a, t1.b as b, t1.c as c, t1.d as d, t1.e as e from t1, t2 where t1.b = t2.b order by t1.c desc",
		// 	map[string]string{"one": "t1", "two": "t2"},
		// 	map[string]string{"t1": "one", "t2": "two"},
		// 	"",
		// },
		// TODO - this passes, and is wrong!
		// {
		// 	"select * from foo, bar where foo.title = bar.title order by bar.views desc",
		// 	"select t1.a as a, t1.b as b, t1.c as c, t1.d as d, t1.e as e from t1, t2 where t1.b = t2.b order by t2.c desc",
		// 	map[string]string{"foo": "t1", "bar": "t2"},
		// 	map[string]string{"t1": "foo", "t2": "bar"},
		// 	"",
		// },
		{
			"select sum(views), avg(views), count(views), max(views), min(views) from foo",
			"select sum(t1.c), avg(t1.c), count(t1.c), max(t1.c), min(t1.c) from t1",
			map[string]string{"foo": "t3"},
			map[string]string{"t1": "foo"},
			"",
		},
		{
			"select title from foo, bar where foo.title = bar.title order by bar.views desc",
			"",
			map[string]string{"foo": "t1", "bar": "t2"},
			nil,
			"column reference 'title' is ambiguous, please specify the dataset name for this table",
		},
	}

	for i, c := range cases {
		r := map[string]*dataset.Dataset{}
		for key, name := range c.resourceNames {
			r[key] = resources[name]
		}

		q := &dataset.Transform{
			Syntax:    "sql",
			Data:      c.inStmt,
			Resources: r,
		}

		stmt, abst, err := Format(q)
		if !(err == nil && c.err == "" || err != nil && err.Error() == c.err) {
			t.Errorf("case %d error mismatch. expected: '%s', got: '%s'", i, c.err, err)
			continue
		}

		if c.err == "" {
			stmtStr := String(stmt)
			if stmtStr != c.stmtStr {
				t.Errorf("case %d statement mismatch:\nexpected:\n'%s',\ngot:\n'%s'", i, c.stmtStr, stmtStr)
				continue
			}

			if len(c.remap) != len(abst.Resources) {
				t.Errorf("case %d remap length mismatch. expected: '%d', got: '%d'", i, len(c.remap), len(abst.Resources))
				t.Errorf("%v", abst)
				continue
			}

			// for key, v := range c.remap {
			// 	if remap[key] != v {
			// 		t.Errorf("case %d key %s mismatch. expected: '%s', got: '%s'", i, key, v, remap[key])
			// 		break
			// 	}
			// }
		}
	}
}
