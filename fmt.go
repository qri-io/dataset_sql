package dataset_sql

// Format places an sql statement in it's standard form.
// This will be *heavily* refined, improved, and moved into a
// separate package
// TODO - milestone & break down this core piece of tech
func Format(sql string) (string, error) {
	stmt, err := Parse(sql)
	if err != nil {
		return "", err
	}

	buf := NewTrackedBuffer(nil)
	stmt.Format(buf)

	return buf.String(), nil
}
