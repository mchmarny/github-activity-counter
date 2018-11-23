package handler

import (
	"testing"
)

func TestFileContentReader(t *testing.T) {
	const testDataPath = "../samples/pull_request.json"
	_, err := getFileContent(testDataPath)
	if err != nil {
		t.Errorf("Error while getting test file %s: %v", testDataPath, err)
	}
}

func TestFileContentReaderNotFoundError(t *testing.T) {
	const testDataPath = "../samples/file-not-exists.json"
	_, err := getFileContent(testDataPath)
	if err == nil {
		t.Errorf("Expected file not found error %s: %v", testDataPath, err)
	}
}
