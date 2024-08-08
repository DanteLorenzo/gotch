package processor

import (
	"github.com/DanteLorenzo/gotch/app/types"
)

// ModifyPatterns updates the byte slice 'data' by replacing old patterns with new patterns.
// 'patterns' is a slice of 'types.Pattern' that contains old and new patterns to be replaced.
func ModifyPatterns(data []byte, patterns []types.Pattern) []byte {
	for _, pattern := range patterns {
		// Convert the new pattern from string format to byte slice format
		newPatternBytes := PatternToBytes(pattern.NewPattern)
		// Replace occurrences of the old pattern with the new pattern
		data = ReplacePattern(data, pattern.OldPattern, newPatternBytes)
	}
	return data
}

// ReplacePattern searches for occurrences of 'pattern' in 'data' and replaces them with 'newPattern'.
// 'pattern' is a string representing the pattern to be replaced, and 'newPattern' is the replacement in byte slice format.
func ReplacePattern(data []byte, pattern string, newPattern []byte) []byte {
	// Convert the old pattern from string format to byte slice format
	patternBytes := PatternToBytes(pattern)
	// Find all positions of the old pattern in the data
	positions := SearchPatterns(data, patternBytes)

	// Replace each occurrence of the old pattern with the new pattern
	for _, pos := range positions {
		// Extract the original opcode at the current position
		originalOpcode := data[pos : pos+len(patternBytes)]
		// Generate the new pattern with actual bytes replacing placeholders
		newPatternWithActualBytes := ReplacePlaceholders(newPattern, originalOpcode)
		// Replace the old pattern in the data at the specified position
		data = ReplaceAtPosition(data, pos, newPatternWithActualBytes)
	}
	return data
}

// ReplaceAtPosition replaces a portion of 'data' starting at 'pos' with 'newPattern'.
// If 'data' is not large enough, it will be expanded to accommodate the new pattern.
func ReplaceAtPosition(data []byte, pos int, newPattern []byte) []byte {
	patternLen := len(newPattern)
	// Expand 'data' if necessary to ensure it can accommodate the new pattern
	if len(data) < pos+patternLen {
		data = append(data, make([]byte, pos+patternLen-len(data))...)
	}
	// Copy the new pattern into the 'data' at the specified position
	copy(data[pos:pos+patternLen], newPattern)
	return data
}
