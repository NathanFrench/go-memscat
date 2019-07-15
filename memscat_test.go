package memscat

import (
	"syscall"
	"testing"
	"unsafe"
)

func TestMemscatRead(t *testing.T) {
	str := []byte{'f', 'o', 'o', 'b', 'a', 'r', '0'}
	pid := syscall.Getpid()

	out := Read(pid, unsafe.Pointer(&str[0]), 6)

	if string(out) != "foobar" {
		t.Errorf("expected \"foobar\" but got %v", string(out))
	}
}

func TestMemscatWrite(t *testing.T) {
	str := []byte{'X', 'Y', 'Z', 'b', 'a', 'r'}
	ovr := []byte{'f', 'o', 'o'}
	pid := syscall.Getpid()

	ret := Write(pid,
		unsafe.Pointer(&str[0]),
		unsafe.Pointer(&ovr[0]), 3)

	if ret < 0 || string(str) != "foobar" {
		t.Errorf("expected the buffer `str` to now contain \"foobar\", but got %v", string(str))
	}
}
