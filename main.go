package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type FileStats struct {
	CharCount     int
	WordCount     int
	LineCount     int
	WordFrequency map[string]int
}

// cleanWord removes special characters and converts to lowercase
func cleanWord(word string) string {
	// Convert to lowercase first
	word = strings.ToLower(word)

	// Keep letters (including unicode letters like umlauts), numbers and combine them
	var result strings.Builder
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			result.WriteRune(r)
		}
	}

	return result.String()
}

func analyzeFile(filePath string) (*FileStats, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	stats := &FileStats{
		WordFrequency: make(map[string]int),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stats.LineCount++
		stats.CharCount += len(line) + 1 // +1 for newline character

		// Split line into words and count
		words := strings.Fields(line)
		stats.WordCount += len(words)

		// Count word frequency
		for _, word := range words {
			// Clean the word before counting
			cleaned := cleanWord(word)
			if cleaned != "" { // Only count non-empty words
				stats.WordFrequency[cleaned]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return stats, nil
}

type WordCount struct {
	Word  string
	Count int
}

func getTop10Words(wordFreq map[string]int) []WordCount {
	var pairs []WordCount
	for word, count := range wordFreq {
		pairs = append(pairs, WordCount{word, count})
	}

	// Sort by count in descending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	// Get top 10 or less if there are fewer words
	n := 10
	if len(pairs) < 10 {
		n = len(pairs)
	}
	return pairs[:n]
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: file-analyzer <filename>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	stats, err := analyzeFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nFile Analysis for: %s\n", filePath)
	fmt.Printf("===================%s\n", strings.Repeat("=", len(filePath)))
	fmt.Printf("Number of Characters: %d\n", stats.CharCount)
	fmt.Printf("Number of Words:      %d\n", stats.WordCount)
	fmt.Printf("Number of Lines:      %d\n", stats.LineCount)

	fmt.Printf("\nTop 10 Most Frequent Words:\n")
	fmt.Println("=========================")
	for i, wc := range getTop10Words(stats.WordFrequency) {
		fmt.Printf("%2d. %-20s %d\n", i+1, wc.Word, wc.Count)
	}
}
