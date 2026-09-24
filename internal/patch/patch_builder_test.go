//go:build !integration

package patch

import "testing"

func TestBuilderReplaceAddRemove(t *testing.T) {
	b := &AbstractBuilder{}
	b.Define = func() {
		b.Compare("/name", "new", "old")       // replace
		b.Compare("/description", "hello", "") // add
		b.Compare("/legacy", "", "gone")       // remove
		b.Compare("/type", "same", "same")     // no-op
	}

	ops := b.Generate()
	if len(ops) != 3 {
		t.Fatalf("expected 3 operations, got %d: %+v", len(ops), ops)
	}

	byPath := map[string]Operation{}
	for _, op := range ops {
		byPath[op.Path] = op
	}

	if op := byPath["/name"]; op.Op != OpReplace || op.Value != "new" {
		t.Fatalf("unexpected /name op: %+v", op)
	}
	if op := byPath["/description"]; op.Op != OpAdd || op.Value != "hello" {
		t.Fatalf("unexpected /description op: %+v", op)
	}
	if op := byPath["/legacy"]; op.Op != OpRemove {
		t.Fatalf("unexpected /legacy op: %+v", op)
	}
	if _, ok := byPath["/type"]; ok {
		t.Fatalf("expected no op for unchanged /type")
	}
}
