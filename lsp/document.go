package lsp

import (
	"github.com/JosephNaberhaus/agnostic/language/lexer"
	"github.com/pulumi/pulumi-lsp/sdk/lsp"
)

type document struct {
	text   lsp.Document
	server *server
}

func (d *document) process(client lsp.Client) {
	ast, _, err := lexer.Parse(d.text.String())
	if err != nil {
		_ = client.LogErrorf("failed to parse document: %w", err)
		return
	}

	_ = client.LogInfof("Document parsed/lexed. Module name: %s", ast.Name)
}
