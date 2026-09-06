package interpreter

import (
	"io"
	"os"
	"testing"

	projecterrors "github.com/lucasch37/nsspi/internal/errors"
	"github.com/lucasch37/nsspi/internal/ir"
	"github.com/lucasch37/nsspi/internal/lexer"
	"github.com/lucasch37/nsspi/internal/parser"
	"github.com/lucasch37/nsspi/internal/semantic"
)

func TestInterpretProgramOutput(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name: "arithmetic and precedence",
			source: `
PROGRAM Test;
BEGIN
	WRITELN(2 + 3 * 4);
	WRITELN((2 + 3) * 4);
	WRITELN(7 DIV 2);
	WRITELN(7 MOD 2)
END.
`,
			want: "14\n20\n3\n1\n",
		},
		{
			name: "real unary and string expressions",
			source: `
PROGRAM Test;
BEGIN
	WRITELN(7 / 2);
	WRITELN(-1.5);
	WRITELN('hello' + ' world')
END.
`,
			want: "3.5\n-1.5\nhello world\n",
		},
		{
			name: "conditionals",
			source: `
PROGRAM Test;
BEGIN
	IF 2 < 3 THEN
		WRITELN('yes')
	ELSE
		WRITELN('no');
	IF 1 = 2 THEN
		WRITE('bad')
	ELSE
		WRITELN('ok')
END.
`,
			want: "yes\nok\n",
		},
		{
			name: "procedure and recursive function",
			source: `
PROGRAM Test;
FUNCTION Fact(n: INTEGER): INTEGER;
BEGIN
	IF n = 0 THEN
		Fact := 1
	ELSE
		Fact := n * Fact(n - 1)
END;
PROCEDURE Show(n: INTEGER);
BEGIN
	WRITELN(n)
END;
BEGIN
	Show(Fact(5))
END.
`,
			want: "120\n",
		},
		{
			name: "for loops",
			source: `
PROGRAM TestNestedFor;

VAR
	i, j, value: INTEGER;

BEGIN
	value := 0;

	FOR i := 1 TO 5 DO
	BEGIN
		FOR j := 1 TO 5 DO
		BEGIN
			value := value + 1
		END
	END;

	WRITELN(value)
END.
`,
			want: "25\n",
		},
		{
			name: "while loops",
			source: `
PROGRAM TestWhile;

VAR
	i, value: INTEGER;

BEGIN
	value := 0;
	i := 0;

	WHILE i <> 10 DO
	BEGIN
		i := i + 1;
		value := value + i
	END;

	WRITELN(value)
END.
`,
			want: "55\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := captureOutput(t, func() {
				if err := interpretSource(t, tt.source); err != nil {
					t.Fatalf("Interpret() error: %v", err)
				}
			})

			if got != tt.want {
				t.Errorf("output = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestInterpretNilTree(t *testing.T) {
	if err := NewInterpreter(nil, false).Interpret(); err != nil {
		t.Fatalf("Interpret() error: %v", err)
	}
}

func TestIntegerDivisionByZero(t *testing.T) {
	source := `
PROGRAM Test;
BEGIN
	WRITELN(1 DIV 0)
END.
`
	p, err := parser.NewParser(lexer.NewLexer(source))
	if err != nil {
		t.Fatal(err)
	}

	tree, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if err := semantic.NewSemanticAnalyzer(false).Analyze(tree); err != nil {
		t.Fatal(err)
	}

	program := tree.(*ir.Program)
	write := program.Block.CompoundStatement.Children[0].(*ir.WriteStatement)

	_, err = NewInterpreter(tree, false).visit(write.Exprs[0])
	if err == nil {
		t.Fatal("division returned nil, want runtime error")
	}

	if runtimeErr, ok := err.(*projecterrors.RuntimeError); !ok || runtimeErr.Code != projecterrors.DivideByZero {
		t.Fatalf("error = %#v, want DivideByZero RuntimeError", err)
	}
}

func interpretSource(t *testing.T, source string) error {
	t.Helper()
	p, err := parser.NewParser(lexer.NewLexer(source))
	if err != nil {
		t.Fatalf("NewParser() error: %v", err)
	}

	tree, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse() error: %v", err)
	}

	if err := semantic.NewSemanticAnalyzer(false).Analyze(tree); err != nil {
		t.Fatalf("Analyze() error: %v", err)
	}

	return NewInterpreter(tree, false).Interpret()
}

func captureOutput(t *testing.T, fn func()) string {
	t.Helper()

	read, write, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	original := os.Stdout
	os.Stdout = write
	defer func() { os.Stdout = original }()

	fn()

	if err := write.Close(); err != nil {
		t.Fatal(err)
	}

	output, err := io.ReadAll(read)
	if err != nil {
		t.Fatal(err)
	}

	if err := read.Close(); err != nil {
		t.Fatal(err)
	}

	return string(output)
}
