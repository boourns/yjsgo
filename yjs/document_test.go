package yjs

import (
	"testing"
)

func TestDocumentWritingWorks(t *testing.T) {
	fox := "quick brown fox"
	d1 := NewTextDocument(&fox)
	d2 := NewTextDocument(nil)

	defer d1.Close()
	defer d2.Close()

	targetStateVector, err := d2.StateVector()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	syncUpdate, err := d1.EncodeStateAsUpdate(targetStateVector)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	err = d2.ApplyUpdate(syncUpdate)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	targetStateVector, err = d2.StateVector()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	result, err := d2.ToString()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if result != "quick brown fox" {
		t.Fatalf("Expected d2 to sync to same state, got %v", result)
	}

	finalStateVector1, err := d1.StateVector()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	finalStateVector2, err := d2.StateVector()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if finalStateVector1 != finalStateVector2 {
		t.Fatalf("Expected equal final state vectors")
	}
}

func TestCreateFromJSON(t *testing.T) {
	obj := `{"foo": [1, 2, 3]}`

	d1 := NewComplexDocument(&obj)
	result, err := d1.ToString()
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if result != `{"foo":[1,2,3]}` {
		t.Fatalf("Expected d2 to sync to same state, got %v", result)
	}
}

func BenchmarkApplyUpdates(b *testing.B) {
	fox := "quick brown fox"
	blank := ""

	for i := 0; i < b.N * 100000; i++ {
		d1 := NewTextDocument(&fox)
		d2 := NewTextDocument(&blank)

		for j := 0; j < 10; j++ {
			d1.Insert(0,"a")
			sv, err := d2.StateVector()
			if err != nil {
				b.Fatalf("Error: %v", err)
			}
			update, err := d1.EncodeStateAsUpdate(sv)
			if err != nil {
				b.Fatalf("Error: %v", err)
			}
			err = d2.ApplyUpdate(update)
			if err != nil {
				b.Fatalf("Error: %v", err)
			}
			flatText, err := d2.ToString()
			if err != nil {
				b.Fatalf("Error: %v", err)
			}
			d3 := NewTextDocument(&flatText)
			_, err = d3.EncodeStateAsUpdate("")
			d3.Close()
		}

		d1.Close()
		d2.Close()
	}
}