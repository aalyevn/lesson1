package semver

import (
  "testing"
)

func TestParseSemVer(t *testing.T) {
  tests := []struct {
    input    string
    expected SemVer
    hasError bool
  }{
    {"1.2.3", SemVer{1, 2, 3}, false},
    {"10.20.30", SemVer{10, 20, 30}, false},
    {"1.2", SemVer{}, true},
    {"a.b.c", SemVer{}, true},
  }

  for _, test := range tests {
    result, err := ParseSemVer(test.input)
    if err != nil && !test.hasError {
      t.Errorf("TestParseSemVer with input %s: unexpected error: %v", test.input, err)
      continue
    }
    if err == nil && test.hasError {
      t.Errorf("TestParseSemVer with input %s: expected error but got none", test.input)
      continue
    }
    if result != test.expected {
      t.Errorf("TestParseSemVer with input %s: expected %v but got %v", test.input, test.expected, result)
    }
  }
}

func TestSemVerIncrementBy(t *testing.T) {
  tests := []struct {
    initial  SemVer
    step     SemVer
    expected SemVer
  }{
    {SemVer{1, 2, 3}, SemVer{0, 0, 1}, SemVer{1, 2, 4}},
    {SemVer{1, 2, 3}, SemVer{1, 0, 2}, SemVer{2, 2, 5}},
  }

  for _, test := range tests {
    test.initial.IncrementBy(test.step)
    if test.initial != test.expected {
      t.Errorf("IncrementBy: from %v with step %v: expected %v but got %v", test.initial, test.step, test.expected, test.initial)
    }
  }
}

func TestCompileSemverIdentificationRegexpPattern(t *testing.T) {
  tests := []struct {
    prefix          string
    expectedPattern string
  }{
    {"^version:", `^version:.*(\d+\.\d+\.\d+)`},
    {"prefix-", `prefix-.*(\d+\.\d+\.\d+)`},
    {"", `.*(\d+\.\d+\.\d+)`},
  }

  for _, tt := range tests {
    t.Run(tt.prefix, func(t *testing.T) {
      got := CompileSemverIdentificationRegexpPattern(tt.prefix)
      if got != tt.expectedPattern {
        t.Errorf("CompileSemverIdentificationRegexpPattern(%s) = %s, want %s", tt.prefix, got, tt.expectedPattern)
      }
    })
  }
}

func TestProcessFileContent(t *testing.T) {
  tests := []struct {
    content  string
    increase string
    prefix   string
    expected string
  }{
    {
      content:  "version: 1.2.3",
      increase: "0.0.1",
      prefix:   "^version:",
      expected: "version: 1.2.4",
    },
    {
      content:  "version: 1.2.3",
      increase: "1.0.2",
      prefix:   "^version:",
      expected: "version: 2.2.5",
    },
  }

  for _, test := range tests {
    output, err := ProcessFileContent(test.content, test.increase, test.prefix)
    if err != nil {
      t.Errorf("TestProcessFileContent with content %s: unexpected error: %v", test.content, err)
    }
    if output != test.expected {
      t.Errorf("TestProcessFileContent with content %s: expected %s, got %s", test.content, test.expected, output)
    }
  }
}

// ... You can add more tests for other functions as necessary ...
