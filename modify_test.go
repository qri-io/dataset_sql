package sqlparser

import "testing"

func TestPushColDefs(t *testing.T) {
	stmt, err := Parse("CREATE TABLE a (noize text NOT NULL)")
	if err != nil {
		t.Errorf("Parse error: %s", err)
	}

	ddlStmt, ok := stmt.(*DDL)
	if !ok {
		t.Error("statement didn't parse into a DDL")
	} else {
		if err := ddlStmt.PushColumnDefs(
			&ColDef{ColName: &TableName{Name: "id"}, ColType: "uuid", Constraints: ColConstrs{&ColConstr{Constraint: ColConstrPrimaryKeyStr}}},
			&ColDef{ColName: &TableName{Name: "created"}, ColType: "integer"},
			&ColDef{ColName: &TableName{Name: "updated"}, ColType: "integer"},
		); err != nil {
			t.Errorf("PushColumnDefs error: %s", err)
		}

		// buf := NewTrackedBuffer(nil)
		// ddlStmt.Format(buf)
		// fmt.Println(buf.String())
	}
}
