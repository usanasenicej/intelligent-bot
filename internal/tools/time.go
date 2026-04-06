package tools

import (
	"context"
	"time"
	"github.com/user/go-ai-agent/pkg/ai"
)

type TimeTool struct{}

func (t *TimeTool) Name() string { return "get_current_time" }

func (t *TimeTool) Description() string {
	return "Returns the current system time."
}

func (t *TimeTool) Parameters() ai.MetaData {
	return ai.MetaData{
		Type:       "object",
		Properties: make(map[string]interface{}),
		Required:   []string{},
	}
}

func (t *TimeTool) Execute(ctx context.Context, args map[string]interface{}) (string, error) {
	return time.Now().Format(time.RFC3339), nil
}
