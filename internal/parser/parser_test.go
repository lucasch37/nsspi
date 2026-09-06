package parser

import (
	stderrors "errors"
	"testing"

	projecterrors "github.com/lucasch37/nsspi/internal/errors"
	"github.com/lucasch37/nsspi/internal/ir"
	"github.com/lucasch37/nsspi/internal/lexer"
	"github.com/lucasch37/nsspi/internal/tokens"
)

func TestParseProgramStructure(t *testing.T) {
	program := parseProgram(t, `
	PROGRAM Example;
	VAR answer: INTEGER;
	BEGIN
		answer := 2 + 3 * 4
	END.`)

	if program.Name != "Example" {
		t.Errorf("program name = %q, want Example", program.Name)
	}

	if len(program.Block.Declarations) != 1 {
		t.Fatalf("declarations = %d, want 1", len(program.Block.Declarations))
	}

	decl, ok := program.Block.Declarations[0].(*ir.VarDecl)
	if !ok {
		t.Fatalf("declaration type = %T, want *ir.VarDecl", program.Block.Declarations[0])
	}

	if decl.IdNode.Value != "answer" || decl.TypeNode.Value != "INTEGER" {
		t.Errorf("declaration = %s: %s, want answer: INTEGER", decl.IdNode.Value, decl.TypeNode.Value)
	}

	assignment, ok := program.Block.CompoundStatement.Children[0].(*ir.Assign)
	if !ok {
		t.Fatalf("statement type = %T, want *ir.Assign", program.Block.CompoundStatement.Children[0])
	}

	addition, ok := assignment.Right.(*ir.BinOp)
	if !ok || addition.Op.Type != tokens.PLUS {
		t.Fatalf("assignment expression = %#v, want addition", assignment.Right)
	}

	multiplication, ok := addition.Right.(*ir.BinOp)
	if !ok || multiplication.Op.Type != tokens.MUL {
		t.Fatalf("addition right operand = %#v, want multiplication", addition.Right)
	}
}

func TestParseLanguageConstructs(t *testing.T) {
	program := parseProgram(t, `
	PROGRAM Example;
	FUNCTION Double(n: INTEGER): INTEGER;
	BEGIN
		Double := n * 2
	END;
	PROCEDURE Show(value: INTEGER);
	BEGIN
		IF value > 0 THEN WRITELN(value) ELSE WRITE('zero')
	END;
	BEGIN
		Show(Double(3))
	END.`)

	if len(program.Block.Declarations) != 2 {
		t.Fatalf("declarations = %d, want 2", len(program.Block.Declarations))
	}

	function, ok := program.Block.Declarations[0].(*ir.FunctionDecl)
	if !ok || function.FuncName != "Double" || len(function.Params) != 1 {
		t.Fatalf("first declaration = %#v, want one-parameter Double function", program.Block.Declarations[0])
	}

	procedure, ok := program.Block.Declarations[1].(*ir.ProcedureDecl)
	if !ok || procedure.ProcName != "Show" || len(procedure.Params) != 1 {
		t.Fatalf("second declaration = %#v, want one-parameter Show procedure", program.Block.Declarations[1])
	}

	if _, ok := procedure.Block.CompoundStatement.Children[0].(*ir.IfStatement); !ok {
		t.Errorf("procedure statement type = %T, want *ir.IfStatement", procedure.Block.CompoundStatement.Children[0])
	}

	call, ok := program.Block.CompoundStatement.Children[0].(*ir.Call)
	if !ok || call.CallName != "Show" || len(call.ActualParams) != 1 {
		t.Fatalf("main statement = %#v, want Show call", program.Block.CompoundStatement.Children[0])
	}
}

func TestParseRejectsInvalidPrograms(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{name: "missing program semicolon", source: "PROGRAM Bad BEGIN END."},
		{name: "missing assignment expression", source: "PROGRAM Bad; VAR x: INTEGER; BEGIN x := END."},
		{name: "trailing input", source: "PROGRAM Bad; BEGIN END. BEGIN END"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewParser(lexer.NewLexer(tt.source))
			if err == nil {
				_, err = p.Parse()
			}
			if err == nil {
				t.Fatal("Parse() returned nil, want syntax error")
			}

			var syntaxErr *projecterrors.SyntaxError
			if !stderrors.As(err, &syntaxErr) {
				t.Fatalf("error type = %T, want *errors.SyntaxError", err)
			}
		})
	}
}

func parseProgram(t *testing.T, source string) *ir.Program {
	t.Helper()

	p, err := NewParser(lexer.NewLexer(source))
	if err != nil {
		t.Fatalf("NewParser() error: %v", err)
	}

	tree, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse() error: %v", err)
	}

	program, ok := tree.(*ir.Program)
	if !ok {
		t.Fatalf("Parse() returned %T, want *ir.Program", tree)
	}

	return program
}
