package tools

import (
	"context"
	"fmt"
	"github.com/user/go-ai-agent/pkg/ai"
)

type Calculator struct{}

func (c *Calculator) Name() string { return "calculator" }

func (c *Calculator) Description() string {
	return "Evaluates simple math expressions."
}

func (c *Calculator) Parameters() ai.MetaData {
	return ai.MetaData{
		Type: "object",
		Properties: map[string]interface{}{
			"expression": map[string]string{"type": "string"},
		},
		Required: []string{"expression"},
	}
}

func (c *Calculator) Execute(ctx context.Context, args map[string]interface{}) (string, error) {
	expr, ok := args["expression"].(string)
	if !ok {
		return "", fmt.Errorf("invalid expression")
	}
	// Logic to evaluate expr. (Real-world would use a parser)
	return fmt.Sprintf("Calculated: %v", expr), nil
}
