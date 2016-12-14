package dataset_sql

import (
	"testing"

	"github.com/qri-io/dataset"
	"github.com/qri-io/datatype"
)

func TestJsonWriter(t *testing.T) {
	cases := []struct {
		ds           *dataset.Dataset
		writeObjects bool
		entries      [][][]byte
		out          string
	}{
		{&dataset.Dataset{Fields: []*dataset.Field{&dataset.Field{Name: "a", Type: datatype.String}}}, true, [][][]byte{[][]byte{[]byte("hello")}}, "[\n{\"a\":\"hello\"}\n]"},
		{&dataset.Dataset{Fields: []*dataset.Field{&dataset.Field{Name: "a", Type: datatype.String}}}, false, [][][]byte{[][]byte{[]byte("hello")}}, "[\n[\"hello\"]\n]"},
		{&dataset.Dataset{Fields: []*dataset.Field{&dataset.Field{Name: "a", Type: datatype.String}}}, true, [][][]byte{
			[][]byte{[]byte("hello")},
			[][]byte{[]byte("world")},
		}, "[\n{\"a\":\"hello\"},\n{\"a\":\"world\"}\n]"},
		{&dataset.Dataset{Fields: []*dataset.Field{&dataset.Field{Name: "a", Type: datatype.String}}}, false, [][][]byte{
			[][]byte{[]byte("hello")},
			[][]byte{[]byte("world")},
		}, "[\n[\"hello\"],\n[\"world\"]\n]"},
	}

	for i, c := range cases {
		w := NewJsonWriter(c.ds, c.writeObjects)
		for _, ent := range c.entries {
			if err := w.WriteRow(ent); err != nil {
				t.Errorf("case %d WriteRow error: %s", i, err.Error())
				break
			}
		}
		if err := w.Close(); err != nil {
			t.Errorf("case %d Close error: %s", i, err.Error())
		}

		if string(w.Bytes()) != c.out {
			t.Errorf("case %d result mismatch. expected:\n%s\ngot:\n%s", i, c.out, string(w.Bytes()))
		}
	}
}
