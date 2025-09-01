#!/usr/bin/env bash
# Initialize session
curl -X POST http://localhost:3000/mcp/init \
  -H "Content-Type: application/json"

# Simple chat completion
curl -X POST http://localhost:3000/mcp/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "messages": [
      {
        "role": "user",
        "content": "Hello, what can you do?"
      }
    ],
    "model": "mcp-local-model-v1"
  }'

# Chat with tools
curl -X POST http://localhost:3000/mcp/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "messages": [
      {
        "role": "user",
        "content": "What is 2+2?"
      }
    ],
    "tools": ["calculator"]
  }'

# Execute tools
curl -X POST http://localhost:3000/mcp/tools/call \
  -H "Content-Type: application/json" \
  -d '{
    "tool_calls": [
      {
        "id": "call_12345",
        "name": "calculator",
        "arguments": "{\"expression\": \"2+2\"}"
      }
    ]
  }'

# Get available tools
curl -X GET http://localhost:3000/mcp/tools
