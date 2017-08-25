// TODO - move this to a separate package
package dataset_sql

import (
	"bytes"
	"strconv"

	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/datatypes"
)

type JsonWriter struct {
	writeObjects bool
	rowsWritten  int
	st           *dataset.Structure
	buf          *bytes.Buffer
}

func NewJsonWriter(st *dataset.Structure, writeObjects bool) *JsonWriter {
	return &JsonWriter{
		writeObjects: writeObjects,
		st:           st,
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
		f := w.st.Schema.Fields[i]
		ent := []byte(",\"" + f.Name + "\":")
		if i == 0 {
			ent = ent[1:]
		}
		if c == nil {
			ent = append(ent, []byte("null")...)
		} else {
			switch f.Type {
			case datatypes.String:
				ent = append(ent, []byte(strconv.Quote(string(c)))...)
			case datatypes.Float, datatypes.Integer:
				ent = append(ent, c...)
			case datatypes.Boolean:
				// TODO - coerce to true & false specifically
				ent = append(ent, c...)
			default:
				ent = append(ent, []byte(strconv.Quote(string(c)))...)
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
		f := w.st.Schema.Fields[i]
		ent := []byte(",")
		if i == 0 {
			ent = ent[1:]
		}
		if c == nil {
			ent = append(ent, []byte("null")...)
		} else {
			switch f.Type {
			case datatypes.String:
				ent = append(ent, []byte(strconv.Quote(string(c)))...)
			case datatypes.Float, datatypes.Integer:
				if len(c) == 0 {
					ent = append(ent, []byte("0")...)
				} else {
					ent = append(ent, c...)
				}
			case datatypes.Boolean:
				// TODO - coerce to true & false specifically
				if len(c) == 0 {
					ent = append(ent, []byte("false")...)
				}
				ent = append(ent, c...)
			default:
				ent = append(ent, []byte(strconv.Quote(string(c)))...)
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
