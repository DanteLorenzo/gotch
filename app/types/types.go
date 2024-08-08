package types

// Pattern represents a structure for holding the old and new patterns used in the patching process.
// The struct fields are tagged with JSON annotations to specify how they should be serialized and deserialized.
type Pattern struct {
    OldPattern string `json:"OldPattern"` // OldPattern is the pattern that needs to be found and replaced.
    NewPattern string `json:"NewPattern"` // NewPattern is the pattern that will replace the OldPattern.
}
