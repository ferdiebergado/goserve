package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestServeIndexHTMLWithFlags(t *testing.T) {
	// Create a temporary directory to simulate the current directory
	tempDir, err := os.MkdirTemp("", "testserver")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create an index.html file in the temporary directory
	indexPath := filepath.Join(tempDir, "index.html")
	indexContent := []byte("<html><body><h1>Hello, World!</h1></body></html>")
	if err := os.WriteFile(indexPath, indexContent, 0644); err != nil {
		t.Fatalf("failed to write index.html: %v", err)
	}

	// Simulate running the program with flags
	cmd := exec.Command(
		"go", "run", "main.go",
		"-a", "localhost",
		"-p", "8080",
		tempDir,
	)
	cmd.Env = append(os.Environ(), "GOSERVE_TEST=1")

	if err := cmd.Run(); err != nil {
		t.Errorf("cmd run: %v", err)
	}

	// Make a GET request to the test server
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatalf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK, got %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if string(body) != string(indexContent) {
		t.Errorf("expected response body to match index.html content, got %s", string(body))
	}
}
