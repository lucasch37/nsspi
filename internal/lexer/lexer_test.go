package lexer

import (
	stderrors "errors"
	"reflect"
	"strings"
	"testing"

	lexererrors "github.com/lucasch37/nsspi/internal/errors"
	"github.com/lucasch37/nsspi/internal/tokens"
)

func TestGetNextToken(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []tokens.Token
	}{
		{
			name:  "single-character tokens",
			input: "+ - * / ( ) ; . : , > < =",
			want: []tokens.Token{
				{Type: tokens.PLUS, Value: "+"}, {Type: tokens.MINUS, Value: "-"},
				{Type: tokens.MUL, Value: "*"}, {Type: tokens.FLOAT_DIV, Value: "/"},
				{Type: tokens.LPAREN, Value: "("}, {Type: tokens.RPAREN, Value: ")"},
				{Type: tokens.SEMI, Value: ";"}, {Type: tokens.DOT, Value: "."},
				{Type: tokens.COLON, Value: ":"}, {Type: tokens.COMMA, Value: ","},
				{Type: tokens.GREATER_THAN, Value: ">"}, {Type: tokens.LESS_THAN, Value: "<"},
				{Type: tokens.EQUAL, Value: "="}, {Type: tokens.EOF},
			},
		},
		{
			name:  "multi-character operators",
			input: ":= <= >= <>",
			want: []tokens.Token{
				{Type: tokens.ASSIGN, Value: ":="},
				{Type: tokens.LESS_THAN_EQUAL, Value: "<="},
				{Type: tokens.GREATER_THAN_EQUAL, Value: ">="},
				{Type: tokens.NOT_EQUAL, Value: "<>"},
				{Type: tokens.EOF},
			},
		},
		{
			name:  "identifiers and literals",
			input: "count2 42 3.14 'hello world' TRUE false",
			want: []tokens.Token{
				{Type: tokens.ID, Value: "count2"},
				{Type: tokens.INTEGER_LIT, Value: 42},
				{Type: tokens.REAL_LIT, Value: 3.14},
				{Type: tokens.STRING_LIT, Value: "hello world"},
				{Type: tokens.TRUE, Value: true},
				{Type: tokens.FALSE, Value: false},
				{Type: tokens.EOF},
			},
		},
		{
			name:  "whitespace and comments are skipped",
			input: "\tVAR\r\n{ ignore these tokens: := 99 }\vanswer\f:= 42",
			want: []tokens.Token{
				{Type: tokens.VAR, Value: "VAR"},
				{Type: tokens.ID, Value: "answer"},
				{Type: tokens.ASSIGN, Value: ":="},
				{Type: tokens.INTEGER_LIT, Value: 42},
				{Type: tokens.EOF},
			},
		},
		{
			name:  "case-insensitive keywords",
			input: "program INTEGER real String div var PROCEDURE function write writeln boolean if then else mod begin end WHILE",
			want: []tokens.Token{
				{Type: tokens.PROGRAM, Value: "PROGRAM"}, {Type: tokens.INTEGER, Value: "INTEGER"},
				{Type: tokens.REAL, Value: "REAL"}, {Type: tokens.STRING, Value: "STRING"},
				{Type: tokens.INTEGER_DIV, Value: "DIV"}, {Type: tokens.VAR, Value: "VAR"},
				{Type: tokens.PROCEDURE, Value: "PROCEDURE"}, {Type: tokens.FUNCTION, Value: "FUNCTION"},
				{Type: tokens.WRITE, Value: "WRITE"}, {Type: tokens.WRITELN, Value: "WRITELN"},
				{Type: tokens.BOOLEAN, Value: "BOOLEAN"}, {Type: tokens.IF, Value: "IF"},
				{Type: tokens.THEN, Value: "THEN"}, {Type: tokens.ELSE, Value: "ELSE"},
				{Type: tokens.MOD, Value: "MOD"}, {Type: tokens.BEGIN, Value: "BEGIN"},
				{Type: tokens.END, Value: "END"}, {Type: tokens.WHILE, Value: "WHILE"}, {Type: tokens.EOF},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lexAll(t, tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("tokens mismatch\n got: %#v\nwant: %#v", got, tt.want)
			}
		})
	}
}

func TestGetNextTokenTracksPosition(t *testing.T) {
	lexer := NewLexer("\n  :=")

	got, err := lexer.GetNextToken()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.LineNo != 2 || got.Column != 3 {
		t.Errorf("got position %d:%d, want 2:3", got.LineNo, got.Column)
	}
}

func TestGetNextTokenErrors(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		message string
	}{
		{name: "invalid character", input: "@", message: "Invalid character -> line: 1 column: 1"},
		{name: "unterminated comment", input: "{ never closed", message: "Unterminated comment -> line: 1 column: 14"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lexer := NewLexer(tt.input)

			_, err := lexer.GetNextToken()
			if err == nil {
				t.Fatal("expected a lexical error")
			}

			var lexicalErr *lexererrors.LexicalError
			if !stderrors.As(err, &lexicalErr) {
				t.Fatalf("got error type %T, want *errors.LexicalError", err)
			}

			if !strings.Contains(err.Error(), tt.message) {
				t.Errorf("error %q does not contain %q", err, tt.message)
			}
		})
	}
}

func lexAll(t *testing.T, input string) []tokens.Token {
	t.Helper()
	lexer := NewLexer(input)
	var result []tokens.Token
	for {
		token, err := lexer.GetNextToken()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		token.LineNo = 0
		token.Column = 0
		result = append(result, token)

		if token.Type == tokens.EOF {
			return result
		}
	}
}
