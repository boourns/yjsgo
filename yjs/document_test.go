package yjs

import (
	"log"
	"testing"
)

func TestBlah(t *testing.T) {
	d1 := NewDocument("quick brown fox")
	d2 := NewDocument("")

	targetStateVector, err := d2.StateVector()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	log.Print("targetStateVector is ", targetStateVector)

	syncUpdate, err := d1.EncodeStateAsUpdate(targetStateVector)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	log.Print("syncUpdate is ", syncUpdate)

	err = d2.ApplyUpdate(syncUpdate)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	targetStateVector, err = d2.StateVector()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	log.Print("new targetStateVector is ", targetStateVector)


	result, err := d2.ToString()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if result != "quick brown fox" {
		t.Fatalf("Expected d2 to sync to same state, got %v", result)
	}

}