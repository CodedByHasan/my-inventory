package utils

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var TestFindDotEnvTest = []struct {
	name          string
	setupEnvFile  bool
	expectedError bool
}{
	{
		name:          "Env file Found",
		setupEnvFile:  true,
		expectedError: false,
	},
	{
		name:          "Env file missing",
		setupEnvFile:  false,
		expectedError: true,
	},
}
func TestFindDotEnv(t *testing.T) {

	for _, tt := range TestFindDotEnvTest {
		baseDir := t.TempDir()

		if tt.setupEnvFile {
			envFilePath := filepath.Join(baseDir, ".env")
			content := []byte("TEST_KEY=test_value")

			if err := os.WriteFile(envFilePath, content, 0644); err != nil {
				t.Fatalf("Failed to write temp .env file: %v", err)
			}
		}

		subDir := filepath.Join(baseDir, "subdir", "nested")
		if err := os.MkdirAll(subDir, 0755); err != nil {
			t.Fatalf("Failed to create nested directory: %v", err)
		}

		if err := os.Chdir(subDir); err != nil {
			t.Fatalf("Failed to change working directory %v", err)
		}

		path, err := findDotEnv()

		if tt.expectedError {
			if err == nil {
				t.Error("Expected error, but got none")
			}
		} else {
			if err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
			if !strings.HasSuffix(path, ".env") {
				t.Errorf("Expected .env path, but got %s", path)
			}
		}
	}
}
