package dataset_sql

import (
	"testing"

	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/datatypes"
)

func TestFormat(t *testing.T) {

	q1 := &dataset.Query{
		Abstract: &dataset.AbstractQuery{
			Statement: "select city, amount, date from precip",
		},
		Resources: map[string]*dataset.Dataset{
			"precip": &dataset.Dataset{
				Structure: &dataset.Structure{
					Format: dataset.CSVDataFormat,
					Schema: &dataset.Schema{
						Fields: []*dataset.Field{
							&dataset.Field{Name: "city", Type: datatypes.String},
							&dataset.Field{Name: "amount", Type: datatypes.Float},
							&dataset.Field{Name: "precip_type", Type: datatypes.String},
							&dataset.Field{Name: "date", Type: datatypes.Date},
						},
					},
				},
			},
		},
	}

	_, resources, err := makeTestStore()
	if err != nil {
		t.Errorf("error creating test data: %s", err.Error())
		return
	}

	q2 := &dataset.Query{
		Abstract: &dataset.AbstractQuery{
			Statement: "select * from one, two where one.title = two.title order by one.views desc",
		},
		Resources: map[string]*dataset.Dataset{
			"one": resources["t1"],
			"two": resources["t2"],
		},
	}

	cases := []struct {
		q       *dataset.Query
		stmtStr string
		remap   map[string]string
		err     string
	}{
		{q1, "select t1.a, t1.b, t1.d from t1", map[string]string{"t1": "precip"}, ""},
		{q2, "select * from t1, t2 where t1.b = t2.b order by t1.c desc", map[string]string{"t1": "one", "t2": "two"}, ""},
	}

	for i, c := range cases {
		stmtStr, _, remap, err := Format(c.q)
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
