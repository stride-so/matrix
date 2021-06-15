package test

import (
	"io"
	"os"
	"testing"
)

func FixtureReader(t *testing.T, path string) io.ReadCloser {
	t.Helper()
	r, err := os.Open(path)
	if err != nil {
		t.Fail()
	}
	return r
}

func TmpWriter(t *testing.T, pattern string) io.WriteCloser {
	t.Helper()
	tmp := "../tmp"
	err := os.MkdirAll(tmp, 0775)
	if err != nil {
		t.Fail()
	}
	w, err := os.CreateTemp(tmp, pattern)
	if err != nil {
		t.Fail()
	}
	return w
}
