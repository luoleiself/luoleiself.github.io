// server.js
import express from 'express';
import cors from 'cors';
import bodyParser from 'body-parser';

const app = express();
const PORT = process.env.PORT || 3000;

app.use(cors());
app.use(bodyParser.json({ limit: '10mb' }));
app.use(bodyParser.urlencoded({ extended: true }));

// In-memory storage for conversations
const conversations = new Map();

// Available tools
const availableTools = {
  calculator: {
    name: "calculator",
    description: "A simple calculator tool for mathematical operations",
    parameters: {
      type: "object",
      properties: {
        expression: {
          type: "string",
          description: "The mathematical expression to evaluate"
        }
      },
      required: ["expression"]
    }
  },
  getCurrentTime: {
    name: "getCurrentTime",
    description: "Get the current time and date",
    parameters: {
      type: "object",
      properties: {}
    }
  }
};

// MCP Endpoints
app.post('/mcp/init', (req, res) => {
  const sessionId = `session_${Date.now()}`;

  res.json({
    session_id: sessionId,
    status: "ready",
    capabilities: {
      chat_completion: true,
      tools: Object.keys(availableTools),
      streaming: false
    },
    models: ["mcp-local-model-v1", "mcp-local-model-v2"]
  });
});

app.post('/mcp/chat/completions', async (req, res) => {
  const {
    session_id,
    messages,
    model,
    tools,
    stream = false
  } = req.body;

  try {
    // Store conversation context
    if (session_id) {
      if (!conversations.has(session_id)) {
        conversations.set(session_id, []);
      }
      conversations.get(session_id).push(...messages);
    }

    // Process the request
    const response = await generateResponse(messages, tools);

    if (stream) {
      // For streaming responses
      res.setHeader('Content-Type', 'text/event-stream');
      res.setHeader('Cache-Control', 'no-cache');
      res.setHeader('Connection', 'keep-alive');

      res.write(`data: ${JSON.stringify(response)}\n\n`);
      res.end();
    } else {
      res.json(response);
    }
  } catch (error) {
    res.status(500).json({
      error: {
        message: error.message,
        type: "server_error"
      }
    });
  }
});

app.post('/mcp/tools/call', (req, res) => {
  const { tool_calls } = req.body;

  try {
    const results = tool_calls.map(call => {
      try {
        const result = executeTool(call.name, call.arguments);
        return {
          id: call.id,
          name: call.name,
          result: result
        };
      } catch (error) {
        return {
          id: call.id,
          name: call.name,
          error: error.message
        };
      }
    });

    res.json({ tool_results: results });
  } catch (error) {
    res.status(500).json({
      error: {
        message: error.message,
        type: "tool_execution_error"
      }
    });
  }
});

app.get('/mcp/tools', (req, res) => {
  res.json({
    tools: Object.values(availableTools)
  });
});

// Helper functions
async function generateResponse(messages, tools = []) {
  // Simulate AI response generation
  const lastMessage = messages[messages.length - 1];

  // Check if we should use a tool
  if (tools && tools.length > 0 && Math.random() > 0.7) {
    // Randomly decide to use a tool
    const toolName = Object.keys(availableTools)[Math.floor(Math.random() * Object.keys(availableTools).length)];
    const tool = availableTools[toolName];

    return {
      id: `chatcmpl-${Date.now()}`,
      object: "chat.completion",
      created: Math.floor(Date.now() / 1000),
      model: "mcp-local-model",
      choices: [{
        index: 0,
        message: {
          role: "assistant",
          content: null,
          tool_calls: [{
            id: `call_${Date.now()}`,
            type: "function",
            function: {
              name: tool.name,
              arguments: JSON.stringify({
                // Generate appropriate arguments based on tool
                expression: "2+2",
                ...(["getCurrentTime"].includes(tool.name) ? {} : {})
              })
            }
          }]
        },
        finish_reason: "tool_calls"
      }]
    };
  }

  // Regular text response
  return {
    id: `chatcmpl-${Date.now()}`,
    object: "chat.completion",
    created: Math.floor(Date.now() / 1000),
    model: "mcp-local-model",
    choices: [{
      index: 0,
      message: {
        role: "assistant",
        content: `I received your message: "${lastMessage.content}". This is a response from the local MCP server.`
      },
      finish_reason: "stop"
    }],
    usage: {
      prompt_tokens: messages.reduce((sum, msg) => sum + (msg.content?.length || 0), 0),
      completion_tokens: 30,
      total_tokens: messages.reduce((sum, msg) => sum + (msg.content?.length || 0), 0) + 30
    }
  };
}

function executeTool(toolName, args) {
  switch (toolName) {
    case 'calculator':
      try {
        // Simple expression evaluation (in real implementation, use a proper math library)
        const result = eval(args.expression); // Note: eval can be dangerous in production
        return { result: result };
      } catch (error) {
        throw new Error(`Calculation error: ${error.message}`);
      }

    case 'getCurrentTime':
      return {
        currentTime: new Date().toISOString(),
        timezone: Intl.DateTimeFormat().resolvedOptions().timeZone
      };

    default:
      throw new Error(`Unknown tool: ${toolName}`);
  }
}

// Error handling middleware
app.use((err, req, res, next) => {
  console.error(err.stack);
  res.status(500).json({
    error: {
      message: "Internal server error",
      type: "internal_error"
    }
  });
});

// 404 handler
app.use((req, res) => {
  res.status(404).json({
    error: {
      message: "Endpoint not found",
      type: "not_found"
    }
  });
});

app.listen(PORT, () => {
  console.log(`Enhanced MCP Server running on port ${PORT}`);
  console.log(`Endpoints:`);
  console.log(`  - POST /mcp/init`);
  console.log(`  - POST /mcp/chat/completions`);
  console.log(`  - POST /mcp/tools/call`);
  console.log(`  - GET  /mcp/tools`);
});
