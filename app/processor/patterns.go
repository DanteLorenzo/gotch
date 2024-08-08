package processor

import (
	"encoding/hex"

	"github.com/fatih/color" // Import the color package for colored output
)

// PatternToBytes converts a pattern string (e.g., "?? 48 89 ??") into a byte slice.
// If a pattern contains "??", it is replaced with a 0 byte in the resulting slice.
func PatternToBytes(pattern string) []byte {
	var patternBytes []byte
	for i := 0; i < len(pattern); i += 3 { // Process each byte in the pattern (separated by spaces)
		if pattern[i] == '?' && pattern[i+1] == '?' {
			patternBytes = append(patternBytes, 0) // Append a 0 byte for "??"
		} else {
			b, _ := hex.DecodeString(pattern[i : i+2]) // Convert the hex string to a byte
			patternBytes = append(patternBytes, b[0])  // Append the converted byte
		}
	}
	return patternBytes
}

// SearchPatterns searches for all occurrences of a given pattern in the data.
// It returns a slice of positions where the pattern matches.
func SearchPatterns(data, pattern []byte) []int {
	patternLen := len(pattern)
	positions := []int{}
	for i := 0; i <= len(data)-patternLen; i++ {
		match := true
		for j := 0; j < patternLen; j++ {
			if pattern[j] != 0 && data[i+j] != pattern[j] { // Check for a match considering wildcards
				match = false
				break
			}
		}
		if match {
			positions = append(positions, i) // Store the position of the match
		}
	}
	return positions
}

// ReplacePlaceholders replaces the placeholder bytes (0) in the pattern with corresponding bytes
// from the original opcode, preserving the rest of the pattern.
func ReplacePlaceholders(pattern, originalOpcode []byte) []byte {
	var result []byte
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == 0 {
			result = append(result, originalOpcode[i]) // Replace 0 with the original byte
		} else {
			result = append(result, pattern[i]) // Keep the existing byte from the pattern
		}
	}
	return result
}

// PrintPatterns searches for and prints the positions of a pattern in the data.
// It provides color-coded output to indicate whether the pattern was found.
func PrintPatterns(data []byte, pattern string, info *color.Color, success *color.Color, errorColor *color.Color) {
	patternBytes := PatternToBytes(pattern)  // Convert the pattern to bytes
	positions := SearchPatterns(data, patternBytes) // Search for the pattern in the data

	if len(positions) > 0 {
		for _, pos := range positions {
			info.Printf("[INFO] Pattern ")
			success.Printf("%s", pattern)
			info.Printf(" found at position ")
			success.Printf("%d\n", pos) // Print the position where the pattern was found
			info.Printf("Opcode at position %d: ", pos)
			success.Printf("%s\n", hex.EncodeToString(data[pos:pos+len(patternBytes)])) // Print the matched opcode
		}
	} else {
		info.Printf("[INFO] Pattern ")
		errorColor.Printf("%s", pattern)
		info.Println(" not found") // Indicate that the pattern was not found
	}
}
