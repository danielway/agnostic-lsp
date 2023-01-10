package lsp

import "github.com/pulumi/pulumi-lsp/sdk/lsp"

type document struct {
	text   lsp.Document
	server *server
}

func (d *document) process(client lsp.Client) {
	// TODO: process file with Agnostic lexer/parser
}
