package main

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeYAML_TwoFiles(t *testing.T) {
	// Create temporary files for testing
	dir, err := ioutil.TempDir("", "yaml-merge-test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(dir)

	file1 := filepath.Join(dir, "file1.yaml")
	file2 := filepath.Join(dir, "file2.yaml")

	// Write test data to files
	err = ioutil.WriteFile(file1, []byte("key1: value1\nkey2: value2"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file1.yaml: %v", err)
	}
	err = ioutil.WriteFile(file2, []byte("key2: overridden\nkey3: added"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file2.yaml: %v", err)
	}

	// Test case: merge two files
	ctx := context.Background()
	result, err := MergeYAML(ctx, file1, file2)
	expected := "key1: value1\nkey2: overridden\nkey3: added\n"
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, expected, result, "Unexpected result")
}

func TestMergeYAML_ThreeFiles(t *testing.T) {
	// Create temporary files for testing
	dir, err := ioutil.TempDir("", "yaml-merge-test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(dir)

	file1 := filepath.Join(dir, "file1.yaml")
	file2 := filepath.Join(dir, "file2.yaml")
	file3 := filepath.Join(dir, "file3.yaml")

	// Write test data to files
	err = ioutil.WriteFile(file1, []byte("key1: value1\nkey2: value2"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file1.yaml: %v", err)
	}
	err = ioutil.WriteFile(file2, []byte("key2: overridden\nkey3: added"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file2.yaml: %v", err)
	}
	err = ioutil.WriteFile(file3, []byte("key1:\n  nested_key: nested_value\nkey4: new_value"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file3.yaml: %v", err)
	}

	// Test case: merge three files
	ctx := context.Background()
	result, err := MergeYAML(ctx, file1, file2, file3)
	expected := "key1:\n  nested_key: nested_value\nkey2: overridden\nkey3: added\nkey4: new_value\n"
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, expected, result, "Unexpected result")
}
