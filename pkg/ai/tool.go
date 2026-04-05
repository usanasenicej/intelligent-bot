package ai

import "context"

// Tool describes a capability the AI Agent can use.
type Tool interface {
	Name() string
	Description() string
	Parameters() MetaData // JSON Schema describing parameters
	Execute(ctx context.Context, args map[string]interface{}) (string, error)
}

// MetaData is a simple container for JSON Schema-like structure.
type MetaData struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Required   []string               `json:"required"`
}
