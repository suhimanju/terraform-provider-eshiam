// Package patch provides a generic RFC-6902 JSON Patch builder used by
// resources that update remote objects via HTTP PATCH.
//
// The builder diffs a "current" object against a "modified" object and emits
// the minimal set of add/replace/remove operations. It is intentionally
// decoupled from any specific SDK: operations are plain structs you can marshal
// to JSON and send to your API.
package patch

// Operation is a single RFC-6902 JSON Patch operation.
type Operation struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value any    `json:"value,omitempty"`
}

// Op values defined by RFC-6902.
const (
	OpAdd     = "add"
	OpReplace = "replace"
	OpRemove  = "remove"
)

// comparison holds a single desired/current value pair for a JSON path.
type comparison struct {
	path     string
	modified any
	current  any
}

// AbstractBuilder provides the diffing machinery. Embed it in a concrete
// builder and populate comparisons in DefineComparisons using Compare.
type AbstractBuilder struct {
	comparisons []comparison
	operations  []Operation
	// Define is set by the concrete builder to its DefineComparisons method so
	// Generate can trigger registration lazily.
	Define func()
}

// Compare registers a desired vs current value pair for the given JSON path.
func (b *AbstractBuilder) Compare(path string, modified, current any) {
	b.comparisons = append(b.comparisons, comparison{path: path, modified: modified, current: current})
}

// Generate runs the registered comparisons and returns the resulting patch.
func (b *AbstractBuilder) Generate() []Operation {
	if b.Define != nil {
		b.Define()
	}
	for _, c := range b.comparisons {
		b.appendOp(c)
	}
	return b.operations
}

// appendOp emits the appropriate operation for a single comparison.
func (b *AbstractBuilder) appendOp(c comparison) {
	switch {
	case isEmpty(c.modified) && !isEmpty(c.current):
		b.operations = append(b.operations, Operation{Op: OpRemove, Path: c.path})
	case isEmpty(c.current) && !isEmpty(c.modified):
		b.operations = append(b.operations, Operation{Op: OpAdd, Path: c.path, Value: c.modified})
	case !valuesEqual(c.modified, c.current):
		b.operations = append(b.operations, Operation{Op: OpReplace, Path: c.path, Value: c.modified})
	}
}
