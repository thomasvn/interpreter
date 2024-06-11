package ast

import "main/token"

// Programs in MonkeyLanguage are just a series of Statements. Statements
// consist of Identifiers and Expressions.
//
// Example:
// let <identifier> = <expression>;

// -----------------------------------------------------------------------------
// Interface for all AST Nodes

type Node interface {
	TokenLiteral() string
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

// -----------------------------------------------------------------------------
// LetStatement Node

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// -----------------------------------------------------------------------------
// Identifier Node

type Identifier struct {
	Token token.Token // the token.IDENT token Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
