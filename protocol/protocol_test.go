package protocol

import (
	"bytes"
	"testing"
)

func TestUnpackString(t *testing.T) {
	buf := []byte{22, 0, 84, 104, 105, 115, 32, 105, 115, 32, 97, 32, 116, 101, 115, 116, 32, 101, 120, 97, 109, 112, 108, 101}
	expected := "This is a test example"
	got, _ := UnpackString(buf, 0)
	if got != expected {
		t.Errorf("UnpackString(...) = %s, expected: %s", got, expected)
	}
}

func TestPackString(t *testing.T) {
	expected := []byte{22, 0, 84, 104, 105, 115, 32, 105, 115, 32, 97, 32, 116, 101, 115, 116, 32, 101, 120, 97, 109, 112, 108, 101}
	str := "This is a test example"
	got := packString(str)
	if !bytes.Equal(got, expected) {
		t.Errorf("packString(...) = %d, expected: %d", got, expected)
	}
}
