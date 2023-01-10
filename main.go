package main

import (
	"context"
	"fmt"
	"io"
	"os"

	agnosticLSP "github.com/danielway/agnostic-lsp/lsp"
	"github.com/pulumi/pulumi-lsp/sdk/lsp"
)

func main() {
	server := lsp.NewServer(agnosticLSP.Methods(), &stdio{false})

	fmt.Println(`Server starting`)
	err := server.Run(context.Background())
	if err != nil {
		fmt.Println("Err: ", err.Error())
	}

	fmt.Println(`Server stopped`)
}

type stdio struct{ bool }

func (s *stdio) Read(p []byte) (n int, err error) {
	if s.bool {
		return 0, io.EOF
	}
	return os.Stdin.Read(p)
}

func (s *stdio) Write(p []byte) (n int, err error) {
	if s.bool {
		return 0, io.EOF
	}
	return os.Stdout.Write(p)
}

func (s *stdio) Close() error {
	s.bool = true
	return nil
}
