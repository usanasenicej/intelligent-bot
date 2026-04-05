package agent

import (
	"context"
	"fmt"
	"github.com/user/go-ai-agent/pkg/ai"
)

type Agent struct {
	provider ai.Provider
	memory   ai.Memory
	tools    map[string]ai.Tool
	maxRetries int
}

func NewAgent(p ai.Provider, m ai.Memory) *Agent {
	return &Agent{
		provider: p,
		memory:   m,
		tools:    make(map[string]ai.Tool),
		maxRetries: 5,
	}
}

func (a *Agent) RegisterTool(t ai.Tool) {
	a.tools[t.Name()] = t
}

func (a *Agent) Run(ctx context.Context, input string) (string, error) {
	a.memory.Add(ai.Message{Role: ai.RoleUser, Content: input})

	for i := 0; i < a.maxRetries; i++ {
		messages := a.memory.Get()
		resp, err := a.provider.Generate(ctx, messages)
		if err != nil {
			return "", err
		}

		// Simplified: The provider's response could include a tool call.
		// For this complex mock, we'll imagine it's an intelligent loop.
		// The agent decides when it has a final answer.
		
		fmt.Printf("[Agent Thought]: %s\n", resp.Message.Content)
		a.memory.Add(resp.Message)

		// Check if we reached a final conclusion.
		// In a real scenario, the LLM would signal this via a specific output format.
		if isFinalAnswer(resp.Message.Content) {
			return resp.Message.Content, nil
		}
	}

	return "", fmt.Errorf("exceeded max retries")
}

func isFinalAnswer(content string) bool {
	// Dummy logic for a complex agent
	return true 
}
