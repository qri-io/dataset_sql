package dataset_sql

import (
	"github.com/qri-io/dataset"
)

// Dataset is a fully dereferenced version of github.com/qri-io/dataset.Dataset
// needed to do internal work. This should move somewhere more sensible, like,
// maybe, the dataset package
type Dataset struct {
	Metadata  *dataset.Dataset
	Resources map[string]*StructuredData
	Query     *dataset.Query
	Structure *dataset.Structure
}

type StructuredData struct {
	Structure *dataset.Structure
	Data      []byte
}
