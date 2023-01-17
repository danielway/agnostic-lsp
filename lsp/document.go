package lsp

import (
	"github.com/JosephNaberhaus/agnostic/ast"
	"github.com/JosephNaberhaus/agnostic/language/lexer"
	"github.com/pulumi/pulumi-lsp/sdk/lsp"
)

type document struct {
	text   lsp.Document
	ast    ast.Module
	tokens []lexer.TokenMeta
	lines  []int // \n character indices
	server *server
}

func (d *document) process(client lsp.Client) {
	text := d.text.String()
	ast, tokens, err := lexer.Parse(text)
	if err != nil {
		_ = client.LogErrorf("failed to parse document: %w", err)
		return
	}

	d.ast = ast
	d.tokens = tokens

	d.lines = make([]int, 0)
	for pos, ch := range text {
		if ch == '\n' {
			d.lines = append(d.lines, pos)
		}
	}

	_ = client.LogInfof("Document parsed/lexed. Module name: %s", ast.Name)
}

func (d *document) getCharLine(find int) int {
	for line, pos := range d.lines {
		if pos > find {
			return line - 1
		}
	}

	return len(d.lines) - 1
}
