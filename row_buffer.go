package dataset_sql

import (
	"fmt"

	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/dsio"
)

// NewResultBuffer returns either a *RowBuffer or *dsio.Buffer depending on
// which is required. RowBuffer is (much) more expensive but supports introspection
// into already-written rows
func NewResultBuffer(stmt Statement, aq *dataset.Transform) (dsio.RowReadWriter, error) {
	if needsRowBuffer(stmt) {
		cfg, err := statementRowBufferCfg(stmt, aq)
		if err != nil {
			return nil, err
		}
		return dsio.NewStructuredRowBuffer(aq.Structure, func(o *dsio.StructuredRowBufferCfg) { *o = *cfg })
	}
	return dsio.NewStructuredBuffer(aq.Structure)
}

// Checks to see if we need a RowBuffer at all. Statements that don't contain
// DISTINCT or ORDER BY clauses don't require row buffering
func needsRowBuffer(stmt Statement) bool {
	sel, ok := stmt.(*Select)
	if !ok {
		// TODO - remove this.
		// for now anything that isn't a select statement is a candidate for
		// row buffering
		return true
	}

	return len(sel.OrderBy) > 0 || sel.Distinct != ""
}

// statementRowBufferCfg gives a configuration for a StructuredRowBuffer based on a sql Statement
func statementRowBufferCfg(stmt Statement, aq *dataset.Transform) (*dsio.StructuredRowBufferCfg, error) {
	sel, ok := stmt.(*Select)
	if !ok {
		// TODO - need to implement this for all types of statements
		// need to add SelectExprs() SelectExprs and Orders() Orders
		// on Statement interface
		return nil, NotYetImplemented("non-select row ordering")
	}

	// TODO - I messed up & wrote StructuredRowBuffer to only accept one "desc" flag
	// need to fix that. For now this hack just checks the last desc flag
	desc := false

	orders := []*dataset.Field{}
	for _, o := range sel.OrderBy {
		if cn, ok := o.Expr.(*ColName); ok {
			st := aq.Resources[cn.Qualifier.String()].Structure
			if st == nil {
				return nil, fmt.Errorf("couldn't find abstract structure reference: %s", cn.Qualifier.String())
			}
			str := String(cn.Name)
			// str = string(bytes.TrimPrefix([]byte(str), []byte("t1.")))
			idx := st.StringFieldIndex(str)
			desc = o.Direction == "desc"
			if idx < 0 {
				return nil, fmt.Errorf("couldn't find sort index: %s", String(o.Expr))
			}
			orders = append(orders, st.Schema.Fields[idx])
		} else {
			return nil, fmt.Errorf("unsupported sort value: %s", String(o.Expr))
		}
	}

	return &dsio.StructuredRowBufferCfg{
		OrderBy:     orders,
		OrderByDesc: desc,
		Unique:      sel.Distinct != "",
	}, nil
}
