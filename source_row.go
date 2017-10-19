package dataset_sql

import (
	"bytes"
	"io"

	"github.com/ipfs/go-datastore"
	"github.com/qri-io/cafs"
	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/dsio"
)

// SourceRow is a row of data from a number of different tables
// identitifed by a string
type SourceRow map[string][][]byte

// SourcRowGenerator consumes dataset data readers
// generating SourceRows.
// It's main job is to generate the exhastive
// set of candidate row combinations for comparison
type SourceRowGenerator struct {
	store   cafs.Filestore
	readers []*rowReader
	init    bool
	err     error
}

// NewSourceRowGenerator initializes a source row generator
func NewSourceRowGenerator(store cafs.Filestore, resources map[string]*dataset.Dataset) (*SourceRowGenerator, error) {
	srg := &SourceRowGenerator{store: store, init: true}
	for name, ds := range resources {
		rdr := &rowReader{
			name: name,
			st:   ds.Structure,
			path: ds.Data,
		}
		if err := rdr.Reset(store); err != nil {
			return nil, err
		}
		srg.readers = append(srg.readers, rdr)
	}
	return srg, nil
}

func (srg *SourceRowGenerator) Next() bool {
	// need init to skip initial call to Next.
	if srg.init {
		srg.init = false
		return true
	}

	if err := srg.incrRow(); err != nil {
		if err == io.EOF {
			return false
		}
		srg.err = err
	}
	return true
}

func (srg *SourceRowGenerator) incrRow() error {
	for i := len(srg.readers) - 1; i >= 0; i-- {
		rdr := srg.readers[i]
		if err := rdr.Next(); err != nil {
			return err
		}

		if rdr.done {
			if i == 0 {
				return io.EOF
			}
			if err := rdr.Reset(srg.store); err != nil {
				return err
			}
		} else {
			return nil
		}
	}
	return nil
}

func (srg *SourceRowGenerator) Row() (SourceRow, error) {
	if srg.err != nil {
		return nil, srg.err
	}
	sr := SourceRow{}
	for _, rdr := range srg.readers {
		sr[rdr.name] = rdr.row
	}
	return sr, nil
}

// rowReader wraps a dsio.reader with additional required state
type rowReader struct {
	reader dsio.Reader
	name   string
	path   datastore.Key
	st     *dataset.Structure
	i      int
	done   bool
	row    [][]byte
}

// next increments the reader, pulling it's row data into
// internal state
func (rr *rowReader) Next() (err error) {
	if rr.done {
		return nil
	}

	rr.row, err = rr.reader.ReadRow()
	if err != nil {
		if err.Error() == "EOF" {
			rr.done = true
			return nil
		}
	}
	rr.i++
	return
}

// reset re-initializes the reader, starting the read process
// from scratch
func (rr *rowReader) Reset(store cafs.Filestore) error {
	f, err := store.Get(rr.path)
	if err != nil {
		return err
	}
	rr.i = 0
	rr.done = false
	rr.reader = dsio.NewReader(rr.st, f)
	return rr.Next()
}

// SourceRowFilter uses type-populated AST to evaluate candidate SourceRows
// to see if they should be added to the resulting dataset internal state
// for example, things like current status in a LIMIT / OFFSET
type SourceRowFilter struct {
	ast     Statement
	passed  int64
	limit   int64
	offset  int64
	test    *Where
	calcAll bool
}

func NewSourceRowFilter(ast Statement) (sr *SourceRowFilter, err error) {
	sr = &SourceRowFilter{}
	err = ast.WalkSubtree(func(node SQLNode) (bool, error) {
		if node == nil {
			return true, nil
		}
		switch n := node.(type) {
		case *Where:
			if n != nil {
				sr.test = n
			}
		case *Limit:
			if n != nil {
				sr.limit, sr.offset, err = n.Counts()
				if err != nil {
					return false, err
				}
			}
		case OrderBy:
			sr.calcAll = true
		}
		return true, nil
	})

	return
}

// Match returns weather the row should be allowed to pass through
// to the table
func (srf *SourceRowFilter) Match() bool {
	_, pass, err := srf.test.Eval()
	if err != nil {
		// fmt.Println(err.Error())
		return false
	}

	if bytes.Equal(pass, trueB) {
		srf.passed++
		if srf.passed < srf.offset {
			return false
		}
		return true
	}
	return false
}

// Done indicates we don't need to iterate anymore
func (srf *SourceRowFilter) Done() bool {
	// TODO - lots of things will complicate this clause, such
	// as needing to calculate all results to sort, etc.
	return !srf.calcAll && srf.limit > 0 && (srf.passed-srf.offset) >= srf.limit
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
