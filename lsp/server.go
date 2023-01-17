package lsp

import (
	"fmt"
	"github.com/pulumi/pulumi-lsp/sdk/lsp"
	"go.lsp.dev/protocol"
)

type server struct {
	docs map[protocol.DocumentURI]*document
}

func (s *server) setDocument(text lsp.Document) *document {
	doc := &document{text: text, server: s}
	s.docs[text.URI()] = doc
	return doc
}

func (s *server) getDocument(uri protocol.DocumentURI) (*document, bool) {
	d, ok := s.docs[uri]
	return d, ok
}

func (s *server) didOpen(client lsp.Client, params *protocol.DidOpenTextDocumentParams) error {
	fileName := params.TextDocument.URI.Filename()
	text := params.TextDocument.Text
	err := client.LogDebugf("Opened file %s:\n---\n%s---", fileName, text)

	s.setDocument(lsp.NewDocument(params.TextDocument)).process(client)
	return err
}

func (s *server) didClose(client lsp.Client, params *protocol.DidCloseTextDocumentParams) error {
	uri := params.TextDocument.URI
	_ = client.LogDebugf("Closing file %s", uri.Filename())

	_, ok := s.docs[uri]
	if !ok {
		_ = client.LogWarningf("Attempted to close unopened file %s", uri.Filename())
	}

	delete(s.docs, uri)
	return nil
}

func (s *server) didChange(client lsp.Client, params *protocol.DidChangeTextDocumentParams) error {
	uri := params.TextDocument.URI

	doc, ok := s.getDocument(params.TextDocument.URI)
	if !ok {
		return fmt.Errorf("could not find document %s(%s)", uri.Filename(), uri)
	}

	if err := doc.text.AcceptChanges(params.ContentChanges); err != nil {
		return fmt.Errorf("document might be unknown: %w", err)
	}

	doc.process(client)
	return nil
}

func (s *server) semanticTokens(
	_ lsp.Client,
	params *protocol.SemanticTokensParams,
) (result *protocol.SemanticTokens, err error) {
	uri := params.TextDocument.URI
	document, ok := s.getDocument(uri)
	if !ok {
		return nil, fmt.Errorf("open document not found %s", uri.Filename())
	}

	// [ lineDelta, charDelta(rel to prev start same line), length, type, modifiers, ... ]
	data := make([]uint32, 0, len(document.tokens)*5)

	previousLine := 0
	previousStart := 0
	for _, token := range document.tokens {
		line := document.getCharLine(token.Start)

		// Compute the line delta (how many lines this token is past the previous)
		data = append(data, uint32(line-previousLine))

		// If we're on a new line, reset the start-character delta
		if line > previousLine {
			previousStart = token.Start
		}

		previousLine = line

		// Compute the start-character delta
		data = append(data, uint32(token.Start-previousStart))
		previousStart = token.Start

		// Compute the token's length
		data = append(data, uint32(token.End-token.Start))

		// Token type (keyword, function, etc)
		data = append(data, uint32(token.Kind))

		// Token modifiers (static, const, etc)
		data = append(data, 0)
	}

	// TODO: endian
	return &protocol.SemanticTokens{Data: data}, nil
}
