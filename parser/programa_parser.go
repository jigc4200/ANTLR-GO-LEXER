// Code generated from Programa.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Programa

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ProgramaParser struct {
	*antlr.BaseParser
}

var ProgramaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func programaParserInit() {
	staticData := &ProgramaParserStaticData
	staticData.LiteralNames = []string{
		"", "'PROGRAMA'", "'FIN'", "'INICIO'", "'SI'", "'ENTONCES'", "'LEER'",
		"'ESCRIBIR'", "'VARIABLES'", "'<--'", "'='", "';'", "','", "'.'", "':'",
		"'['", "']'", "", "", "", "", "", "'+'", "'-'", "'*'", "'/'", "'('",
		"')'", "'=='", "'>'", "'<'", "'>='", "'<='", "'!='", "'%'", "'^'", "'++'",
		"'--'",
	}
	staticData.SymbolicNames = []string{
		"", "Reservada_PROGRAMA", "Reservada_FIN", "Reservada_INICIO", "Reservada_SI",
		"Reservada_ENTONCES", "Reservada_LEER", "Reservada_ESCRIBIR", "Reservada_VARIABLES",
		"ASIGNACION", "IGUALDAD", "PUNTOYCOMA", "COMA", "PUNTO", "DOS_PUNTOS",
		"CORCHETE_IZQ", "CORCHETE_DER", "CADENA", "CADENA_CARACTER", "BOOLEANO",
		"NUMERO_ENTERO", "IDENTIFICADOR_VARIABLE", "OPERADOR_SUMA", "OPERADOR_RESTA",
		"OPERADOR_MULTIPLICACION", "OPERADOR_DIVISION", "OPERADOR_PARENTESIS_IZQ",
		"OPERADOR_PARENTESIS_DER", "COMPARADOR_IGUAL", "COMPARADOR_MAYOR", "COMPARADOR_MENOR",
		"COMPARADOR_MAYOR_IGUAL", "COMPARADOR_MENOR_IGUAL", "COMPARADOR_DIFERENTE",
		"OPERADOR_MODULO", "OPERADOR_POTENCIA", "OPERADOR_INCREMENTO", "OPERADOR_DECREMENTO",
		"WS", "COMMENT", "COMMENT_MULTILINE",
	}
	staticData.RuleNames = []string{
		"programa", "declaraciones", "lista_variables", "bloque", "instrucciones",
		"instruccion", "asignacion", "condicional", "expresion", "termino",
		"asignacionvar", "comparador", "operador",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 40, 100, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 42, 8, 2, 10,
		2, 12, 2, 45, 9, 2, 1, 3, 1, 3, 1, 3, 1, 4, 5, 4, 51, 8, 4, 10, 4, 12,
		4, 54, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 3, 5, 67, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 84, 8, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 3, 9, 92, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 0, 0, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0,
		3, 1, 0, 9, 10, 2, 0, 10, 10, 29, 33, 1, 0, 22, 25, 94, 0, 26, 1, 0, 0,
		0, 2, 34, 1, 0, 0, 0, 4, 38, 1, 0, 0, 0, 6, 46, 1, 0, 0, 0, 8, 52, 1, 0,
		0, 0, 10, 66, 1, 0, 0, 0, 12, 68, 1, 0, 0, 0, 14, 72, 1, 0, 0, 0, 16, 79,
		1, 0, 0, 0, 18, 91, 1, 0, 0, 0, 20, 93, 1, 0, 0, 0, 22, 95, 1, 0, 0, 0,
		24, 97, 1, 0, 0, 0, 26, 27, 5, 1, 0, 0, 27, 28, 5, 21, 0, 0, 28, 29, 5,
		11, 0, 0, 29, 30, 3, 2, 1, 0, 30, 31, 3, 6, 3, 0, 31, 32, 5, 2, 0, 0, 32,
		33, 5, 11, 0, 0, 33, 1, 1, 0, 0, 0, 34, 35, 5, 8, 0, 0, 35, 36, 3, 4, 2,
		0, 36, 37, 5, 11, 0, 0, 37, 3, 1, 0, 0, 0, 38, 43, 5, 21, 0, 0, 39, 40,
		5, 12, 0, 0, 40, 42, 5, 21, 0, 0, 41, 39, 1, 0, 0, 0, 42, 45, 1, 0, 0,
		0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 5, 1, 0, 0, 0, 45, 43, 1,
		0, 0, 0, 46, 47, 5, 3, 0, 0, 47, 48, 3, 8, 4, 0, 48, 7, 1, 0, 0, 0, 49,
		51, 3, 10, 5, 0, 50, 49, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0,
		0, 0, 52, 53, 1, 0, 0, 0, 53, 9, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 56,
		3, 12, 6, 0, 56, 57, 5, 11, 0, 0, 57, 67, 1, 0, 0, 0, 58, 59, 5, 6, 0,
		0, 59, 60, 5, 21, 0, 0, 60, 67, 5, 11, 0, 0, 61, 62, 5, 7, 0, 0, 62, 63,
		3, 16, 8, 0, 63, 64, 5, 11, 0, 0, 64, 67, 1, 0, 0, 0, 65, 67, 3, 14, 7,
		0, 66, 55, 1, 0, 0, 0, 66, 58, 1, 0, 0, 0, 66, 61, 1, 0, 0, 0, 66, 65,
		1, 0, 0, 0, 67, 11, 1, 0, 0, 0, 68, 69, 5, 21, 0, 0, 69, 70, 3, 20, 10,
		0, 70, 71, 3, 16, 8, 0, 71, 13, 1, 0, 0, 0, 72, 73, 5, 4, 0, 0, 73, 74,
		3, 16, 8, 0, 74, 75, 3, 22, 11, 0, 75, 76, 3, 16, 8, 0, 76, 77, 5, 5, 0,
		0, 77, 78, 3, 8, 4, 0, 78, 15, 1, 0, 0, 0, 79, 83, 3, 18, 9, 0, 80, 81,
		3, 24, 12, 0, 81, 82, 3, 16, 8, 0, 82, 84, 1, 0, 0, 0, 83, 80, 1, 0, 0,
		0, 83, 84, 1, 0, 0, 0, 84, 17, 1, 0, 0, 0, 85, 92, 5, 20, 0, 0, 86, 92,
		5, 21, 0, 0, 87, 88, 5, 26, 0, 0, 88, 89, 3, 16, 8, 0, 89, 90, 5, 27, 0,
		0, 90, 92, 1, 0, 0, 0, 91, 85, 1, 0, 0, 0, 91, 86, 1, 0, 0, 0, 91, 87,
		1, 0, 0, 0, 92, 19, 1, 0, 0, 0, 93, 94, 7, 0, 0, 0, 94, 21, 1, 0, 0, 0,
		95, 96, 7, 1, 0, 0, 96, 23, 1, 0, 0, 0, 97, 98, 7, 2, 0, 0, 98, 25, 1,
		0, 0, 0, 5, 43, 52, 66, 83, 91,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ProgramaParserInit initializes any static state used to implement ProgramaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProgramaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ProgramaParserInit() {
	staticData := &ProgramaParserStaticData
	staticData.once.Do(programaParserInit)
}

// NewProgramaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewProgramaParser(input antlr.TokenStream) *ProgramaParser {
	ProgramaParserInit()
	this := new(ProgramaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ProgramaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Programa.g4"

	return this
}

// ProgramaParser tokens.
const (
	ProgramaParserEOF                     = antlr.TokenEOF
	ProgramaParserReservada_PROGRAMA      = 1
	ProgramaParserReservada_FIN           = 2
	ProgramaParserReservada_INICIO        = 3
	ProgramaParserReservada_SI            = 4
	ProgramaParserReservada_ENTONCES      = 5
	ProgramaParserReservada_LEER          = 6
	ProgramaParserReservada_ESCRIBIR      = 7
	ProgramaParserReservada_VARIABLES     = 8
	ProgramaParserASIGNACION              = 9
	ProgramaParserIGUALDAD                = 10
	ProgramaParserPUNTOYCOMA              = 11
	ProgramaParserCOMA                    = 12
	ProgramaParserPUNTO                   = 13
	ProgramaParserDOS_PUNTOS              = 14
	ProgramaParserCORCHETE_IZQ            = 15
	ProgramaParserCORCHETE_DER            = 16
	ProgramaParserCADENA                  = 17
	ProgramaParserCADENA_CARACTER         = 18
	ProgramaParserBOOLEANO                = 19
	ProgramaParserNUMERO_ENTERO           = 20
	ProgramaParserIDENTIFICADOR_VARIABLE  = 21
	ProgramaParserOPERADOR_SUMA           = 22
	ProgramaParserOPERADOR_RESTA          = 23
	ProgramaParserOPERADOR_MULTIPLICACION = 24
	ProgramaParserOPERADOR_DIVISION       = 25
	ProgramaParserOPERADOR_PARENTESIS_IZQ = 26
	ProgramaParserOPERADOR_PARENTESIS_DER = 27
	ProgramaParserCOMPARADOR_IGUAL        = 28
	ProgramaParserCOMPARADOR_MAYOR        = 29
	ProgramaParserCOMPARADOR_MENOR        = 30
	ProgramaParserCOMPARADOR_MAYOR_IGUAL  = 31
	ProgramaParserCOMPARADOR_MENOR_IGUAL  = 32
	ProgramaParserCOMPARADOR_DIFERENTE    = 33
	ProgramaParserOPERADOR_MODULO         = 34
	ProgramaParserOPERADOR_POTENCIA       = 35
	ProgramaParserOPERADOR_INCREMENTO     = 36
	ProgramaParserOPERADOR_DECREMENTO     = 37
	ProgramaParserWS                      = 38
	ProgramaParserCOMMENT                 = 39
	ProgramaParserCOMMENT_MULTILINE       = 40
)

// ProgramaParser rules.
const (
	ProgramaParserRULE_programa        = 0
	ProgramaParserRULE_declaraciones   = 1
	ProgramaParserRULE_lista_variables = 2
	ProgramaParserRULE_bloque          = 3
	ProgramaParserRULE_instrucciones   = 4
	ProgramaParserRULE_instruccion     = 5
	ProgramaParserRULE_asignacion      = 6
	ProgramaParserRULE_condicional     = 7
	ProgramaParserRULE_expresion       = 8
	ProgramaParserRULE_termino         = 9
	ProgramaParserRULE_asignacionvar   = 10
	ProgramaParserRULE_comparador      = 11
	ProgramaParserRULE_operador        = 12
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Reservada_PROGRAMA() antlr.TerminalNode
	IDENTIFICADOR_VARIABLE() antlr.TerminalNode
	AllPUNTOYCOMA() []antlr.TerminalNode
	PUNTOYCOMA(i int) antlr.TerminalNode
	Declaraciones() IDeclaracionesContext
	Bloque() IBloqueContext
	Reservada_FIN() antlr.TerminalNode

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) Reservada_PROGRAMA() antlr.TerminalNode {
	return s.GetToken(ProgramaParserReservada_PROGRAMA, 0)
}

func (s *ProgramaContext) IDENTIFICADOR_VARIABLE() antlr.TerminalNode {
	return s.GetToken(ProgramaParserIDENTIFICADOR_VARIABLE, 0)
}

func (s *ProgramaContext) AllPUNTOYCOMA() []antlr.TerminalNode {
	return s.GetTokens(ProgramaParserPUNTOYCOMA)
}

func (s *ProgramaContext) PUNTOYCOMA(i int) antlr.TerminalNode {
	return s.GetToken(ProgramaParserPUNTOYCOMA, i)
}

func (s *ProgramaContext) Declaraciones() IDeclaracionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionesContext)
}

func (s *ProgramaContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ProgramaContext) Reservada_FIN() antlr.TerminalNode {
	return s.GetToken(ProgramaParserReservada_FIN, 0)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (p *ProgramaParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ProgramaParserRULE_programa)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(ProgramaParserReservada_PROGRAMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.Match(ProgramaParserIDENTIFICADOR_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.Match(ProgramaParserPUNTOYCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Declaraciones()
	}
	{
		p.SetState(30)
		p.Bloque()
	}
	{
		p.SetState(31)
		p.Match(ProgramaParserReservada_FIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Match(ProgramaParserPUNTOYCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracionesContext is an interface to support dynamic dispatch.
type IDeclaracionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Reservada_VARIABLES() antlr.TerminalNode
	Lista_variables() ILista_variablesContext
	PUNTOYCOMA() antlr.TerminalNode

	// IsDeclaracionesContext differentiates from other interfaces.
	IsDeclaracionesContext()
}

type DeclaracionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionesContext() *DeclaracionesContext {
	var p = new(DeclaracionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_declaraciones
	return p
}

func InitEmptyDeclaracionesContext(p *DeclaracionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_declaraciones
}

func (*DeclaracionesContext) IsDeclaracionesContext() {}

func NewDeclaracionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionesContext {
	var p = new(DeclaracionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_declaraciones

	return p
}

func (s *DeclaracionesContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionesContext) Reservada_VARIABLES() antlr.TerminalNode {
	return s.GetToken(ProgramaParserReservada_VARIABLES, 0)
}

func (s *DeclaracionesContext) Lista_variables() ILista_variablesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_variablesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILista_variablesContext)
}

func (s *DeclaracionesContext) PUNTOYCOMA() antlr.TerminalNode {
	return s.GetToken(ProgramaParserPUNTOYCOMA, 0)
}

func (s *DeclaracionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterDeclaraciones(s)
	}
}

func (s *DeclaracionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitDeclaraciones(s)
	}
}

func (p *ProgramaParser) Declaraciones() (localctx IDeclaracionesContext) {
	localctx = NewDeclaracionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ProgramaParserRULE_declaraciones)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(ProgramaParserReservada_VARIABLES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
		p.Lista_variables()
	}
	{
		p.SetState(36)
		p.Match(ProgramaParserPUNTOYCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILista_variablesContext is an interface to support dynamic dispatch.
type ILista_variablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFICADOR_VARIABLE() []antlr.TerminalNode
	IDENTIFICADOR_VARIABLE(i int) antlr.TerminalNode
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsLista_variablesContext differentiates from other interfaces.
	IsLista_variablesContext()
}

type Lista_variablesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_variablesContext() *Lista_variablesContext {
	var p = new(Lista_variablesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_lista_variables
	return p
}

func InitEmptyLista_variablesContext(p *Lista_variablesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_lista_variables
}

func (*Lista_variablesContext) IsLista_variablesContext() {}

func NewLista_variablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_variablesContext {
	var p = new(Lista_variablesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_lista_variables

	return p
}

func (s *Lista_variablesContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_variablesContext) AllIDENTIFICADOR_VARIABLE() []antlr.TerminalNode {
	return s.GetTokens(ProgramaParserIDENTIFICADOR_VARIABLE)
}

func (s *Lista_variablesContext) IDENTIFICADOR_VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(ProgramaParserIDENTIFICADOR_VARIABLE, i)
}

func (s *Lista_variablesContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(ProgramaParserCOMA)
}

func (s *Lista_variablesContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(ProgramaParserCOMA, i)
}

func (s *Lista_variablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_variablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lista_variablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterLista_variables(s)
	}
}

func (s *Lista_variablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitLista_variables(s)
	}
}

func (p *ProgramaParser) Lista_variables() (localctx ILista_variablesContext) {
	localctx = NewLista_variablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ProgramaParserRULE_lista_variables)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(ProgramaParserIDENTIFICADOR_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ProgramaParserCOMA {
		{
			p.SetState(39)
			p.Match(ProgramaParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Match(ProgramaParserIDENTIFICADOR_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Reservada_INICIO() antlr.TerminalNode
	Instrucciones() IInstruccionesContext

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_bloque
	return p
}

func InitEmptyBloqueContext(p *BloqueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_bloque
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) Reservada_INICIO() antlr.TerminalNode {
	return s.GetToken(ProgramaParserReservada_INICIO, 0)
}

func (s *BloqueContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (p *ProgramaParser) Bloque() (localctx IBloqueContext) {
	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ProgramaParserRULE_bloque)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(ProgramaParserReservada_INICIO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Instrucciones()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInstruccion() []IInstruccionContext
	Instruccion(i int) IInstruccionContext

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_instrucciones
	return p
}

func InitEmptyInstruccionesContext(p *InstruccionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_instrucciones
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionContext); ok {
			tst[i] = t.(IInstruccionContext)
			i++
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *ProgramaParser) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ProgramaParserRULE_instrucciones)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(49)
				p.Instruccion()
			}

		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Asignacion() IAsignacionContext
	PUNTOYCOMA() antlr.TerminalNode
	Reservada_LEER() antlr.TerminalNode
	IDENTIFICADOR_VARIABLE() antlr.TerminalNode
	Reservada_ESCRIBIR() antlr.TerminalNode
	Expresion() IExpresionContext
	Condicional() ICondicionalContext

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_instruccion
	return p
}

func InitEmptyInstruccionContext(p *InstruccionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_instruccion
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *InstruccionContext) PUNTOYCOMA() antlr.TerminalNode {
	return s.GetToken(ProgramaParserPUNTOYCOMA, 0)
}

func (s *InstruccionContext) Reservada_LEER() antlr.TerminalNode {
	return s.GetToken(ProgramaParserReservada_LEER, 0)
}

func (s *InstruccionContext) IDENTIFICADOR_VARIABLE() antlr.TerminalNode {
	return s.GetToken(ProgramaParserIDENTIFICADOR_VARIABLE, 0)
}

func (s *InstruccionContext) Reservada_ESCRIBIR() antlr.TerminalNode {
	return s.GetToken(ProgramaParserReservada_ESCRIBIR, 0)
}

func (s *InstruccionContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *InstruccionContext) Condicional() ICondicionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondicionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondicionalContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *ProgramaParser) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ProgramaParserRULE_instruccion)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ProgramaParserIDENTIFICADOR_VARIABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Asignacion()
		}
		{
			p.SetState(56)
			p.Match(ProgramaParserPUNTOYCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ProgramaParserReservada_LEER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Match(ProgramaParserReservada_LEER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Match(ProgramaParserIDENTIFICADOR_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Match(ProgramaParserPUNTOYCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ProgramaParserReservada_ESCRIBIR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.Match(ProgramaParserReservada_ESCRIBIR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Expresion()
		}
		{
			p.SetState(63)
			p.Match(ProgramaParserPUNTOYCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ProgramaParserReservada_SI:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)
			p.Condicional()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFICADOR_VARIABLE() antlr.TerminalNode
	Asignacionvar() IAsignacionvarContext
	Expresion() IExpresionContext

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_asignacion
	return p
}

func InitEmptyAsignacionContext(p *AsignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_asignacion
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) IDENTIFICADOR_VARIABLE() antlr.TerminalNode {
	return s.GetToken(ProgramaParserIDENTIFICADOR_VARIABLE, 0)
}

func (s *AsignacionContext) Asignacionvar() IAsignacionvarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionvarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionvarContext)
}

func (s *AsignacionContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (p *ProgramaParser) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ProgramaParserRULE_asignacion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(ProgramaParserIDENTIFICADOR_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Asignacionvar()
	}
	{
		p.SetState(70)
		p.Expresion()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICondicionalContext is an interface to support dynamic dispatch.
type ICondicionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Reservada_SI() antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	Comparador() IComparadorContext
	Reservada_ENTONCES() antlr.TerminalNode
	Instrucciones() IInstruccionesContext

	// IsCondicionalContext differentiates from other interfaces.
	IsCondicionalContext()
}

type CondicionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondicionalContext() *CondicionalContext {
	var p = new(CondicionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_condicional
	return p
}

func InitEmptyCondicionalContext(p *CondicionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_condicional
}

func (*CondicionalContext) IsCondicionalContext() {}

func NewCondicionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondicionalContext {
	var p = new(CondicionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_condicional

	return p
}

func (s *CondicionalContext) GetParser() antlr.Parser { return s.parser }

func (s *CondicionalContext) Reservada_SI() antlr.TerminalNode {
	return s.GetToken(ProgramaParserReservada_SI, 0)
}

func (s *CondicionalContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *CondicionalContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CondicionalContext) Comparador() IComparadorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparadorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparadorContext)
}

func (s *CondicionalContext) Reservada_ENTONCES() antlr.TerminalNode {
	return s.GetToken(ProgramaParserReservada_ENTONCES, 0)
}

func (s *CondicionalContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *CondicionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondicionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondicionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterCondicional(s)
	}
}

func (s *CondicionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitCondicional(s)
	}
}

func (p *ProgramaParser) Condicional() (localctx ICondicionalContext) {
	localctx = NewCondicionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ProgramaParserRULE_condicional)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(ProgramaParserReservada_SI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Expresion()
	}
	{
		p.SetState(74)
		p.Comparador()
	}
	{
		p.SetState(75)
		p.Expresion()
	}
	{
		p.SetState(76)
		p.Match(ProgramaParserReservada_ENTONCES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Instrucciones()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Termino() ITerminoContext
	Operador() IOperadorContext
	Expresion() IExpresionContext

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) Termino() ITerminoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminoContext)
}

func (s *ExpresionContext) Operador() IOperadorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperadorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperadorContext)
}

func (s *ExpresionContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *ProgramaParser) Expresion() (localctx IExpresionContext) {
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ProgramaParserRULE_expresion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Termino()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62914560) != 0 {
		{
			p.SetState(80)
			p.Operador()
		}
		{
			p.SetState(81)
			p.Expresion()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITerminoContext is an interface to support dynamic dispatch.
type ITerminoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERO_ENTERO() antlr.TerminalNode
	IDENTIFICADOR_VARIABLE() antlr.TerminalNode
	OPERADOR_PARENTESIS_IZQ() antlr.TerminalNode
	Expresion() IExpresionContext
	OPERADOR_PARENTESIS_DER() antlr.TerminalNode

	// IsTerminoContext differentiates from other interfaces.
	IsTerminoContext()
}

type TerminoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminoContext() *TerminoContext {
	var p = new(TerminoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_termino
	return p
}

func InitEmptyTerminoContext(p *TerminoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_termino
}

func (*TerminoContext) IsTerminoContext() {}

func NewTerminoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerminoContext {
	var p = new(TerminoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_termino

	return p
}

func (s *TerminoContext) GetParser() antlr.Parser { return s.parser }

func (s *TerminoContext) NUMERO_ENTERO() antlr.TerminalNode {
	return s.GetToken(ProgramaParserNUMERO_ENTERO, 0)
}

func (s *TerminoContext) IDENTIFICADOR_VARIABLE() antlr.TerminalNode {
	return s.GetToken(ProgramaParserIDENTIFICADOR_VARIABLE, 0)
}

func (s *TerminoContext) OPERADOR_PARENTESIS_IZQ() antlr.TerminalNode {
	return s.GetToken(ProgramaParserOPERADOR_PARENTESIS_IZQ, 0)
}

func (s *TerminoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *TerminoContext) OPERADOR_PARENTESIS_DER() antlr.TerminalNode {
	return s.GetToken(ProgramaParserOPERADOR_PARENTESIS_DER, 0)
}

func (s *TerminoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerminoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterTermino(s)
	}
}

func (s *TerminoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitTermino(s)
	}
}

func (p *ProgramaParser) Termino() (localctx ITerminoContext) {
	localctx = NewTerminoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ProgramaParserRULE_termino)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ProgramaParserNUMERO_ENTERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.Match(ProgramaParserNUMERO_ENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ProgramaParserIDENTIFICADOR_VARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.Match(ProgramaParserIDENTIFICADOR_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ProgramaParserOPERADOR_PARENTESIS_IZQ:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.Match(ProgramaParserOPERADOR_PARENTESIS_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Expresion()
		}
		{
			p.SetState(89)
			p.Match(ProgramaParserOPERADOR_PARENTESIS_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionvarContext is an interface to support dynamic dispatch.
type IAsignacionvarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASIGNACION() antlr.TerminalNode
	IGUALDAD() antlr.TerminalNode

	// IsAsignacionvarContext differentiates from other interfaces.
	IsAsignacionvarContext()
}

type AsignacionvarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionvarContext() *AsignacionvarContext {
	var p = new(AsignacionvarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_asignacionvar
	return p
}

func InitEmptyAsignacionvarContext(p *AsignacionvarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_asignacionvar
}

func (*AsignacionvarContext) IsAsignacionvarContext() {}

func NewAsignacionvarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionvarContext {
	var p = new(AsignacionvarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_asignacionvar

	return p
}

func (s *AsignacionvarContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionvarContext) ASIGNACION() antlr.TerminalNode {
	return s.GetToken(ProgramaParserASIGNACION, 0)
}

func (s *AsignacionvarContext) IGUALDAD() antlr.TerminalNode {
	return s.GetToken(ProgramaParserIGUALDAD, 0)
}

func (s *AsignacionvarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionvarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionvarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterAsignacionvar(s)
	}
}

func (s *AsignacionvarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitAsignacionvar(s)
	}
}

func (p *ProgramaParser) Asignacionvar() (localctx IAsignacionvarContext) {
	localctx = NewAsignacionvarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ProgramaParserRULE_asignacionvar)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ProgramaParserASIGNACION || _la == ProgramaParserIGUALDAD) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparadorContext is an interface to support dynamic dispatch.
type IComparadorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IGUALDAD() antlr.TerminalNode
	COMPARADOR_MENOR() antlr.TerminalNode
	COMPARADOR_MAYOR() antlr.TerminalNode
	COMPARADOR_MENOR_IGUAL() antlr.TerminalNode
	COMPARADOR_MAYOR_IGUAL() antlr.TerminalNode
	COMPARADOR_DIFERENTE() antlr.TerminalNode

	// IsComparadorContext differentiates from other interfaces.
	IsComparadorContext()
}

type ComparadorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparadorContext() *ComparadorContext {
	var p = new(ComparadorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_comparador
	return p
}

func InitEmptyComparadorContext(p *ComparadorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_comparador
}

func (*ComparadorContext) IsComparadorContext() {}

func NewComparadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparadorContext {
	var p = new(ComparadorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_comparador

	return p
}

func (s *ComparadorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparadorContext) IGUALDAD() antlr.TerminalNode {
	return s.GetToken(ProgramaParserIGUALDAD, 0)
}

func (s *ComparadorContext) COMPARADOR_MENOR() antlr.TerminalNode {
	return s.GetToken(ProgramaParserCOMPARADOR_MENOR, 0)
}

func (s *ComparadorContext) COMPARADOR_MAYOR() antlr.TerminalNode {
	return s.GetToken(ProgramaParserCOMPARADOR_MAYOR, 0)
}

func (s *ComparadorContext) COMPARADOR_MENOR_IGUAL() antlr.TerminalNode {
	return s.GetToken(ProgramaParserCOMPARADOR_MENOR_IGUAL, 0)
}

func (s *ComparadorContext) COMPARADOR_MAYOR_IGUAL() antlr.TerminalNode {
	return s.GetToken(ProgramaParserCOMPARADOR_MAYOR_IGUAL, 0)
}

func (s *ComparadorContext) COMPARADOR_DIFERENTE() antlr.TerminalNode {
	return s.GetToken(ProgramaParserCOMPARADOR_DIFERENTE, 0)
}

func (s *ComparadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparadorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparadorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterComparador(s)
	}
}

func (s *ComparadorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitComparador(s)
	}
}

func (p *ProgramaParser) Comparador() (localctx IComparadorContext) {
	localctx = NewComparadorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ProgramaParserRULE_comparador)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16642999296) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperadorContext is an interface to support dynamic dispatch.
type IOperadorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPERADOR_SUMA() antlr.TerminalNode
	OPERADOR_RESTA() antlr.TerminalNode
	OPERADOR_MULTIPLICACION() antlr.TerminalNode
	OPERADOR_DIVISION() antlr.TerminalNode

	// IsOperadorContext differentiates from other interfaces.
	IsOperadorContext()
}

type OperadorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperadorContext() *OperadorContext {
	var p = new(OperadorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_operador
	return p
}

func InitEmptyOperadorContext(p *OperadorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ProgramaParserRULE_operador
}

func (*OperadorContext) IsOperadorContext() {}

func NewOperadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperadorContext {
	var p = new(OperadorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProgramaParserRULE_operador

	return p
}

func (s *OperadorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperadorContext) OPERADOR_SUMA() antlr.TerminalNode {
	return s.GetToken(ProgramaParserOPERADOR_SUMA, 0)
}

func (s *OperadorContext) OPERADOR_RESTA() antlr.TerminalNode {
	return s.GetToken(ProgramaParserOPERADOR_RESTA, 0)
}

func (s *OperadorContext) OPERADOR_MULTIPLICACION() antlr.TerminalNode {
	return s.GetToken(ProgramaParserOPERADOR_MULTIPLICACION, 0)
}

func (s *OperadorContext) OPERADOR_DIVISION() antlr.TerminalNode {
	return s.GetToken(ProgramaParserOPERADOR_DIVISION, 0)
}

func (s *OperadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperadorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperadorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.EnterOperador(s)
	}
}

func (s *OperadorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProgramaListener); ok {
		listenerT.ExitOperador(s)
	}
}

func (p *ProgramaParser) Operador() (localctx IOperadorContext) {
	localctx = NewOperadorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ProgramaParserRULE_operador)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62914560) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
