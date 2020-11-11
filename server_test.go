package main

import "testing"

func TestGreetingSpecific(t *testing.T) {
	greeting := CreateGreeting("George")
	if greeting != "Hello, George\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, George\n")
	}
}

func TestGreetingDefault(t *testing.T) {
	greeting := CreateGreeting("")
	if greeting != "Hello, Guest\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Guest\n")
	}
}
