package print_ast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	sql "github.com/qri-io/dataset_sql"
)

func PrintAst(stmt string) ([]byte, []byte, error) {
	tree, err := sql.Parse(stmt)
	if err != nil {
		return nil, nil, err
	}

	ttree := NewTypedSQLNode(tree)
	buf := &bytes.Buffer{}

	var addWalk func(d int, node *TypedSQLNode) func(node sql.SQLNode) (bool, error)
	addWalk = func(d int, parent *TypedSQLNode) func(node sql.SQLNode) (bool, error) {
		fmt.Fprintf(buf, "%s%s\n", strings.Repeat("  ", d), parent.NodeType)
		return func(node sql.SQLNode) (bool, error) {
			child := NewTypedSQLNode(node)
			parent.Children = append(parent.Children, child)
			child.WalkSubtree(addWalk(d+1, child))
			return false, nil
		}
	}
	ttree.WalkSubtree(addWalk(0, ttree))
	jsondata, err := json.MarshalIndent(ttree, "", "  ")

	return buf.Bytes(), jsondata, err
}

type TypedSQLNode struct {
	NodeType string `json:"type"`
	sql.SQLNode
	Children []*TypedSQLNode
}

func NewTypedSQLNode(node sql.SQLNode) *TypedSQLNode {
	return &TypedSQLNode{
		NodeType: reflect.TypeOf(node).String(),
		SQLNode:  node,
	}
}

// func NodeType(node sql.SQLNode) string {
//   return reflect.TypeOf(node).String()
// }
