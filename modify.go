package sqlparser

import "errors"

// PushStandardColumns pushes columns onto the front of a
// DDL statement's
// TODO - break DDL out into seperate structs?
func (stmt *DDL) PushColumnDefs(defs ...*ColDef) error {
	if stmt.Action == CreateStr {
		stmt.ColDefs = append(defs, stmt.ColDefs...)
	} else {
		return errors.New("this statement is not a create table statement")
	}

	return nil
}
