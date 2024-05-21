package main

import "testing"

func TestClient(t *testing.T) {
	client := NewClient()
	if err := client.Connect(":8081"); err != nil {
		t.Fatalf("Error in Connect: %v", err)
	}

	client.Start()
}
