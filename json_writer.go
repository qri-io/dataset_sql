package dataset_sql

import (
	"bytes"

	"github.com/qri-io/dataset"
	"github.com/qri-io/datatype"
)

type JsonWriter struct {
	writeObjects bool
	rowsWritten  int
	ds           *dataset.Dataset
	buf          *bytes.Buffer
}

func NewJsonWriter(ds *dataset.Dataset, writeObjects bool) *JsonWriter {
	return &JsonWriter{
		writeObjects: writeObjects,
		ds:           ds,
		buf:          bytes.NewBuffer([]byte{'['}),
	}
}

func (w *JsonWriter) WriteRow(row [][]byte) error {
	if w.writeObjects {
		return w.writeObjectRow(row)
	}
	return w.writeArrayRow(row)
}

func (w *JsonWriter) writeObjectRow(row [][]byte) error {
	enc := []byte{',', '\n', '{'}
	if w.rowsWritten == 0 {
		enc = enc[1:]
	}
	for i, c := range row {
		f := w.ds.Fields[i]
		ent := []byte(",\"" + f.Name + "\":")
		if i == 0 {
			ent = ent[1:]
		}
		if c == nil {
			ent = append(ent, []byte("null")...)
		} else {
			switch f.Type {
			case datatype.String:
				ent = append(ent, []byte("\""+string(c)+"\"")...)
			case datatype.Float, datatype.Integer:
				ent = append(ent, c...)
			case datatype.Boolean:
				// TODO - coerce to true & false specifically
				ent = append(ent, c...)
			default:
				ent = append(ent, []byte("\""+string(c)+"\"")...)
			}
		}

		enc = append(enc, ent...)
	}

	enc = append(enc, '}')
	if _, err := w.buf.Write(enc); err != nil {
		return err
	}

	w.rowsWritten++
	return nil
}

func (w *JsonWriter) writeArrayRow(row [][]byte) error {
	enc := []byte{',', '\n', '['}
	if w.rowsWritten == 0 {
		enc = enc[1:]
	}
	for i, c := range row {
		f := w.ds.Fields[i]
		ent := []byte(",")
		if i == 0 {
			ent = ent[1:]
		}
		if c == nil {
			ent = append(ent, []byte("null")...)
		} else {
			switch f.Type {
			case datatype.String:
				ent = append(ent, []byte("\""+string(c)+"\"")...)
			case datatype.Float, datatype.Integer:
				ent = append(ent, c...)
			case datatype.Boolean:
				// TODO - coerce to true & false specifically
				ent = append(ent, c...)
			default:
				ent = append(ent, []byte("\""+string(c)+"\"")...)
			}
		}

		enc = append(enc, ent...)
	}

	enc = append(enc, ']')
	if _, err := w.buf.Write(enc); err != nil {
		return err
	}

	w.rowsWritten++
	return nil
}

func (w *JsonWriter) Close() error {
	_, err := w.buf.Write([]byte{'\n', ']'})
	return err
}

func (w *JsonWriter) Bytes() []byte {
	return w.buf.Bytes()
}
