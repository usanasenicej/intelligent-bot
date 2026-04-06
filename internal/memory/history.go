package memory

import (
	"github.com/user/go-ai-agent/pkg/ai"
)

type HistoryMemory struct {
	messages []ai.Message
}

func NewHistoryMemory() *HistoryMemory {
	return &HistoryMemory{messages: make([]ai.Message, 0)}
}

func (h *HistoryMemory) Add(msg ai.Message) {
	h.messages = append(h.messages, msg)
}

func (h *HistoryMemory) Clear() {
	h.messages = make([]ai.Message, 0)
}

func (h *HistoryMemory) Get() []ai.Message {
	return h.messages
}

func (h *HistoryMemory) Size() int {
	return len(h.messages)
}
