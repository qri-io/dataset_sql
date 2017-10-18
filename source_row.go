package dataset_sql

import (
	"github.com/ipfs/go-datastore"
	"github.com/qri-io/cafs"
	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/dsio"
)

// SourceRow is a row of data from a number of different tables
// identitifed by a string
type SourceRow map[string][][]byte

// NewSourceRowGenerator initializes a source row generator
func NewSourceRowGenerator(store cafs.Filestore, resources map[string]*dataset.Dataset) (*SourceRowGenerator, error) {
	srg := &SourceRowGenerator{store: store}
	for name, ds := range resources {
		rdr := &rowReader{
			name: name,
			st:   ds.Structure,
		}
		if err := rdr.Reset(store); err != nil {
			return nil, err
		}
		srg.readers = append(srg.readers, rdr)
	}
	return srg, nil
}

// SourcRowGenerator consumes dataset data readers
// generating SourceRows.
// It's main job is to generate the exhastive
// set of candidate row combinations for comparison
type SourceRowGenerator struct {
	store   cafs.Filestore
	readers []*rowReader
}

func (srg *SourceRowGenerator) Next() SourceRow {
	sr := SourceRow{}
	if srg.readers[0].i == srg.readers[0].length-1 {
		return nil
	}
	srg.incrRow()
	return srg.row()
}

func (srg *SourceRowGenerator) incrRow() error {
	for i := len(srg.readers) - 1; i >= 0; i-- {
		rdr := srg.readers[i]
		if rdr.i < rdr.length-1 {
			return rdr.Next()
		} else if rdr.i == rdr.length-1 && i > 0 {
			return rdr.Reset(srg.store)
		}
	}
	return nil
}

func (srg *SourceRowGenerator) row() SourceRow {
	sr := SourceRow{}
	for _, rdr := range srg.readers {
		SourceRow[rdr.name] == rdr.row
	}
	return sr
}

// rowReader wraps a dsio.reader with additional required state
type rowReader struct {
	reader dsio.Reader
	name   string
	key    datastore.Key
	st     *dataset.Structure
	i      int
	done   bool
	row    [][]byte
}

// next increments the reader, pulling it's row data into
// internal state
func (rr *rowReader) Next() (err error) {
	rr.i++
	rr.row, err = rr.reader.ReadRow()
	return
}

// reset re-initializes the reader, starting the read process
// from scratch
func (rr *rowReader) Reset(store cafs.Filestore) error {
	f, err := store.Get(rr.key)
	if err != nil {
		return err
	}
	rr.i = 0
	dsio.NewReader(rr.st, f)
	return nil
}

func NewSourceRowFilter(ast Statement) *SourceRowFilter {
	return &SourceRowFilter{}
}

// SourceRowFilter uses type-populated AST to evaluate candidate SourceRows
// to see if they should be added to the resulting dataset internal state
// for example, things like current status in a LIMIT / OFFSET
type SourceRowFilter struct {
	ast    Statement
	added  int
	limit  int
	offset int
}

// Filter returns weather the row should be allowed to pass through
// to the table
func (srf *SourceRowFilter) Filter() bool {
	return true, false
}
