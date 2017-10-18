package dataset_sql

import (
	"bytes"

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
	srg := &SourceRowGenerator{store: store, init: true}
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
	init    bool
}

func (srg *SourceRowGenerator) Next() bool {
	if srg.readers[0].done {
		return false
	}

	// need init to skip initial call to Next.
	if srg.init {
		srg.init = false
		return true
	}

	srg.incrRow()
	return true
}

func (srg *SourceRowGenerator) incrRow() error {
	for i := len(srg.readers) - 1; i >= 0; i-- {
		rdr := srg.readers[i]
		if rdr.done {
			return rdr.Reset(srg.store)
		} else {
			return rdr.Next()
		}
	}
	return nil
}

func (srg *SourceRowGenerator) Row() SourceRow {
	sr := SourceRow{}
	for _, rdr := range srg.readers {
		sr[rdr.name] = rdr.row
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
	if err != nil {
		rr.done = true
	}
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

func NewSourceRowFilter(ast Statement) (*SourceRowFilter, error) {
	// TODO - generalize limit/offset to Statement searching
	// limit, offset, err := ast.Limit.Counts()
	// if err != nil {
	// 	return nil, err
	// }

	funcs, err := AggregateFuncs(ast)
	if err != nil {
		return nil, err
	}

	return &SourceRowFilter{
		limit:   -1,
		offset:  0,
		calcAll: true,
		agg:     len(funcs) > 0,
	}, nil
}

// SourceRowFilter uses type-populated AST to evaluate candidate SourceRows
// to see if they should be added to the resulting dataset internal state
// for example, things like current status in a LIMIT / OFFSET
type SourceRowFilter struct {
	ast     Statement
	added   int
	limit   int
	offset  int
	calcAll bool
	agg     bool
}

// Filter returns weather the row should be allowed to pass through
// to the table
func (srf *SourceRowFilter) Filter(sr SourceRow) bool {
	return !srf.agg
}

// Done indicates we don't need to iterate anymore
func (srf *SourceRowFilter) Done() bool {
	// TODO - lots of things will complicate this clause, such
	// as needing to calculate all results to sort, etc.
	return !srf.calcAll && srf.limit > 0 && srf.added == srf.limit
}

// rowsEqual checks to see if two rows are identitical
func rowsEqual(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ai := range a {
		if !bytes.Equal(ai, b[i]) {
			return false
		}
	}
	return true
}
