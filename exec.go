package dataset_sql

import (
	"github.com/qri-io/cafs"
	"github.com/qri-io/dataset"
)

type ExecOpt struct {
	Format dataset.DataFormat
}

func DefaultExecOpts() *ExecOpt {
	return &ExecOpt{
		Format: dataset.CSVDataFormat,
	}
}

func opts(options ...func(*ExecOpt)) *ExecOpt {
	o := &ExecOpt{
		Format: dataset.CSVDataFormat,
	}
	for _, option := range options {
		option(o)
	}
	return o
}

func Exec(store cafs.Filestore, query *dataset.Transform, options ...func(o *ExecOpt)) (result *dataset.Transform, resultBytes []byte, err error) {
	opts := &ExecOpt{
		Format: dataset.CSVDataFormat,
	}
	for _, option := range options {
		option(opts)
	}

	stmt, abst, err := Format(query)
	if err != nil {
		return nil, nil, err
	}

	return stmt.exec(store, query, abst)
}

func (stmt *Select) exec(store cafs.Filestore, query, absq *dataset.Transform) (result *dataset.Transform, resultBytes []byte, err error) {
	cols := CollectColNames(stmt)
	buf, err := NewResultBuffer(stmt, absq)
	if err != nil {
		return result, nil, err
	}
	srg, err := NewSourceRowGenerator(store, absq.Resources)
	if err != nil {
		return result, nil, err
	}
	srf, err := NewSourceRowFilter(stmt, buf)
	if err != nil {
		return result, nil, err
	}

	rrg, err := NewResultRowGenerator(stmt, absq.Structure)
	if err != nil {
		return result, nil, err
	}

	for srg.Next() && !srf.Done() {
		sr, err := srg.Row()
		if err != nil {
			return result, nil, err
		}

		if err := SetSourceRow(cols, sr); err != nil {
			return result, nil, err
		}

		if srf.Match() {
			row, err := rrg.GenerateRow()
			if err == ErrAggStmt {
				continue
			} else if err != nil {
				return result, nil, err
			}

			if srf.ShouldWriteRow(row) {
				if err := buf.WriteRow(row); err != nil {
					return result, nil, err
				}
			}

		}
	}

	if rrg.HasAggregates() {
		row, err := rrg.GenerateAggregateRow()
		if err != nil {
			return result, nil, err
		}
		buf.WriteRow(row)
	}

	if err := buf.Close(); err != nil {
		return result, nil, err
	}

	// TODO - rename / deref result var
	result = query
	resultBytes = buf.Bytes()
	return
}

func (node *Union) exec(store cafs.Filestore, query, abst *dataset.Transform) (*dataset.Transform, []byte, error) {
	return nil, nil, NotYetImplemented("union statements")
}
func (node *Insert) exec(store cafs.Filestore, query, abst *dataset.Transform) (*dataset.Transform, []byte, error) {
	return nil, nil, NotYetImplemented("insert statements")
}
func (node *Update) exec(store cafs.Filestore, query, abst *dataset.Transform) (*dataset.Transform, []byte, error) {
	return nil, nil, NotYetImplemented("update statements")
}
func (node *Delete) exec(store cafs.Filestore, query, abst *dataset.Transform) (*dataset.Transform, []byte, error) {
	return nil, nil, NotYetImplemented("delete statements")
}
func (node *Set) exec(store cafs.Filestore, query, abst *dataset.Transform) (*dataset.Transform, []byte, error) {
	return nil, nil, NotYetImplemented("set statements")
}
func (node *DDL) exec(store cafs.Filestore, query, abst *dataset.Transform) (*dataset.Transform, []byte, error) {
	return nil, nil, NotYetImplemented("ddl statements")
}
func (node *ParenSelect) exec(store cafs.Filestore, query, abst *dataset.Transform) (*dataset.Transform, []byte, error) {
	return nil, nil, NotYetImplemented("ParenSelect statements")
}
func (node *Show) exec(store cafs.Filestore, query, abst *dataset.Transform) (*dataset.Transform, []byte, error) {
	return nil, nil, NotYetImplemented("Show statements")
}
func (node *Use) exec(store cafs.Filestore, query, abst *dataset.Transform) (*dataset.Transform, []byte, error) {
	return nil, nil, NotYetImplemented("Use statements")
}
func (node *OtherRead) exec(store cafs.Filestore, query, abst *dataset.Transform) (*dataset.Transform, []byte, error) {
	return nil, nil, NotYetImplemented("OtherRead statements")
}
func (node *OtherAdmin) exec(store cafs.Filestore, query, abst *dataset.Transform) (*dataset.Transform, []byte, error) {
	return nil, nil, NotYetImplemented("OtherAdmin statements")
}
