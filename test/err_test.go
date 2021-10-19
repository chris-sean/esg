package test

import (
	"testing"

	"github.com/SimpleFelix/esg"
)

func TestError(t *testing.T) {
	e := ErrTestErr("test")
	modifyE(&e)
	want := "TestErr"
	got := e.ErrorCode()
	if got != want {
		t.Errorf("want=%v; got=%v;", want, got)
	}
	want = "test"
	got = e.Error()
	if got != want {
		t.Errorf("want=%v; got=%v;", want, got)
	}
	wantSC := 500
	gotSC := e.StatusCode()
	if got != want {
		t.Errorf("want=%v; got=%v;", wantSC, gotSC)
	}
	want = "extra"
	gotExtra := e.Extra()
	got, ok := gotExtra.(string)
	if !ok {
		t.Errorf("%#v is not a string", got)
	}
	if got != want {
		t.Errorf("want=%v; got=%v;", want, got)
	}
}

func modifyE(e esg.ErrorTypeWriteable) {
	e.SetExtra("extra")
}
