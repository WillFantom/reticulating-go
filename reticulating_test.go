package reticulating

import "testing"

func BenchmarkLoadingMessages(b *testing.B) {
	for i := 0; i < 500; i++ {
		GetLoadingMessage()
	}
}

func TestLoadingMessage(t *testing.T) {
	var messages []string
	for i := 0; i < 5; i++ {
		messages = append(messages, GetLoadingMessage())
	}
	t.Logf("Random Messages: %v\n", messages)
}
