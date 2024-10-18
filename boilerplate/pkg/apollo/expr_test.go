package apollo

import (
	"testing"

	"github.com/expr-lang/expr/ast"
	"github.com/stretchr/testify/assert"
)

func TestParseVariableNode(t *testing.T) {
	tests := []struct {
		input    Operand
		expected ast.Node
	}{
		{
			input: Operand{Values: []string{"a"}},
			expected: &ast.IdentifierNode{
				Value: "a",
			},
		},
		{
			input: Operand{Values: []string{"a.b"}},
			expected: &ast.MemberNode{
				Node: &ast.IdentifierNode{
					Value: "a",
				},
				Property: &ast.StringNode{
					Value: "b",
				},
			},
		},
		{
			input: Operand{Values: []string{"a.b.c"}},
			expected: &ast.MemberNode{
				Node: &ast.MemberNode{
					Node: &ast.IdentifierNode{
						Value: "a",
					},
					Property: &ast.StringNode{
						Value: "b",
					},
				},
				Property: &ast.StringNode{
					Value: "c",
				},
			},
		},
		{
			input: Operand{Values: []string{"a.b.c.d"}},
			expected: &ast.MemberNode{
				Node: &ast.MemberNode{
					Node: &ast.MemberNode{
						Node: &ast.IdentifierNode{
							Value: "a",
						},
						Property: &ast.StringNode{
							Value: "b",
						},
					},
					Property: &ast.StringNode{
						Value: "c",
					},
				},
				Property: &ast.StringNode{
					Value: "d",
				},
			},
		},
		{
			input: Operand{Values: []string{"a.b.c.d.e"}},
			expected: &ast.MemberNode{
				Node: &ast.MemberNode{
					Node: &ast.MemberNode{
						Node: &ast.MemberNode{
							Node: &ast.IdentifierNode{
								Value: "a",
							},
							Property: &ast.StringNode{
								Value: "b",
							},
						},
						Property: &ast.StringNode{
							Value: "c",
						},
					},
					Property: &ast.StringNode{
						Value: "d",
					},
				},
				Property: &ast.StringNode{
					Value: "e",
				},
			},
		},
	}
	for _, test := range tests {
		result := parseVariableNode(test.input)
		assert.Equal(t, test.expected, result, "parseVariableNode(%s)", test.input.Values[0])
	}
}

func TestParseValueNode(t *testing.T) {
	tests := []struct {
		name     string
		operand  Operand
		expected ast.Node
	}{
		{
			name:     "Float operand",
			operand:  Operand{Type: "float", Values: []string{"3.14"}},
			expected: &ast.FloatNode{Value: 3.14},
		},
		{
			name:     "String operand",
			operand:  Operand{Type: "string", Values: []string{"hello"}},
			expected: &ast.StringNode{Value: "hello"},
		},
		{
			name:     "Bool operand",
			operand:  Operand{Type: "boolean", Values: []string{"true"}},
			expected: &ast.BoolNode{Value: true},
		},
		{
			name:     "Integer operand",
			operand:  Operand{Type: "int", Values: []string{"42"}},
			expected: &ast.IntegerNode{Value: 42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseValueNode(tt.operand)
			assert.Equal(t, tt.expected, result)
		})
	}
}
