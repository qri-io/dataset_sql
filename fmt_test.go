package dataset_sql

import (
	"testing"

	"github.com/qri-io/dataset"
)

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
			"",
			map[string]string{"precip": "nonexistent"},
			nil,
			"couldn't find resource for table name: precip",
		},
		{
			"select * from one, two where one.title = two.title order by one.views desc",
			"select * from t1, t2 where t1.b = t2.b order by t1.c desc",
			map[string]string{"one": "t1", "two": "t2"},
			map[string]string{"t1": "one", "t2": "two"},
			"",
		},
		{
			"select * from foo, bar where foo.title = bar.title order by bar.views desc",
			"select * from t1, t2 where t1.b = t2.b order by t2.c desc",
			map[string]string{"foo": "t1", "bar": "t2"},
			map[string]string{"t1": "foo", "t2": "bar"},
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

		q := &dataset.Query{
			Abstract: &dataset.AbstractQuery{
				Statement: c.inStmt,
			},
			Resources: r,
		}

		stmtStr, _, remap, err := Format(q)
		if !(err == nil && c.err == "" || err != nil && err.Error() == c.err) {
			t.Errorf("case %d error mismatch. expected: '%s', got: '%s'", i, c.err, err)
			continue
		}

		if stmtStr != c.stmtStr {
			t.Errorf("case %d statement mismatch:\nexpected:\n'%s',\ngot:\n'%s'", i, c.stmtStr, stmtStr)
			continue
		}

		if len(c.remap) != len(remap) {
			t.Errorf("case %d remap length mismatch. expected: '%d', got: '%d'", i, len(c.remap), len(remap))
			t.Errorf("%v", remap)
			continue
		}

		for key, v := range c.remap {
			if remap[key] != v {
				t.Errorf("case %d key %s mismatch. expected: '%s', got: '%s'", i, key, v, remap[key])
				break
			}
		}
	}
}
