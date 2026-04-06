package ai

import "context"

// Provider is an interface for LLM models like Gemini, OpenAI, etc.
type Provider interface {
	Generate(ctx context.Context, messages []Message) (Response, error)
	Name() string
}

// Memory is an interface for storing context between turns.
type Memory interface {
	Add(message Message)
	Clear()
	Get() []Message
	Size() int
}
