// Generated from /home/jigc4200/Documents/antlr-go-example/Programa.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link ProgramaParser}.
 */
public interface ProgramaListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#programa}.
	 * @param ctx the parse tree
	 */
	void enterPrograma(ProgramaParser.ProgramaContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#programa}.
	 * @param ctx the parse tree
	 */
	void exitPrograma(ProgramaParser.ProgramaContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#declaraciones}.
	 * @param ctx the parse tree
	 */
	void enterDeclaraciones(ProgramaParser.DeclaracionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#declaraciones}.
	 * @param ctx the parse tree
	 */
	void exitDeclaraciones(ProgramaParser.DeclaracionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#lista_variables}.
	 * @param ctx the parse tree
	 */
	void enterLista_variables(ProgramaParser.Lista_variablesContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#lista_variables}.
	 * @param ctx the parse tree
	 */
	void exitLista_variables(ProgramaParser.Lista_variablesContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#bloque}.
	 * @param ctx the parse tree
	 */
	void enterBloque(ProgramaParser.BloqueContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#bloque}.
	 * @param ctx the parse tree
	 */
	void exitBloque(ProgramaParser.BloqueContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterInstrucciones(ProgramaParser.InstruccionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitInstrucciones(ProgramaParser.InstruccionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#instruccion}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion(ProgramaParser.InstruccionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#instruccion}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion(ProgramaParser.InstruccionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion(ProgramaParser.AsignacionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion(ProgramaParser.AsignacionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#condicional}.
	 * @param ctx the parse tree
	 */
	void enterCondicional(ProgramaParser.CondicionalContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#condicional}.
	 * @param ctx the parse tree
	 */
	void exitCondicional(ProgramaParser.CondicionalContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterExpresion(ProgramaParser.ExpresionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitExpresion(ProgramaParser.ExpresionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#termino}.
	 * @param ctx the parse tree
	 */
	void enterTermino(ProgramaParser.TerminoContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#termino}.
	 * @param ctx the parse tree
	 */
	void exitTermino(ProgramaParser.TerminoContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#comparador}.
	 * @param ctx the parse tree
	 */
	void enterComparador(ProgramaParser.ComparadorContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#comparador}.
	 * @param ctx the parse tree
	 */
	void exitComparador(ProgramaParser.ComparadorContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#operador}.
	 * @param ctx the parse tree
	 */
	void enterOperador(ProgramaParser.OperadorContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#operador}.
	 * @param ctx the parse tree
	 */
	void exitOperador(ProgramaParser.OperadorContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#identificador}.
	 * @param ctx the parse tree
	 */
	void enterIdentificador(ProgramaParser.IdentificadorContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#identificador}.
	 * @param ctx the parse tree
	 */
	void exitIdentificador(ProgramaParser.IdentificadorContext ctx);
	/**
	 * Enter a parse tree produced by {@link ProgramaParser#numero}.
	 * @param ctx the parse tree
	 */
	void enterNumero(ProgramaParser.NumeroContext ctx);
	/**
	 * Exit a parse tree produced by {@link ProgramaParser#numero}.
	 * @param ctx the parse tree
	 */
	void exitNumero(ProgramaParser.NumeroContext ctx);
}