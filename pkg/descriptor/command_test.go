package descriptor

import "testing"

var command = Command{
	Value: "/usr/bin/find -r password.txt",
}

func TestCommandDescriptor(t *testing.T) {
	if command.GetTitle() != "command" {
		t.Error("invalid title returned", "received", command.GetTitle())
	}
	if command.GetMin() != 0 {
		t.Error("min value should be 0", "received", command.GetMin())
	}
	if command.GetMax() != 0 {
		t.Error("max value should be 0", "received", command.GetMax())
	}
	if command.Expression() != command.Value {
		t.Error("value provide is not persisted", "received", command.Expression())
	}
}
