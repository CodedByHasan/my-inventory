package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindDotEnv(t *testing.T) {
	baseDir := t.TempDir()

	envFilePath := filepath.Join(baseDir, ".env")
	content := []byte("TEST_KEY=test_value")

	if err := os.WriteFile(envFilePath, content, 0644); err != nil{
		t.Fatalf("Failed to write temp .env file: %v", err)
	}

	subDir := filepath.Join(baseDir, "subdir", "nested")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("Failed to create nested directory: %v", err)
	}

	if err := os.Chdir(subDir); err != nil {
		t.Fatalf("Failed to change working directory %v", err)
	}

	// foundPath := "test"
	foundPath := findDotEnv()

	if foundPath != envFilePath {
		t.Errorf("Expected findDotEnv to find %s, but got %s", envFilePath, foundPath)
	}
}
