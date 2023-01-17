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
	server *server
}

func (d *document) process(client lsp.Client) {
	ast, tokens, err := lexer.Parse(d.text.String())
	if err != nil {
		_ = client.LogErrorf("failed to parse document: %w", err)
		return
	}

	d.ast = ast
	d.tokens = tokens

	_ = client.LogInfof("Document parsed/lexed. Module name: %s", ast.Name)
}
