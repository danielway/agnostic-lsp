package lsp

import (
	"github.com/pulumi/pulumi-lsp/sdk/lsp"
	"go.lsp.dev/protocol"
)

func Methods() *lsp.Methods {
	server := &server{
		docs: map[protocol.DocumentURI]*document{},
	}

	methods := lsp.Methods{
		DidOpenFunc:            server.didOpen,
		DidCloseFunc:           server.didClose,
		DidChangeFunc:          server.didChange,
		SemanticTokensFullFunc: server.semanticTokens,
	}

	methods.InitializeFunc = func(client lsp.Client, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
		return &protocol.InitializeResult{
			Capabilities: protocol.ServerCapabilities{
				TextDocumentSync: &protocol.TextDocumentSyncOptions{
					OpenClose: methods.DidOpenFunc != nil || methods.DidCloseFunc != nil,
					Change:    protocol.TextDocumentSyncKindFull,
				},
				SemanticTokensProvider: &SemanticTokensOptions{
					Types: []string{
						// See `lexer.TokenKind`
						`keyword`,
						`type`,
						`module`,
						`function`,
						`comment`,
					},
				},
			},
			ServerInfo: &protocol.ServerInfo{
				Name:    `agnostic-lsp`,
				Version: `0.0.1`,
			},
		}, nil
	}

	return &methods
}

type SemanticTokensOptions struct {
	Types []string `json:"tokenTypes"`
}
