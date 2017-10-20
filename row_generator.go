package dataset_sql

import (
	"fmt"
	"github.com/qri-io/dataset"
)

// RowGenerator makes rows from SourceRows
// calling eval on a set of select expressions from a given
// SourceRow
type RowGenerator struct {
	exprs SelectExprs
	aggs  []AggFunc
	st    *dataset.Structure
}

func NewRowGenerator(sel *Select, result *dataset.Structure) (rg *RowGenerator, err error) {
	rg = &RowGenerator{
		exprs: sel.SelectExprs,
		st:    result,
	}

	rg.aggs, err = AggregateFuncs(sel)
	if err != nil {
		return nil, err
	}
	return
}

var (
	ErrAggStmt   = fmt.Errorf("this statement only generates an aggregate result row")
	ErrTableStmt = fmt.Errorf("this statement doesn't generate an aggregate result row")
)

// GenerateRow generates a row
func (rg *RowGenerator) GenerateRow() ([][]byte, error) {
	row := make([][]byte, len(rg.exprs))
	for i, expr := range rg.exprs {
		_, data, err := expr.Eval()
		if err != nil {
			return nil, err
		}
		row[i] = data
	}

	if !rg.HasAggregates() {
		return row, nil
	}
	return nil, ErrAggStmt
}

func (rg *RowGenerator) HasAggregates() bool {
	return len(rg.aggs) > 0
}

func (rg *RowGenerator) GenerateAggregateRow() ([][]byte, error) {
	if rg.HasAggregates() {

		row := make([][]byte, len(rg.exprs))
		for i, expr := range rg.exprs {
			_, data, err := expr.Eval()
			if err != nil {
				return nil, err
			}
			row[i] = data
		}
		return row, nil
	}
	return nil, ErrTableStmt
}

func (rg *RowGenerator) Structure() *dataset.Structure {
	return rg.st
}
