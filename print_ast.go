package dataset_sql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func PrintAst(tree Statement) ([]byte, []byte, error) {
	ttree := NewTypedSQLNode(tree)
	buf := &bytes.Buffer{}

	var addWalk func(d int, node *TypedSQLNode) func(node SQLNode) (bool, error)
	addWalk = func(d int, parent *TypedSQLNode) func(node SQLNode) (bool, error) {
		fmt.Fprintf(buf, "%s%s\n", strings.Repeat("  ", d), parent.NodeType)
		return func(node SQLNode) (bool, error) {
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
	SQLNode
	Children []*TypedSQLNode
}

func NewTypedSQLNode(node SQLNode) *TypedSQLNode {
	return &TypedSQLNode{
		NodeType: reflect.TypeOf(node).String(),
		SQLNode:  node,
	}
}

// func NodeType(node SQLNode) string {
//   return reflect.TypeOf(node).String()
// }
