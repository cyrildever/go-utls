package bip32

import (
	"bytes"
	"testing"
)

func TestReverseCopy(t *testing.T) {
	testCases := []struct {
		dst, src []byte
		expect   []byte
	}{
		{
			make([]byte, 4),
			[]byte{0x12, 0x34},
			[]byte{0x00, 0x00, 0x12, 0x34},
		},
		{
			make([]byte, 2),
			[]byte{0x12, 0x34, 0x56},
			[]byte{0x34, 0x56},
		},
	}

	for i, c := range testCases {
		ReverseCopy(c.dst, c.src)

		if !bytes.Equal(c.dst, c.expect) {
			t.Fatalf("#%d failed: got %x, expect %x", i, c.dst, c.expect)
		}
	}
}

func TestPaddedAppend(t *testing.T) {
	testCases := []struct {
		size     uint
		dst, src []byte
		expect   []byte
	}{
		{2, nil, []byte{0x12}, []byte{0x00, 0x12}},
		{1, nil, []byte{0x12}, []byte{0x12}},
		{0, nil, []byte{0x12}, []byte{0x12}},
	}

	for i, c := range testCases {
		if got := paddedAppend(c.size, c.dst, c.src); !bytes.Equal(got, c.expect) {
			t.Fatalf("#%d failed: got %x, expect %x", i, got, c.expect)
		}
	}
}
