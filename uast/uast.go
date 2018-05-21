// Copyright 2017 Sourced Technologies SL
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

// Package uast defines a UAST (Universal Abstract Syntax Tree) representation
// and operations to manipulate them.
package uast

import (
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/bblfsh/sdk.v2/uast/role"
)

// Special field keys for Object
const (
	KeyType  = "@type"  // InternalType
	KeyToken = "@token" // Token
	KeyRoles = "@role"  // Roles, for representations see RoleList
	KeyPos   = "@pos"   // All positional information is stored in this field
)

const (
	// NS is a namespace for the UAST types.
	NS = "uast"

	// TypePosition is a node type for positional information in AST. See AsPosition.
	TypePosition = NS + ":Position"
	// TypePositions is a node type for a root node of positional information in AST. See AsPositions.
	TypePositions = NS + ":Positions"
	// TypeOperator is a node type for an operator AST node. See Operator.
	TypeOperator = NS + ":Operator"
	// KeyPosOff is a name for a Position object field that stores a bytes offset.
	KeyPosOff = "offset"
	// KeyPosLine is a name for a Position object field that stores a source line.
	KeyPosLine = "line"
	// KeyPosCol is a name for a Position object field that stores a source column.
	KeyPosCol = "col"

	KeyStart = "start" // StartPosition
	KeyEnd   = "end"   // EndPosition
)

// Position represents a position in a source code file.
type Position struct {
	// Offset is the position as an absolute byte offset. It is a 0-based index.
	Offset uint32 `json:"offset"`
	// Line is the line number. It is a 1-based index.
	Line uint32 `json:"line"`
	// Col is the column number (the byte offset of the position relative to
	// a line. It is a 1-based index.
	Col uint32 `json:"col"`
}

// Positions is a container object that stores all positional information for a node.
type Positions map[string]Position

func (p Positions) Start() *Position {
	if p, ok := p[KeyStart]; ok {
		return &p
	}
	return nil
}
func (p Positions) End() *Position {
	if p, ok := p[KeyEnd]; ok {
		return &p
	}
	return nil
}

// AsPosition transforms a generic AST node to a Position object.
func AsPosition(m Object) *Position {
	if TypeOf(m) != TypePosition {
		return nil
	}
	var p Position
	if err := NodeAs(m, &p); err != nil {
		panic(err)
	}
	return &p
}

// PositionsOf returns an object with all positional information for a node.
func PositionsOf(m Object) Positions {
	o, _ := m[KeyPos].(Object)
	if len(o) == 0 {
		return nil
	}
	ps := make(Positions, len(o))
	for k, v := range o {
		po, _ := v.(Object)
		if p := AsPosition(po); p != nil {
			ps[k] = *p
		}
	}
	return ps
}

// ToObject converts Position to a generic AST node.
func (p Position) ToObject() Object {
	n, err := toNodeReflect(reflect.ValueOf(&p))
	if err != nil {
		panic(err)
	}
	return n.(Object)
}

// RoleList converts a set of roles into a list node.
func RoleList(roles ...role.Role) Array {
	arr := make(Array, 0, len(roles))
	for _, r := range roles {
		arr = append(arr, String(r.String()))
	}
	return arr
}

// RolesOf is a helper for getting node UAST roles (see KeyRoles).
func RolesOf(m Object) role.Roles {
	arr, ok := m[KeyRoles].(Array)
	if !ok || len(arr) == 0 {
		if tp := TypeOf(m); tp == "" || strings.HasPrefix(tp, NS+":") {
			return nil
		}
		return role.Roles{role.Unannotated}
	}
	out := make(role.Roles, 0, len(arr))
	for _, v := range arr {
		if r, ok := v.(String); ok {
			out = append(out, role.FromString(string(r)))
		}
	}
	return out
}

// TokenOf is a helper for getting node token (see KeyToken).
func TokenOf(m Object) string {
	t := m[KeyToken]
	s, ok := t.(String)
	if ok {
		return string(s)
	}
	v, _ := t.(Value)
	if v != nil {
		return fmt.Sprint(v)
	}
	return ""
}

// Tokens collects all tokens of the tree recursively (pre-order).
func Tokens(n Node) []string {
	var tokens []string
	WalkPreOrder(n, func(n Node) bool {
		if obj, ok := n.(Object); ok {
			if tok := TokenOf(obj); tok != "" {
				tokens = append(tokens, tok)
			}
		}
		return true
	})
	return tokens
}
