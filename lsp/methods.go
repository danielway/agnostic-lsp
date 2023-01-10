package lsp

import (
	"github.com/pulumi/pulumi-lsp/sdk/lsp"
	"go.lsp.dev/protocol"
)

func Methods() *lsp.Methods {
	server := &server{
		docs: map[protocol.DocumentURI]*document{},
	}

	return lsp.Methods{
		DidOpenFunc:   server.didOpen,
		DidCloseFunc:  server.didClose,
		DidChangeFunc: server.didChange,
		// TODO: implement agnostic methods
	}.DefaultInitializer("agnostic-lsp", `0.0.1`)
}
