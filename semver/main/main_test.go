package main

import (
  "io/ioutil"
  "os"
  "testing"
)

func TestProcessPath(t *testing.T) {
  tmpDir, err := ioutil.TempDir("", "testdir")
  if err != nil {
    t.Fatalf("could not create temp dir: %v", err)
  }
  defer os.RemoveAll(tmpDir)

  filePath := tmpDir + "/Chart.yaml"
  content := []byte("version: 1.2.3")
  err = ioutil.WriteFile(filePath, content, 0644)
  if err != nil {
    t.Fatalf("could not write temp file: %v", err)
  }

  // Test
  err = processPath(filePath, "Chart.yaml", "0.0.1", "^version:")
  if err != nil {
    t.Errorf("unexpected error in processPath: %v", err)
  }

  updatedContent, err := os.ReadFile(filePath)
  if err != nil {
    t.Fatalf("could not read updated temp file: %v", err)
  }

  expected := "version: 1.2.4"
  if string(updatedContent) != expected {
    t.Errorf("expected file content %q, but got %q", expected, updatedContent)
  }
}
