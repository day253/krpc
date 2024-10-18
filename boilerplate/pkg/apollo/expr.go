package apollo

import (
	"strconv"
	"strings"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/expr-lang/expr/ast"
)

func parseOperand(op *Operand) *ast.BinaryNode {
	leafNode := func() bool {
		for _, operand := range op.Operands {
			if operand.Operator == "var" || operand.Operator == "val" {
				return true
			}
		}
		return false
	}()
	if leafNode && len(op.Operands) >= 2 {
		return &ast.BinaryNode{
			Operator: parseOperator(op.Operator),
			Left:     parseLeafNode(op.Operands[0]),
			Right:    parseLeafNode(op.Operands[1]),
		}
	}
	left := parseOperand(&op.Operands[0])
	for i := 1; i < len(op.Operands); i++ {
		right := parseOperand(&op.Operands[i])
		left = &ast.BinaryNode{
			Operator: parseOperator(op.Operator),
			Left:     left,
			Right:    right,
		}
	}
	return left
}

func parseOperator(operator string) string {
	operatorMap := map[string]string{
		"!=":      "!=",
		"<":       "<",
		"<=":      "<=",
		"=":       "==",
		">":       ">",
		">=":      ">=",
		"and":     "&&",
		"not":     "!",
		"or":      "||",
		"subsetp": "in",
		"contain": "contains",
		"eq":      "==",
		"neq":     "!=",
	}
	if mappedOperator, exists := operatorMap[operator]; exists {
		return mappedOperator
	}
	klog.Fatal("unsupported operator: ", operator)
	return operator
}

func parseLeafNode(operand Operand) ast.Node {
	if operand.Operator == "var" {
		return parseVariableNode(operand)
	} else if operand.Operator == "val" {
		return parseValueNode(operand)
	}
	return nil
}

func parseVariableNode(operand Operand) ast.Node {
	parts := strings.Split(operand.Values[0], ".")
	var node ast.Node = &ast.IdentifierNode{Value: parts[0]}
	for _, part := range parts[1:] {
		node = &ast.MemberNode{
			Node: node,
			Property: &ast.StringNode{
				Value: part,
			},
		}
	}
	return node
}

func parseValueNode(operand Operand) ast.Node {
	switch operand.Type {
	case "array_boolean":
		boolValues := make([]ast.Node, len(operand.Values))
		for i, v := range operand.Values {
			value, _ := strconv.ParseBool(v)
			boolValues[i] = &ast.BoolNode{Value: value}
		}
		return &ast.ArrayNode{Nodes: boolValues}
	case "array_int":
		intValues := make([]ast.Node, len(operand.Values))
		for i, v := range operand.Values {
			value, _ := strconv.Atoi(v)
			intValues[i] = &ast.IntegerNode{Value: value}
		}
		return &ast.ArrayNode{Nodes: intValues}
	case "array_float":
		fallthrough
	case "array_number":
		numberValues := make([]ast.Node, len(operand.Values))
		for i, v := range operand.Values {
			value, _ := strconv.ParseFloat(v, 64)
			numberValues[i] = &ast.FloatNode{Value: value}
		}
		return &ast.ArrayNode{Nodes: numberValues}
	case "array":
		fallthrough
	case "array_string":
		stringValues := make([]ast.Node, len(operand.Values))
		for i, v := range operand.Values {
			stringValues[i] = &ast.StringNode{Value: v}
		}
		return &ast.ArrayNode{Nodes: stringValues}
	case "boolean":
		value, _ := strconv.ParseBool(operand.Values[0])
		return &ast.BoolNode{Value: value}
	case "int":
		value, _ := strconv.Atoi(operand.Values[0])
		return &ast.IntegerNode{Value: value}
	case "float":
		fallthrough
	case "number":
		value, _ := strconv.ParseFloat(operand.Values[0], 64)
		return &ast.FloatNode{Value: value}
	case "string":
		return &ast.StringNode{Value: operand.Values[0]}
	default:
		klog.Fatal("unsupported type: ", operand.Type)
		return nil
	}
}
