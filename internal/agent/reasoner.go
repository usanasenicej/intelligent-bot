package agent

import (
	"context"
	"fmt"
	"github.com/user/go-ai-agent/pkg/ai"
)

type Agent struct {
	provider   ai.Provider
	memory     ai.Memory
	tools      map[string]ai.Tool
	maxRetries int
	logger     ai.Logger
}

func NewAgent(p ai.Provider, m ai.Memory) *Agent {
	return &Agent{
		provider:   p,
		memory:     m,
		tools:      make(map[string]ai.Tool),
		maxRetries: 5,
		logger:     ai.DefaultLogger{},
	}
}

func (a *Agent) RegisterTool(t ai.Tool) {
	a.tools[t.Name()] = t
}

func (a *Agent) Run(ctx context.Context, input string) (string, error) {
	a.memory.Add(ai.Message{Role: ai.RoleUser, Content: input})

	for i := 0; i < a.maxRetries; i++ {
		a.logger.Log(fmt.Sprintf("Memory Size: %d messages", a.memory.Size()))
		messages := a.memory.Get()
		resp, err := a.provider.Generate(ctx, messages)
		if err != nil {
			a.logger.Errorf("Provider generation failed: %v", err)
			return "", err
		}

		a.logger.Log(fmt.Sprintf("Agent Thought: %s", resp.Message.Content))
		a.memory.Add(resp.Message)

		// In a real agent, we would parse 'resp.Message' for tool calls.
		// For this complex mock, we assume 'isFinalAnswer' handles termination.
		if isFinalAnswer(resp.Message.Content) {
			return resp.Message.Content, nil
		}
	}

	return "", fmt.Errorf("exceeded max retries of %d", a.maxRetries)
}

func isFinalAnswer(content string) bool {
	// Dummy logic for a complex agent
	return true 
}
