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
	client lsp.Client,
	params *protocol.SemanticTokensParams,
) (result *protocol.SemanticTokens, err error) {
	var data []uint32

	// TODO: populate `data`

	return &protocol.SemanticTokens{Data: data}, nil
}
