package dataset_sql

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/qri-io/dataset/datatypes"
	q "github.com/qri-io/dataset_sql/vt/proto/query"
)

type AggFunc interface {
	Eval(row [][]byte) (q.Type, []byte, error)
	Value() float32
}

// AggregateFuncs extracts a slice of any aggregate functions from an AST, while also writing
// pointers to newly-generated funcs to the AST
func AggregateFuncs(root SQLNode, schemas map[string]*StructureData) (funcs []AggFunc, err error) {
	err = root.WalkSubtree(func(node SQLNode) (bool, error) {
		switch t := node.(type) {
		case *FuncExpr:
			fn, err := t.Function(schemas)
			if err != nil {
				return false, err
			}
			funcs = append(funcs, fn)
		}
		return true, nil
	})
	return
}

// Function gives the backing function to perform
func (node *FuncExpr) Function(from map[string]*StructureData) (fn AggFunc, err error) {
	switch node.Name.Lowered() {
	case "sum":
		fn, err = node.newAggFuncSum(from)
		if err != nil {
			return
		}
		node.fn = fn
	default:
		return nil, fmt.Errorf("unrecognized aggregate function: %s", node.Name)
	}
	return fn, nil
}

// type AggFuncAvg struct{ Value float32 }
// type AggFuncBitAnd struct{ Value float32 }
// type AggFuncBitOr struct{ Value float32 }
// type AggFuncBitXor struct{ Value float32 }
// type AggFuncCount struct{ Value float32 }
// type AggFuncGroupConcat struct{ Value float32 }
// type AggFuncMax struct{ Value float32 }
// type AggFuncMin struct{ Value float32 }
// type AggFuncStd struct{ Value float32 }
// type AggFuncStddevPop struct{ Value float32 }
// type AggFuncStddevSamp struct{ Value float32 }
// type AggFuncStddev struct{ Value float32 }

func (node *FuncExpr) newAggFuncSum(from map[string]*StructureData) (AggFunc, error) {
	if !datatypes.EachNumeric(node.Exprs.FieldTypes(from)) {
		return nil, fmt.Errorf("sum only works with numeric fields")
	}

	return &AggFuncSum{
		Exprs: node.Exprs,
		value: 0,
	}, nil
}

type AggFuncSum struct {
	Exprs SelectExprs
	value float32
}

func (af *AggFuncSum) Eval(row [][]byte) (q.Type, []byte, error) {
	ts, vs, err := af.Exprs.Values(row)
	if err != nil {
		return q.Type_NULL_TYPE, nil, err
	}

	for i, val := range vs {
		switch ts[i] {
		case q.Type_INT64:
			v, err := readInt(val)
			if err != nil {
				return q.Type_NULL_TYPE, nil, err
			}
			af.value += float32(v)
		case q.Type_FLOAT32:
			af.value += readFloat32(val)
		}
	}

	// TODO - possible to debug by printing intermediate
	// steps here
	return q.Type_FLOAT32, nil, nil
}

func (af *AggFuncSum) Value() float32 {
	return af.value
}

func readInt(data []byte) (int64, error) {
	return binary.ReadVarint(bytes.NewBuffer(data))
}

func readFloat32(data []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(data))
}

// type AggFuncVarPop struct{ Value float32 }
// type AggFuncVarSamp struct{ Value float32 }
// type AggFuncVariance struct{ Value float32 }
