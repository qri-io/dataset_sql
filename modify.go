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

// AddStandardColumns adds id, created, and updated columns to a
// CREATE TABLE ddl statement.
// TODO - this should do additional checks on the existing columns
// to make sure they're kosher
func (stmt *DDL) AddStdColumns() error {
	if stmt.Action == CreateStr {
		if err := stmt.PushColumnDefs(
			&ColDef{ColName: &TableName{Name: "id"}, ColType: &DataType{Type: "uuid"}, Constraints: ColConstrs{&ColConstr{Constraint: ColConstrPrimaryKeyStr}}},
			&ColDef{ColName: &TableName{Name: "created"}, ColType: &DataType{Type: "integer"}},
			&ColDef{ColName: &TableName{Name: "updated"}, ColType: &DataType{Type: "integer"}},
		); err != nil {
			return err
		}
	} else {
		return errors.New("this statement is not a create table statement")
	}

	return nil
}
