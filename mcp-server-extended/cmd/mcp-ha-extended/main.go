// Command mcp-ha-extended serves Home Assistant automation management over the
// Model Context Protocol, speaking MCP on stdin/stdout.
package main

import (
	"context"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	logger "github.com/sirupsen/logrus"
)

// version is the add-on version, overridden at build time with
// -ldflags "-X main.version=<config.yaml version>".
var version = "dev"

func main() {
	// stdout carries the MCP protocol stream, so every log line must go to
	// stderr — anything else written to stdout corrupts the session.
	logger.SetOutput(os.Stderr)

	controller, err := injectController()
	if err != nil {
		logger.WithFields(logger.Fields{"error": err}).Fatal("failed to start")
	}

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "home-assistant-automations",
		Version: version,
	}, nil)
	controller.Register(server)

	logger.WithFields(logger.Fields{"version": version}).Info("serving MCP over stdio")

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		logger.WithFields(logger.Fields{"error": err}).Fatal("server stopped")
	}
}
