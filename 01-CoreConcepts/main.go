package main

import (
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

// Tool arguments are just structs, annotated with jsonschema tags
// More at https://mcpgolang.com/tools#schema-generation
type WeatherData struct {
	Temperature float32 `json:"temperature"`
	Conditions  string  `json:"conditions"`
	Location    string  `json:"location"`
}

type WeatherArgs struct {
	Location string `json:"location" jsonschema:"required"`
}

// run: npx @modelcontextprotocol/inspector
// command: <executable-path>
func main() {
	done := make(chan struct{})

	transport := stdio.NewStdioServerTransport()
	server := mcp.NewServer(
		transport,
		mcp.WithName("Weather MCP Server"),
		mcp.WithVersion("1.0.0"),
	)

	err := server.RegisterTool(
		"weather tool",
		"get current weather for a given location",
		func(args WeatherArgs) (*mcp.ToolResponse, error) {
			return mcp.NewToolResponse(mcp.NewTextContent(
				fmt.Sprintf("%v", WeatherData{
					Temperature: 75.5,
					Conditions:  "sunny",
					Location:    args.Location,
				}),
			)), nil
		})

	if err != nil {
		panic(err)
	}

	err = server.Serve()
	if err != nil {
		panic(err)
	}

	<-done
}
