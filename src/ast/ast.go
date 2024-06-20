package ast

import (
	"bytes"
	"main/token"
)

// Programs in MonkeyLanguage are just a series of Statements. Statements
// consist of Identifiers and Expressions.
//
// Example:
// let <identifier> = <expression>;

// -----------------------------------------------------------------------------
// Interface for all AST Nodes

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// -----------------------------------------------------------------------------
// Root Node

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// -----------------------------------------------------------------------------
// LetStatement Node
//
// Example:
//   let x = 5;
//   let y = 10;
//   let foobar = 838383;
//
// Generally:
//   let <identifier> = <expression>;

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// -----------------------------------------------------------------------------
// Identifier Node

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// -----------------------------------------------------------------------------
// ReturnStatement Node
//
// Example:
//   return 5;
//   return 10;
//   return add(15);
//
// Generally:
//   return <expression>;

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// -----------------------------------------------------------------------------
// ExpressionStatement Node
//
// Example:
//   x;
//   x + 10;
//   y;
//   add(x);
//
// Generally:
//   <expression>;

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	var out bytes.Buffer
	out.WriteString(es.Expression.String())
	out.WriteString(";")
	return out.String()
}
