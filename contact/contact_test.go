package contact

import (
	"testing"
)

func TestData(t *testing.T) {
	if len(Data.Emergency) == 0 {
		t.Error("Emergency contacts should not be empty")
	}
	if len(Data.Security.Peace) == 0 {
		t.Error("Peace security contacts should not be empty")
	}
	if Data.Emergency[0].Phones[0] != "07-7172930" {
		t.Errorf("Expected 07-7172930, got %s", Data.Emergency[0].Phones[0])
	}
}
