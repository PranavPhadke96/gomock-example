package greeting

import (
	"testing"
	"github.com/golang/mock/gomock"
)

func TestGreetingService_CreateGreeting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGreeter := NewMockGreeter(ctrl)
	mockGreeter.EXPECT().Greet("Alice").Return("Hello, Alice!")

	service := NewGreetingService(mockGreeter)
	greeting := service.CreateGreeting("Alice")

	if greeting != "Hello, Alice!" {
		t.Errorf("Expected 'Hello, Alice!', got '%s'", greeting)
	}
}