package dataset_sql

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"

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

func (node *FuncExpr) Datatype() datatypes.Type {
	switch node.Name.Lowered() {
	case "sum":
		return datatypes.Float
	}
	return datatypes.Any
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

func (af *AggFuncSum) Datatype() datatypes.Type {
	return datatypes.Float
}

func (af *AggFuncSum) Eval(row [][]byte) (q.Type, []byte, error) {
	fmt.Printf("%#v\n", row)
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
			fmt.Println("adding int", v)
			af.value = af.value + float32(v)
		case q.Type_FLOAT32:
			v, err := readFloat32(val)
			if err != nil {
				return q.Type_NULL_TYPE, nil, err
			}
			fmt.Println("adding float", v)
			af.value = af.value + v
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

func readFloat32(data []byte) (float32, error) {
	f64, err := strconv.ParseFloat(string(data), 32)
	return float32(f64), err
}

// type AggFuncVarPop struct{ Value float32 }
// type AggFuncVarSamp struct{ Value float32 }
// type AggFuncVariance struct{ Value float32 }
