// Generated from /home/jigc4200/Documents/antlr-go-example/Programa.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class ProgramaParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		Reservada_PROGRAMA=1, Reservada_FIN=2, Reservada_INICIO=3, Reservada_SI=4, 
		Reservada_ENTONCES=5, Reservada_LEER=6, Reservada_ESCRIBIR=7, Reservada_VARIABLES=8, 
		ASIGNACION=9, IGUALDAD=10, PUNTOYCOMA=11, COMA=12, PUNTO=13, DOS_PUNTOS=14, 
		CORCHETE_IZQ=15, CORCHETE_DER=16, CADENA=17, CADENA_CARACTER=18, BOOLEANO=19, 
		NUMERO_ENTERO=20, IDENTIFICADOR_VARIABLE=21, OPERADOR_SUMA=22, OPERADOR_RESTA=23, 
		OPERADOR_MULTIPLICACION=24, OPERADOR_DIVISION=25, OPERADOR_PARENTESIS_IZQ=26, 
		OPERADOR_PARENTESIS_DER=27, COMPARADOR_IGUAL=28, COMPARADOR_MAYOR=29, 
		COMPARADOR_MENOR=30, COMPARADOR_MAYOR_IGUAL=31, COMPARADOR_MENOR_IGUAL=32, 
		COMPARADOR_DIFERENTE=33, OPERADOR_MODULO=34, OPERADOR_POTENCIA=35, OPERADOR_INCREMENTO=36, 
		OPERADOR_DECREMENTO=37, WS=38, COMMENT=39;
	public static final int
		RULE_programa = 0, RULE_declaraciones = 1, RULE_lista_variables = 2, RULE_bloque = 3, 
		RULE_instrucciones = 4, RULE_instruccion = 5, RULE_asignacion = 6, RULE_condicional = 7, 
		RULE_expresion = 8, RULE_termino = 9, RULE_asignacionvar = 10, RULE_comparador = 11, 
		RULE_operador = 12;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "declaraciones", "lista_variables", "bloque", "instrucciones", 
			"instruccion", "asignacion", "condicional", "expresion", "termino", "asignacionvar", 
			"comparador", "operador"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'PROGRAMA'", "'FIN'", "'INICIO'", "'SI'", "'ENTONCES'", "'LEER'", 
			"'ESCRIBIR'", "'VARIABLES'", "'<--'", "'='", "';'", "','", "'.'", "':'", 
			"'['", "']'", null, null, null, null, null, "'+'", "'-'", "'*'", "'/'", 
			"'('", "')'", "'=='", "'>'", "'<'", "'>='", "'<='", "'!='", "'%'", "'^'", 
			"'++'", "'--'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "Reservada_PROGRAMA", "Reservada_FIN", "Reservada_INICIO", "Reservada_SI", 
			"Reservada_ENTONCES", "Reservada_LEER", "Reservada_ESCRIBIR", "Reservada_VARIABLES", 
			"ASIGNACION", "IGUALDAD", "PUNTOYCOMA", "COMA", "PUNTO", "DOS_PUNTOS", 
			"CORCHETE_IZQ", "CORCHETE_DER", "CADENA", "CADENA_CARACTER", "BOOLEANO", 
			"NUMERO_ENTERO", "IDENTIFICADOR_VARIABLE", "OPERADOR_SUMA", "OPERADOR_RESTA", 
			"OPERADOR_MULTIPLICACION", "OPERADOR_DIVISION", "OPERADOR_PARENTESIS_IZQ", 
			"OPERADOR_PARENTESIS_DER", "COMPARADOR_IGUAL", "COMPARADOR_MAYOR", "COMPARADOR_MENOR", 
			"COMPARADOR_MAYOR_IGUAL", "COMPARADOR_MENOR_IGUAL", "COMPARADOR_DIFERENTE", 
			"OPERADOR_MODULO", "OPERADOR_POTENCIA", "OPERADOR_INCREMENTO", "OPERADOR_DECREMENTO", 
			"WS", "COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Programa.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public ProgramaParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramaContext extends ParserRuleContext {
		public TerminalNode Reservada_PROGRAMA() { return getToken(ProgramaParser.Reservada_PROGRAMA, 0); }
		public TerminalNode IDENTIFICADOR_VARIABLE() { return getToken(ProgramaParser.IDENTIFICADOR_VARIABLE, 0); }
		public List<TerminalNode> PUNTOYCOMA() { return getTokens(ProgramaParser.PUNTOYCOMA); }
		public TerminalNode PUNTOYCOMA(int i) {
			return getToken(ProgramaParser.PUNTOYCOMA, i);
		}
		public DeclaracionesContext declaraciones() {
			return getRuleContext(DeclaracionesContext.class,0);
		}
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public TerminalNode Reservada_FIN() { return getToken(ProgramaParser.Reservada_FIN, 0); }
		public ProgramaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_programa; }
	}

	public final ProgramaContext programa() throws RecognitionException {
		ProgramaContext _localctx = new ProgramaContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_programa);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(26);
			match(Reservada_PROGRAMA);
			setState(27);
			match(IDENTIFICADOR_VARIABLE);
			setState(28);
			match(PUNTOYCOMA);
			setState(29);
			declaraciones();
			setState(30);
			bloque();
			setState(31);
			match(Reservada_FIN);
			setState(32);
			match(PUNTOYCOMA);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DeclaracionesContext extends ParserRuleContext {
		public TerminalNode Reservada_VARIABLES() { return getToken(ProgramaParser.Reservada_VARIABLES, 0); }
		public Lista_variablesContext lista_variables() {
			return getRuleContext(Lista_variablesContext.class,0);
		}
		public TerminalNode PUNTOYCOMA() { return getToken(ProgramaParser.PUNTOYCOMA, 0); }
		public DeclaracionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaraciones; }
	}

	public final DeclaracionesContext declaraciones() throws RecognitionException {
		DeclaracionesContext _localctx = new DeclaracionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_declaraciones);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(34);
			match(Reservada_VARIABLES);
			setState(35);
			lista_variables();
			setState(36);
			match(PUNTOYCOMA);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Lista_variablesContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFICADOR_VARIABLE() { return getTokens(ProgramaParser.IDENTIFICADOR_VARIABLE); }
		public TerminalNode IDENTIFICADOR_VARIABLE(int i) {
			return getToken(ProgramaParser.IDENTIFICADOR_VARIABLE, i);
		}
		public List<TerminalNode> COMA() { return getTokens(ProgramaParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(ProgramaParser.COMA, i);
		}
		public Lista_variablesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista_variables; }
	}

	public final Lista_variablesContext lista_variables() throws RecognitionException {
		Lista_variablesContext _localctx = new Lista_variablesContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_lista_variables);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(38);
			match(IDENTIFICADOR_VARIABLE);
			setState(43);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(39);
				match(COMA);
				setState(40);
				match(IDENTIFICADOR_VARIABLE);
				}
				}
				setState(45);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BloqueContext extends ParserRuleContext {
		public TerminalNode Reservada_INICIO() { return getToken(ProgramaParser.Reservada_INICIO, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public BloqueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bloque; }
	}

	public final BloqueContext bloque() throws RecognitionException {
		BloqueContext _localctx = new BloqueContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_bloque);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			match(Reservada_INICIO);
			setState(47);
			instrucciones();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstruccionesContext extends ParserRuleContext {
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_instrucciones);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(52);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(49);
					instruccion();
					}
					} 
				}
				setState(54);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstruccionContext extends ParserRuleContext {
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
		}
		public TerminalNode PUNTOYCOMA() { return getToken(ProgramaParser.PUNTOYCOMA, 0); }
		public TerminalNode Reservada_LEER() { return getToken(ProgramaParser.Reservada_LEER, 0); }
		public TerminalNode IDENTIFICADOR_VARIABLE() { return getToken(ProgramaParser.IDENTIFICADOR_VARIABLE, 0); }
		public TerminalNode Reservada_ESCRIBIR() { return getToken(ProgramaParser.Reservada_ESCRIBIR, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public CondicionalContext condicional() {
			return getRuleContext(CondicionalContext.class,0);
		}
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_instruccion);
		try {
			setState(66);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFICADOR_VARIABLE:
				enterOuterAlt(_localctx, 1);
				{
				setState(55);
				asignacion();
				setState(56);
				match(PUNTOYCOMA);
				}
				break;
			case Reservada_LEER:
				enterOuterAlt(_localctx, 2);
				{
				setState(58);
				match(Reservada_LEER);
				setState(59);
				match(IDENTIFICADOR_VARIABLE);
				setState(60);
				match(PUNTOYCOMA);
				}
				break;
			case Reservada_ESCRIBIR:
				enterOuterAlt(_localctx, 3);
				{
				setState(61);
				match(Reservada_ESCRIBIR);
				setState(62);
				expresion();
				setState(63);
				match(PUNTOYCOMA);
				}
				break;
			case Reservada_SI:
				enterOuterAlt(_localctx, 4);
				{
				setState(65);
				condicional();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionContext extends ParserRuleContext {
		public TerminalNode IDENTIFICADOR_VARIABLE() { return getToken(ProgramaParser.IDENTIFICADOR_VARIABLE, 0); }
		public AsignacionvarContext asignacionvar() {
			return getRuleContext(AsignacionvarContext.class,0);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_asignacion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			match(IDENTIFICADOR_VARIABLE);
			setState(69);
			asignacionvar();
			setState(70);
			expresion();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CondicionalContext extends ParserRuleContext {
		public TerminalNode Reservada_SI() { return getToken(ProgramaParser.Reservada_SI, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public ComparadorContext comparador() {
			return getRuleContext(ComparadorContext.class,0);
		}
		public TerminalNode Reservada_ENTONCES() { return getToken(ProgramaParser.Reservada_ENTONCES, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public CondicionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condicional; }
	}

	public final CondicionalContext condicional() throws RecognitionException {
		CondicionalContext _localctx = new CondicionalContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_condicional);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(72);
			match(Reservada_SI);
			setState(73);
			expresion();
			setState(74);
			comparador();
			setState(75);
			expresion();
			setState(76);
			match(Reservada_ENTONCES);
			setState(77);
			instrucciones();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpresionContext extends ParserRuleContext {
		public TerminoContext termino() {
			return getRuleContext(TerminoContext.class,0);
		}
		public OperadorContext operador() {
			return getRuleContext(OperadorContext.class,0);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
	}

	public final ExpresionContext expresion() throws RecognitionException {
		ExpresionContext _localctx = new ExpresionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_expresion);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(79);
			termino();
			setState(83);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 62914560L) != 0)) {
				{
				setState(80);
				operador();
				setState(81);
				expresion();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TerminoContext extends ParserRuleContext {
		public TerminalNode NUMERO_ENTERO() { return getToken(ProgramaParser.NUMERO_ENTERO, 0); }
		public TerminalNode IDENTIFICADOR_VARIABLE() { return getToken(ProgramaParser.IDENTIFICADOR_VARIABLE, 0); }
		public TerminalNode OPERADOR_PARENTESIS_IZQ() { return getToken(ProgramaParser.OPERADOR_PARENTESIS_IZQ, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode OPERADOR_PARENTESIS_DER() { return getToken(ProgramaParser.OPERADOR_PARENTESIS_DER, 0); }
		public TerminoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_termino; }
	}

	public final TerminoContext termino() throws RecognitionException {
		TerminoContext _localctx = new TerminoContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_termino);
		try {
			setState(91);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NUMERO_ENTERO:
				enterOuterAlt(_localctx, 1);
				{
				setState(85);
				match(NUMERO_ENTERO);
				}
				break;
			case IDENTIFICADOR_VARIABLE:
				enterOuterAlt(_localctx, 2);
				{
				setState(86);
				match(IDENTIFICADOR_VARIABLE);
				}
				break;
			case OPERADOR_PARENTESIS_IZQ:
				enterOuterAlt(_localctx, 3);
				{
				setState(87);
				match(OPERADOR_PARENTESIS_IZQ);
				setState(88);
				expresion();
				setState(89);
				match(OPERADOR_PARENTESIS_DER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionvarContext extends ParserRuleContext {
		public TerminalNode ASIGNACION() { return getToken(ProgramaParser.ASIGNACION, 0); }
		public TerminalNode IGUALDAD() { return getToken(ProgramaParser.IGUALDAD, 0); }
		public AsignacionvarContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacionvar; }
	}

	public final AsignacionvarContext asignacionvar() throws RecognitionException {
		AsignacionvarContext _localctx = new AsignacionvarContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_asignacionvar);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(93);
			_la = _input.LA(1);
			if ( !(_la==ASIGNACION || _la==IGUALDAD) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ComparadorContext extends ParserRuleContext {
		public TerminalNode IGUALDAD() { return getToken(ProgramaParser.IGUALDAD, 0); }
		public TerminalNode COMPARADOR_MENOR() { return getToken(ProgramaParser.COMPARADOR_MENOR, 0); }
		public TerminalNode COMPARADOR_MAYOR() { return getToken(ProgramaParser.COMPARADOR_MAYOR, 0); }
		public TerminalNode COMPARADOR_MENOR_IGUAL() { return getToken(ProgramaParser.COMPARADOR_MENOR_IGUAL, 0); }
		public TerminalNode COMPARADOR_MAYOR_IGUAL() { return getToken(ProgramaParser.COMPARADOR_MAYOR_IGUAL, 0); }
		public TerminalNode COMPARADOR_DIFERENTE() { return getToken(ProgramaParser.COMPARADOR_DIFERENTE, 0); }
		public ComparadorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comparador; }
	}

	public final ComparadorContext comparador() throws RecognitionException {
		ComparadorContext _localctx = new ComparadorContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_comparador);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 16642999296L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperadorContext extends ParserRuleContext {
		public TerminalNode OPERADOR_SUMA() { return getToken(ProgramaParser.OPERADOR_SUMA, 0); }
		public TerminalNode OPERADOR_RESTA() { return getToken(ProgramaParser.OPERADOR_RESTA, 0); }
		public TerminalNode OPERADOR_MULTIPLICACION() { return getToken(ProgramaParser.OPERADOR_MULTIPLICACION, 0); }
		public TerminalNode OPERADOR_DIVISION() { return getToken(ProgramaParser.OPERADOR_DIVISION, 0); }
		public OperadorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operador; }
	}

	public final OperadorContext operador() throws RecognitionException {
		OperadorContext _localctx = new OperadorContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_operador);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(97);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 62914560L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001\'d\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0005\u0002*\b\u0002"+
		"\n\u0002\f\u0002-\t\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004"+
		"\u0005\u00043\b\u0004\n\u0004\f\u00046\t\u0004\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005C\b\u0005\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0003\bT\b\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0003\t\\\b\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\f\u0001"+
		"\f\u0001\f\u0000\u0000\r\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012"+
		"\u0014\u0016\u0018\u0000\u0003\u0001\u0000\t\n\u0002\u0000\n\n\u001d!"+
		"\u0001\u0000\u0016\u0019^\u0000\u001a\u0001\u0000\u0000\u0000\u0002\""+
		"\u0001\u0000\u0000\u0000\u0004&\u0001\u0000\u0000\u0000\u0006.\u0001\u0000"+
		"\u0000\u0000\b4\u0001\u0000\u0000\u0000\nB\u0001\u0000\u0000\u0000\fD"+
		"\u0001\u0000\u0000\u0000\u000eH\u0001\u0000\u0000\u0000\u0010O\u0001\u0000"+
		"\u0000\u0000\u0012[\u0001\u0000\u0000\u0000\u0014]\u0001\u0000\u0000\u0000"+
		"\u0016_\u0001\u0000\u0000\u0000\u0018a\u0001\u0000\u0000\u0000\u001a\u001b"+
		"\u0005\u0001\u0000\u0000\u001b\u001c\u0005\u0015\u0000\u0000\u001c\u001d"+
		"\u0005\u000b\u0000\u0000\u001d\u001e\u0003\u0002\u0001\u0000\u001e\u001f"+
		"\u0003\u0006\u0003\u0000\u001f \u0005\u0002\u0000\u0000 !\u0005\u000b"+
		"\u0000\u0000!\u0001\u0001\u0000\u0000\u0000\"#\u0005\b\u0000\u0000#$\u0003"+
		"\u0004\u0002\u0000$%\u0005\u000b\u0000\u0000%\u0003\u0001\u0000\u0000"+
		"\u0000&+\u0005\u0015\u0000\u0000\'(\u0005\f\u0000\u0000(*\u0005\u0015"+
		"\u0000\u0000)\'\u0001\u0000\u0000\u0000*-\u0001\u0000\u0000\u0000+)\u0001"+
		"\u0000\u0000\u0000+,\u0001\u0000\u0000\u0000,\u0005\u0001\u0000\u0000"+
		"\u0000-+\u0001\u0000\u0000\u0000./\u0005\u0003\u0000\u0000/0\u0003\b\u0004"+
		"\u00000\u0007\u0001\u0000\u0000\u000013\u0003\n\u0005\u000021\u0001\u0000"+
		"\u0000\u000036\u0001\u0000\u0000\u000042\u0001\u0000\u0000\u000045\u0001"+
		"\u0000\u0000\u00005\t\u0001\u0000\u0000\u000064\u0001\u0000\u0000\u0000"+
		"78\u0003\f\u0006\u000089\u0005\u000b\u0000\u00009C\u0001\u0000\u0000\u0000"+
		":;\u0005\u0006\u0000\u0000;<\u0005\u0015\u0000\u0000<C\u0005\u000b\u0000"+
		"\u0000=>\u0005\u0007\u0000\u0000>?\u0003\u0010\b\u0000?@\u0005\u000b\u0000"+
		"\u0000@C\u0001\u0000\u0000\u0000AC\u0003\u000e\u0007\u0000B7\u0001\u0000"+
		"\u0000\u0000B:\u0001\u0000\u0000\u0000B=\u0001\u0000\u0000\u0000BA\u0001"+
		"\u0000\u0000\u0000C\u000b\u0001\u0000\u0000\u0000DE\u0005\u0015\u0000"+
		"\u0000EF\u0003\u0014\n\u0000FG\u0003\u0010\b\u0000G\r\u0001\u0000\u0000"+
		"\u0000HI\u0005\u0004\u0000\u0000IJ\u0003\u0010\b\u0000JK\u0003\u0016\u000b"+
		"\u0000KL\u0003\u0010\b\u0000LM\u0005\u0005\u0000\u0000MN\u0003\b\u0004"+
		"\u0000N\u000f\u0001\u0000\u0000\u0000OS\u0003\u0012\t\u0000PQ\u0003\u0018"+
		"\f\u0000QR\u0003\u0010\b\u0000RT\u0001\u0000\u0000\u0000SP\u0001\u0000"+
		"\u0000\u0000ST\u0001\u0000\u0000\u0000T\u0011\u0001\u0000\u0000\u0000"+
		"U\\\u0005\u0014\u0000\u0000V\\\u0005\u0015\u0000\u0000WX\u0005\u001a\u0000"+
		"\u0000XY\u0003\u0010\b\u0000YZ\u0005\u001b\u0000\u0000Z\\\u0001\u0000"+
		"\u0000\u0000[U\u0001\u0000\u0000\u0000[V\u0001\u0000\u0000\u0000[W\u0001"+
		"\u0000\u0000\u0000\\\u0013\u0001\u0000\u0000\u0000]^\u0007\u0000\u0000"+
		"\u0000^\u0015\u0001\u0000\u0000\u0000_`\u0007\u0001\u0000\u0000`\u0017"+
		"\u0001\u0000\u0000\u0000ab\u0007\u0002\u0000\u0000b\u0019\u0001\u0000"+
		"\u0000\u0000\u0005+4BS[";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}