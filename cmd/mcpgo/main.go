package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/guigui42/mcpgo/pkg/models"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	serverPort = 8080
)

func main() {
	// Set up logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Initializing MCP server...")

	// Create a new hooks instance
	hooks := &server.Hooks{}

	// Add hooks using the proper methods
	hooks.AddBeforeAny(func(id any, method mcp.MCPMethod, message any) {
		log.Printf("REQUEST [%v] Method: %s\n", id, method)
	})

	hooks.AddOnSuccess(func(id any, method mcp.MCPMethod, message any, result any) {
		log.Printf("SUCCESS [%v] Method: %s\n", id, method)
	})

	hooks.AddOnError(func(id any, method mcp.MCPMethod, message any, err error) {
		log.Printf("ERROR [%v] Method: %s, Error: %v\n", id, method, err)
	})

	hooks.AddAfterCallTool(func(id any, message *mcp.CallToolRequest, result *mcp.CallToolResult) {
		log.Printf("TOOL COMPLETED [%v] Tool: %s\n", id, message.Params.Name)
	})

	// Create MCP server with hooks
	s := server.NewMCPServer(
		"MCP Demo Server",
		"1.0.0",
		server.WithInstructions("A Go MCP server implementation based on the octodemo/mcp-octodemo TypeScript example"),
		server.WithLogging(),
		server.WithHooks(hooks),
	)

	// Add get_services tool
	getServicesTool := mcp.NewTool("get_services",
		mcp.WithDescription("Retrieve all services"),
	)
	s.AddTool(getServicesTool, logWrapper("get_services", getServicesHandler))

	// Add get_service_by_id tool
	getServiceByIdTool := mcp.NewTool("get_service_by_id",
		mcp.WithDescription("Retrieve a specific service by its ID"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("ID of the service"),
		),
	)
	s.AddTool(getServiceByIdTool, logWrapper("get_service_by_id", getServiceByIdHandler))

	// Add get_services_by_owner tool
	getServicesByOwnerTool := mcp.NewTool("get_services_by_owner",
		mcp.WithDescription("Retrieve services owned by a specific person"),
		mcp.WithString("owner",
			mcp.Required(),
			mcp.Description("Name of the owner"),
		),
	)
	s.AddTool(getServicesByOwnerTool, logWrapper("get_services_by_owner", getServicesByOwnerHandler))

	// Start the HTTP server on port 8080
	log.Printf("Starting MCP server on port %d...\n", serverPort)

	// Create SSE server with HTTP support
	sseServer := server.NewSSEServer(s,
		server.WithBasePath("/mcp"),
		server.WithSSEEndpoint("/sse"),
		server.WithMessageEndpoint("/message"),
	)

	// Log server start
	log.Printf("Server endpoints:\n"+
		"- Base URL: http://localhost:%d/mcp\n"+
		"- SSE endpoint: http://localhost:%d/mcp/sse\n"+
		"- Message endpoint: http://localhost:%d/mcp/message\n",
		serverPort, serverPort, serverPort)

	// Start the server on port 8080 using the correct Start method
	if err := sseServer.Start(fmt.Sprintf(":%d", serverPort)); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}

// logWrapper wraps a handler function with logging
func logWrapper(toolName string, handler server.ToolHandlerFunc) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		startTime := time.Now()

		// Log the incoming request
		argJSON, _ := json.Marshal(request.Params.Arguments)
		log.Printf("TOOL INVOKED: %s, Arguments: %s", toolName, string(argJSON))

		// Call the original handler
		result, err := handler(ctx, request)

		// Log the outcome
		if err != nil {
			log.Printf("TOOL ERROR: %s, Error: %v, Duration: %v", toolName, err, time.Since(startTime))
		} else {
			resultJSON, _ := json.Marshal(result)
			log.Printf("TOOL SUCCESS: %s, Result: %s, Duration: %v", toolName, string(resultJSON), time.Since(startTime))
		}

		return result, err
	}
}

func getServicesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	services, err := models.GetAllServices()
	if err != nil {
		return nil, fmt.Errorf("error retrieving services: %v", err)
	}

	jsonData, err := json.Marshal(services)
	if err != nil {
		return nil, fmt.Errorf("error marshaling services data: %v", err)
	}

	return mcp.NewToolResultText(string(jsonData)), nil
}

func getServiceByIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	idVal, ok := request.Params.Arguments["id"]
	if !ok {
		return nil, errors.New("id is required")
	}

	// Handle different number formats
	var id int
	switch v := idVal.(type) {
	case float64:
		id = int(v)
	case int:
		id = v
	case string:
		var err error
		id, err = strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("id must be a number")
		}
	default:
		return nil, errors.New("id must be a number")
	}

	service, err := models.GetServiceByID(id)
	if err != nil {
		return nil, fmt.Errorf("error retrieving service: %v", err)
	}

	if service == nil {
		return mcp.NewToolResultText(`{"error": "Service not found"}`), nil
	}

	jsonData, err := json.Marshal(service)
	if err != nil {
		return nil, fmt.Errorf("error marshaling service data: %v", err)
	}

	return mcp.NewToolResultText(string(jsonData)), nil
}

func getServicesByOwnerHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	owner, ok := request.Params.Arguments["owner"].(string)
	if !ok {
		return nil, errors.New("owner must be a string")
	}

	services, err := models.GetServicesByOwner(owner)
	if err != nil {
		return nil, fmt.Errorf("error retrieving services: %v", err)
	}

	jsonData, err := json.Marshal(services)
	if err != nil {
		return nil, fmt.Errorf("error marshaling services data: %v", err)
	}

	return mcp.NewToolResultText(string(jsonData)), nil
}
