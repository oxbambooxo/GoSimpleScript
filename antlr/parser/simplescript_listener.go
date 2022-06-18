// Code generated from SimpleScript.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SimpleScriptListener is a complete listener for a parse tree produced by SimpleScriptParser.
type SimpleScriptListener interface {
	antlr.ParseTreeListener

	// EnterProgramm is called when entering the programm production.
	EnterProgramm(c *ProgrammContext)

	// EnterIntDeclaration is called when entering the intDeclaration production.
	EnterIntDeclaration(c *IntDeclarationContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAdditive is called when entering the additive production.
	EnterAdditive(c *AdditiveContext)

	// EnterMultiplicative is called when entering the multiplicative production.
	EnterMultiplicative(c *MultiplicativeContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// ExitProgramm is called when exiting the programm production.
	ExitProgramm(c *ProgrammContext)

	// ExitIntDeclaration is called when exiting the intDeclaration production.
	ExitIntDeclaration(c *IntDeclarationContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAdditive is called when exiting the additive production.
	ExitAdditive(c *AdditiveContext)

	// ExitMultiplicative is called when exiting the multiplicative production.
	ExitMultiplicative(c *MultiplicativeContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)
}
