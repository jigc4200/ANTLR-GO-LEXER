// Code generated from Programa.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Programa

import "github.com/antlr4-go/antlr/v4"

// BaseProgramaListener is a complete listener for a parse tree produced by ProgramaParser.
type BaseProgramaListener struct{}

var _ ProgramaListener = &BaseProgramaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProgramaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProgramaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProgramaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProgramaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseProgramaListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseProgramaListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterDeclaraciones is called when production declaraciones is entered.
func (s *BaseProgramaListener) EnterDeclaraciones(ctx *DeclaracionesContext) {}

// ExitDeclaraciones is called when production declaraciones is exited.
func (s *BaseProgramaListener) ExitDeclaraciones(ctx *DeclaracionesContext) {}

// EnterLista_variables is called when production lista_variables is entered.
func (s *BaseProgramaListener) EnterLista_variables(ctx *Lista_variablesContext) {}

// ExitLista_variables is called when production lista_variables is exited.
func (s *BaseProgramaListener) ExitLista_variables(ctx *Lista_variablesContext) {}

// EnterBloque is called when production bloque is entered.
func (s *BaseProgramaListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BaseProgramaListener) ExitBloque(ctx *BloqueContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseProgramaListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseProgramaListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseProgramaListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseProgramaListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BaseProgramaListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BaseProgramaListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterCondicional is called when production condicional is entered.
func (s *BaseProgramaListener) EnterCondicional(ctx *CondicionalContext) {}

// ExitCondicional is called when production condicional is exited.
func (s *BaseProgramaListener) ExitCondicional(ctx *CondicionalContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BaseProgramaListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BaseProgramaListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterTermino is called when production termino is entered.
func (s *BaseProgramaListener) EnterTermino(ctx *TerminoContext) {}

// ExitTermino is called when production termino is exited.
func (s *BaseProgramaListener) ExitTermino(ctx *TerminoContext) {}

// EnterAsignacionvar is called when production asignacionvar is entered.
func (s *BaseProgramaListener) EnterAsignacionvar(ctx *AsignacionvarContext) {}

// ExitAsignacionvar is called when production asignacionvar is exited.
func (s *BaseProgramaListener) ExitAsignacionvar(ctx *AsignacionvarContext) {}

// EnterComparador is called when production comparador is entered.
func (s *BaseProgramaListener) EnterComparador(ctx *ComparadorContext) {}

// ExitComparador is called when production comparador is exited.
func (s *BaseProgramaListener) ExitComparador(ctx *ComparadorContext) {}

// EnterOperador is called when production operador is entered.
func (s *BaseProgramaListener) EnterOperador(ctx *OperadorContext) {}

// ExitOperador is called when production operador is exited.
func (s *BaseProgramaListener) ExitOperador(ctx *OperadorContext) {}
