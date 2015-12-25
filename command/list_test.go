package command

import "testing"

func TestDoneLabel(t *testing.T) {
	if doneLabel(0) != "-" {
		t.Error("Failed doneLabel(0)")
	}
	if doneLabel(1) != "Done" {
		t.Error("Failed doneLabel(1)")
	}
}
