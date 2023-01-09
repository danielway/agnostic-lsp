package main

import (
	"fmt"
	"context"
	"go.lsp.dev/protocol"
)

var _ protocol.Server = &AgnosticServer{}

type AgnosticServer struct {}

func (server *AgnosticServer) Initialize(ctx context.Context, params *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
	return nil, nil
}

func (server *AgnosticServer) Initialized(ctx context.Context, params *protocol.InitializedParams) (err error) {
	return nil
}

func (server *AgnosticServer) Shutdown(ctx context.Context) (err error) {
	return nil
}

func (server *AgnosticServer) Exit(ctx context.Context) (err error) {
	return nil
}

func (server *AgnosticServer) WorkDoneProgressCancel(ctx context.Context, params *protocol.WorkDoneProgressCancelParams) (err error) {
	return nil
}

func (server *AgnosticServer) LogTrace(ctx context.Context, params *protocol.LogTraceParams) (err error) {
	return nil
}

func (server *AgnosticServer) SetTrace(ctx context.Context, params *protocol.SetTraceParams) (err error) {
	return nil
}

func (server *AgnosticServer) CodeAction(ctx context.Context, params *protocol.CodeActionParams) (result []protocol.CodeAction, err error) {
	return nil, nil
}

func (server *AgnosticServer) CodeLens(ctx context.Context, params *protocol.CodeLensParams) (result []protocol.CodeLens, err error) {
	return nil, nil
}

func (server *AgnosticServer) CodeLensResolve(ctx context.Context, params *protocol.CodeLens) (result *protocol.CodeLens, err error) {
	return nil, nil
}

func (server *AgnosticServer) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) (result []protocol.ColorPresentation, err error) {
	return nil, nil
}

func (server *AgnosticServer) Completion(ctx context.Context, params *protocol.CompletionParams) (result *protocol.CompletionList, err error) {
	return nil, nil
}

func (server *AgnosticServer) CompletionResolve(ctx context.Context, params *protocol.CompletionItem) (result *protocol.CompletionItem, err error) {
	return nil, nil
}

func (server *AgnosticServer) Declaration(ctx context.Context, params *protocol.DeclarationParams) (result []protocol.Location /* Declaration | DeclarationLink[] | null */, err error) {
	return nil, nil
}

func (server *AgnosticServer) Definition(ctx context.Context, params *protocol.DefinitionParams) (result []protocol.Location /* Definition | DefinitionLink[] | null */, err error) {
	return nil, nil
}

func (server *AgnosticServer) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (err error) {
	return nil
}

func (server *AgnosticServer) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) (err error) {
	return nil
}

func (server *AgnosticServer) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) (err error) {
	return nil
}

func (server *AgnosticServer) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) (err error) {
	return nil
}

func (server *AgnosticServer) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	return nil
}

func (server *AgnosticServer) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	return nil
}

func (server *AgnosticServer) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	return nil
}

func (server *AgnosticServer) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) (result []protocol.ColorInformation, err error) {
	return nil, nil
}

func (server *AgnosticServer) DocumentHighlight(ctx context.Context, params *protocol.DocumentHighlightParams) (result []protocol.DocumentHighlight, err error) {
	return nil, nil
}

func (server *AgnosticServer) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) (result []protocol.DocumentLink, err error) {
	return nil, nil
}

func (server *AgnosticServer) DocumentLinkResolve(ctx context.Context, params *protocol.DocumentLink) (result *protocol.DocumentLink, err error) {
	return nil, nil
}

func (server *AgnosticServer) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) (result []interface{} /* []SymbolInformation | []DocumentSymbol */, err error) {
	return nil, nil
}

func (server *AgnosticServer) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (result interface{}, err error) {
	return nil, nil
}

func (server *AgnosticServer) FoldingRanges(ctx context.Context, params *protocol.FoldingRangeParams) (result []protocol.FoldingRange, err error) {
	return nil, nil
}

func (server *AgnosticServer) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) (result []protocol.TextEdit, err error) {
	return nil, nil
}

func (server *AgnosticServer) Hover(ctx context.Context, params *protocol.HoverParams) (result *protocol.Hover, err error) {
	return nil, nil
}

func (server *AgnosticServer) Implementation(ctx context.Context, params *protocol.ImplementationParams) (result []protocol.Location, err error) {
	return nil, nil
}

func (server *AgnosticServer) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) (result []protocol.TextEdit, err error) {
	return nil, nil
}

func (server *AgnosticServer) PrepareRename(ctx context.Context, params *protocol.PrepareRenameParams) (result *protocol.Range, err error) {
	return nil, nil
}

func (server *AgnosticServer) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) (result []protocol.TextEdit, err error) {
	return nil, nil
}

func (server *AgnosticServer) References(ctx context.Context, params *protocol.ReferenceParams) (result []protocol.Location, err error) {
	return nil, nil
}

func (server *AgnosticServer) Rename(ctx context.Context, params *protocol.RenameParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, nil
}

func (server *AgnosticServer) SignatureHelp(ctx context.Context, params *protocol.SignatureHelpParams) (result *protocol.SignatureHelp, err error) {
	return nil, nil
}

func (server *AgnosticServer) Symbols(ctx context.Context, params *protocol.WorkspaceSymbolParams) (result []protocol.SymbolInformation, err error) {
	return nil, nil
}

func (server *AgnosticServer) TypeDefinition(ctx context.Context, params *protocol.TypeDefinitionParams) (result []protocol.Location, err error) {
	return nil, nil
}

func (server *AgnosticServer) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (err error) {
	return nil
}

func (server *AgnosticServer) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (result []protocol.TextEdit, err error) {
	return nil, nil
}

func (server *AgnosticServer) ShowDocument(ctx context.Context, params *protocol.ShowDocumentParams) (result *protocol.ShowDocumentResult, err error) {
	return nil, nil
}

func (server *AgnosticServer) WillCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, nil
}

func (server *AgnosticServer) DidCreateFiles(ctx context.Context, params *protocol.CreateFilesParams) (err error) {
	return nil
}

func (server *AgnosticServer) WillRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, nil
}

func (server *AgnosticServer) DidRenameFiles(ctx context.Context, params *protocol.RenameFilesParams) (err error) {
	return nil
}

func (server *AgnosticServer) WillDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (result *protocol.WorkspaceEdit, err error) {
	return nil, nil
}

func (server *AgnosticServer) DidDeleteFiles(ctx context.Context, params *protocol.DeleteFilesParams) (err error) {
	return nil
}

func (server *AgnosticServer) CodeLensRefresh(ctx context.Context) (err error) {
	return nil
}

func (server *AgnosticServer) PrepareCallHierarchy(ctx context.Context, params *protocol.CallHierarchyPrepareParams) (result []protocol.CallHierarchyItem, err error) {
	return nil, nil
}

func (server *AgnosticServer) IncomingCalls(ctx context.Context, params *protocol.CallHierarchyIncomingCallsParams) (result []protocol.CallHierarchyIncomingCall, err error) {
	return nil, nil
}

func (server *AgnosticServer) OutgoingCalls(ctx context.Context, params *protocol.CallHierarchyOutgoingCallsParams) (result []protocol.CallHierarchyOutgoingCall, err error) {
	return nil, nil
}

func (server *AgnosticServer) SemanticTokensFull(ctx context.Context, params *protocol.SemanticTokensParams) (result *protocol.SemanticTokens, err error) {
	return nil, nil
}

func (server *AgnosticServer) SemanticTokensFullDelta(ctx context.Context, params *protocol.SemanticTokensDeltaParams) (result interface{} /* SemanticTokens | SemanticTokensDelta */, err error) {
	return nil, nil
}

func (server *AgnosticServer) SemanticTokensRange(ctx context.Context, params *protocol.SemanticTokensRangeParams) (result *protocol.SemanticTokens, err error) {
	return nil, nil
}

func (server *AgnosticServer) SemanticTokensRefresh(ctx context.Context) (err error) {
	return nil
}

func (server *AgnosticServer) LinkedEditingRange(ctx context.Context, params *protocol.LinkedEditingRangeParams) (result *protocol.LinkedEditingRanges, err error) {
	return nil, nil
}

func (server *AgnosticServer) Moniker(ctx context.Context, params *protocol.MonikerParams) (result []protocol.Moniker, err error) {
	return nil, nil
}

func (server *AgnosticServer) Request(ctx context.Context, method string, params interface{}) (result interface{}, err error) {
	return nil, nil
}

func main() {
	fmt.Println("Hello, world!")
}