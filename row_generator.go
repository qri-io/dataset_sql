package dataset_sql

// RowGenerator makes rows from SourceRows
// calling eval on a set of select expressions from a given
// SourceRow
type RowGenerator struct {
	ast Statement
}

// GenerateRow generates a row
func (rg *RowGenerator) GenerateRow(sr SourceRow) ([][]byte, error) {
	return nil, nil
}
