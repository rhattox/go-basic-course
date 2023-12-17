package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	// Use the context in a function
	result, err := performOperation(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}

func performOperation(ctx context.Context) (string, error) {
	// Simulate a time-consuming operation
	select {
	case <-time.After(3 * time.Second):
		return "Operation completed successfully", nil
	case <-ctx.Done():
		// Context is canceled or timed out
		return "", ctx.Err()
	}
}
