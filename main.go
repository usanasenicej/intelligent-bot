package main

import (
	"context"
	"fmt"
	"github.com/user/go-ai-agent/internal/agent"
	"github.com/user/go-ai-agent/internal/memory"
	"github.com/user/go-ai-agent/pkg/ai"
)

type MockProvider struct{}

func (m MockProvider) Name() string { return "mock-llm" }

func (m MockProvider) Generate(ctx context.Context, messages []ai.Message) (ai.Response, error) {
	// A real LLM would generate this based on its context.
	lastMsg := messages[len(messages)-1].Content
	return ai.Response{
		Message: ai.Message{
			Role:    ai.RoleAssistant,
			Content: fmt.Sprintf("I received: %s. Processing it intelligently now...", lastMsg),
		},
		Usage: ai.TokenUsage{PromptTokens: 10, CompletionTokens: 5, TotalTokens: 15},
	}, nil
}

func main() {
	ctx := context.Background()
	provider := MockProvider{}
	mem := memory.NewHistoryMemory()
	
	bot := agent.NewAgent(provider, mem)
	
	fmt.Println("AI Bot is running...")
	response, err := bot.Run(ctx, "Calculate 1 + 1")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Final Bot Response: %s\n", response)
}
