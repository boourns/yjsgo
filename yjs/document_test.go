package yjs

import (
	"testing"
)

func TestBlah(t *testing.T) {
	d1 := NewDocument("quick brown fox")
	d2 := NewDocument("")

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

func BenchmarkApplyUpdates(b *testing.B) {
	for i := 0; i < b.N * 100000; i++ {
		d1 := NewDocument("quick brown fox")
		d2 := NewDocument("")

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
			d3 := NewDocument(flatText)
			_, err = d3.EncodeStateAsUpdate("")
			d3.Close()
		}

		d1.Close()
		d2.Close()
	}
}