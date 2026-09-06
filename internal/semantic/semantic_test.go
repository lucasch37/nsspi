package semantic

import (
	stderrors "errors"
	"os"
	"path/filepath"
	"testing"

	projecterrors "github.com/lucasch37/nsspi/internal/errors"
	"github.com/lucasch37/nsspi/internal/lexer"
	"github.com/lucasch37/nsspi/internal/parser"
)

func TestAnalyzeExamples(t *testing.T) {
	entries, err := os.ReadDir(filepath.Join("..", "..", "examples"))
	if err != nil {
		t.Fatalf("read examples directory: %v", err)
	}
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || filepath.Ext(name) != ".pas" {
			continue
		}
		t.Run(name, func(t *testing.T) {
			source, err := os.ReadFile(filepath.Join("..", "..", "examples", name))
			if err != nil {
				t.Fatalf("read example: %v", err)
			}
			if err := analyzeSource(t, string(source)); err != nil {
				t.Fatalf("Analyze() error: %v", err)
			}
		})
	}
}

func TestAnalyzeAcceptsValidPrograms(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{
			name: "assignments and expressions",
			source: `
PROGRAM Valid;
VAR
	i: INTEGER;
	r: REAL;
	s: STRING;
	b: BOOLEAN;
BEGIN
	i := 2 + 3 * 4;
	r := 1.5 + 2.0;
	s := 'a' + 'b';
	b := i < 10
END.
`,
		},
		{
			name: "procedure and function calls",
			source: `
PROGRAM Valid;
FUNCTION Double(n: INTEGER): INTEGER;
BEGIN
	Double := n * 2
END;
PROCEDURE Show(n: INTEGER);
BEGIN
	WRITELN(n)
END;
BEGIN
	Show(Double(3))
END.
`,
		},
		{
			name: "nested scope",
			source: `
PROGRAM Valid;
VAR
	x: INTEGER;
PROCEDURE SetValue(n: INTEGER);
VAR
	y: INTEGER;
BEGIN
	y := n;
	x := y
END;
BEGIN
	x := 0;
	SetValue(4)
END.
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := analyzeSource(t, tt.source); err != nil {
				t.Fatalf("Analyze() error: %v", err)
			}
		})
	}
}

func TestAnalyzeRejectsInvalidPrograms(t *testing.T) {
	tests := []struct {
		name   string
		code   projecterrors.ErrorCode
		source string
	}{
		{
			name: "unknown identifier",
			code: projecterrors.IDNotFound,
			source: `
PROGRAM Bad;
BEGIN
	missing := 1
END.
`,
		},
		{
			name: "duplicate identifier",
			code: projecterrors.DuplicateID,
			source: `
PROGRAM Bad;
VAR
	x: INTEGER;
	x: REAL;
BEGIN
END.
`,
		},
		{
			name: "assignment type mismatch",
			code: projecterrors.TypeMismatch,
			source: `
PROGRAM Bad;
VAR
	x: INTEGER;
BEGIN
	x := 'text'
END.
`,
		},
		{
			name: "invalid arithmetic operands",
			code: projecterrors.TypeMismatch,
			source: `
PROGRAM Bad;
VAR
	x: INTEGER;
BEGIN
	x := TRUE + 1
END.
`,
		},
		{
			name: "wrong parameter count",
			code: projecterrors.WrongParamCount,
			source: `
PROGRAM Bad;
PROCEDURE P(x: INTEGER);
BEGIN
END;
BEGIN
	P()
END.
`,
		},
		{
			name: "wrong parameter type",
			code: projecterrors.TypeMismatch,
			source: `
PROGRAM Bad;
PROCEDURE P(x: INTEGER);
BEGIN
END;
BEGIN
	P('text')
END.
`,
		},
		{
			name: "invalid for statement",
			code: projecterrors.TypeMismatch,
			source: `
PROGRAM Bad;
VAR
	a: INTEGER;
PROCEDURE P(x: INTEGER);
BEGIN
	FOR a := 1 TO '10' DO
		WRITELN(x)
END;
BEGIN
	P('text')
END.
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := analyzeSource(t, tt.source)
			if err == nil {
				t.Fatal("Analyze() returned nil, want semantic error")
			}

			var semanticErr *projecterrors.SemanticError
			if !stderrors.As(err, &semanticErr) {
				t.Fatalf("error type = %T, want *errors.SemanticError", err)
			}

			if semanticErr.Code != tt.code {
				t.Errorf("error code = %v, want %v", semanticErr.Code, tt.code)
			}
		})
	}
}

func analyzeSource(t *testing.T, source string) error {
	t.Helper()

	p, err := parser.NewParser(lexer.NewLexer(source))
	if err != nil {
		t.Fatalf("NewParser() error: %v", err)
	}

	tree, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse() error: %v", err)
	}

	return NewSemanticAnalyzer(false).Analyze(tree)
}
