package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCleanWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Basic word", "hello", "hello"},
		{"Word with punctuation", "hello!", "hello"},
		{"Mixed case", "Hello", "hello"},
		{"With numbers", "user123", "user123"},
		{"Special characters", "café", "café"},
		{"Hyphenated word", "good-bye", "goodbye"},
		{"Only special chars", "!@#$%", ""},
		{"Empty string", "", ""},
		{"Multiple special chars", "!!!hello???", "hello"},
		{"German umlauts", "über", "über"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cleanWord(tt.input)
			if result != tt.expected {
				t.Errorf("cleanWord(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetTop10Words(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []WordCount
	}{
		{
			name: "Less than 10 words",
			input: map[string]int{
				"hello": 3,
				"world": 2,
				"go":    1,
			},
			expected: []WordCount{
				{Word: "hello", Count: 3},
				{Word: "world", Count: 2},
				{Word: "go", Count: 1},
			},
		},
		{
			name: "More than 10 words",
			input: map[string]int{
				"a": 11, "b": 10, "c": 9, "d": 8, "e": 7,
				"f": 6, "g": 5, "h": 4, "i": 3, "j": 2,
				"k": 1,
			},
			expected: []WordCount{
				{Word: "a", Count: 11},
				{Word: "b", Count: 10},
				{Word: "c", Count: 9},
				{Word: "d", Count: 8},
				{Word: "e", Count: 7},
				{Word: "f", Count: 6},
				{Word: "g", Count: 5},
				{Word: "h", Count: 4},
				{Word: "i", Count: 3},
				{Word: "j", Count: 2},
			},
		},
		{
			name:     "Empty map",
			input:    map[string]int{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getTop10Words(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("getTop10Words() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAnalyzeFile(t *testing.T) {
	// Create a temporary test file
	content := `Hello, World!
This is a test file.
It has three lines and some repeated words.
test test test`

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	stats, err := analyzeFile(tmpFile)
	if err != nil {
		t.Fatalf("analyzeFile() error = %v", err)
	}

	// Test character count (including newlines)
	expectedCharCount := len(content) + 1 // +1 for the last newline
	if stats.CharCount != expectedCharCount {
		t.Errorf("CharCount = %d, want %d", stats.CharCount, expectedCharCount)
	}

	// Test line count
	expectedLineCount := 4
	if stats.LineCount != expectedLineCount {
		t.Errorf("LineCount = %d, want %d", stats.LineCount, expectedLineCount)
	}

	// Test word count (manually counted from the test content)
	expectedWordCount := 18 // Updated to match actual word count
	if stats.WordCount != expectedWordCount {
		t.Errorf("WordCount = %d, want %d", stats.WordCount, expectedWordCount)
	}

	// Test word frequency
	if count := stats.WordFrequency["test"]; count != 4 {
		t.Errorf("Frequency of 'test' = %d, want 4", count)
	}
}

func TestAnalyzeFileErrors(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
	}{
		{"Non-existent file", "nonexistent.txt"},
		{"Directory instead of file", "."}, // Current directory
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := analyzeFile(tt.filepath)
			if err == nil {
				t.Error("analyzeFile() error = nil, want error")
			}
		})
	}
}
