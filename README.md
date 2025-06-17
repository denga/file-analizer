# File Analyzer

A command-line tool written in Go that analyzes text files and provides various statistics about their content.

The main task of this project is to try out github actions.

## Features

- Character count (including newlines)
- Word count
- Line count
- Word frequency analysis (Top 10 most frequent words)
- Special character handling (removes special characters from words)

## Requirements

- Go 1.24 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/denga/file-analyzer.git
cd file-analyzer
```

2. Build the program:
```bash
go build
```

## Usage

Run the program with a text file as an argument:

```bash
./file-analyzer <path-to-file>
```

### Example

```bash
./file-analyzer sample.txt
```

This will output something like:

```
File Analysis for: sample.txt
=======================
Number of Characters: 1234
Number of Words:      200
Number of Lines:      15

Top 10 Most Frequent Words:
=========================
 1. the                 24
 2. and                 18
 3. in                  15
 4. to                  12
 5. of                  10
 6. a                    9
 7. is                   8
 8. that                 7
 9. for                  6
10. with                 5
```

## Word Processing

The analyzer processes words in the following way:
- Converts all text to lowercase
- Removes special characters, keeping only letters and numbers
- Combines similar words (e.g., "hello!" and "hello" are counted as the same word)
- Empty strings after cleaning are ignored in the word count

Examples of word cleaning:
- "Hello!" → "hello"
- "good-bye" → "goodbye"
- "café" → "cafe"
- "user123" → "user123"
- "!@#$%" → "" (ignored in counting)

## License

This project is open source and available under the MIT License. 