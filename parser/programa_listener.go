// Code generated from Programa.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Programa

import "github.com/antlr4-go/antlr/v4"

// ProgramaListener is a complete listener for a parse tree produced by ProgramaParser.
type ProgramaListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterDeclaraciones is called when entering the declaraciones production.
	EnterDeclaraciones(c *DeclaracionesContext)

	// EnterLista_variables is called when entering the lista_variables production.
	EnterLista_variables(c *Lista_variablesContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterCondicional is called when entering the condicional production.
	EnterCondicional(c *CondicionalContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterTermino is called when entering the termino production.
	EnterTermino(c *TerminoContext)

	// EnterAsignacionvar is called when entering the asignacionvar production.
	EnterAsignacionvar(c *AsignacionvarContext)

	// EnterComparador is called when entering the comparador production.
	EnterComparador(c *ComparadorContext)

	// EnterOperador is called when entering the operador production.
	EnterOperador(c *OperadorContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitDeclaraciones is called when exiting the declaraciones production.
	ExitDeclaraciones(c *DeclaracionesContext)

	// ExitLista_variables is called when exiting the lista_variables production.
	ExitLista_variables(c *Lista_variablesContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitCondicional is called when exiting the condicional production.
	ExitCondicional(c *CondicionalContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitTermino is called when exiting the termino production.
	ExitTermino(c *TerminoContext)

	// ExitAsignacionvar is called when exiting the asignacionvar production.
	ExitAsignacionvar(c *AsignacionvarContext)

	// ExitComparador is called when exiting the comparador production.
	ExitComparador(c *ComparadorContext)

	// ExitOperador is called when exiting the operador production.
	ExitOperador(c *OperadorContext)
}
