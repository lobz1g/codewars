package src

import (
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "Alex", want: "Hello, Alex how are you doing today?"},
		{name: "Ryan", want: "Hello, Ryan how are you doing today?"},
		{name: "Ololo", want: "Hello, Ololo how are you doing today?"},
		{name: "Golanger", want: "Hello, Golanger how are you doing today?"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.name); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
