package gotype

import (
	"go/build"
	"testing"
)

func TestCheck(t *testing.T) {
	const filename = "/Users/Charlie/go/src/git.vieth.io/mgo/gotype/gotype.go"
	ctxt := &build.Default
	list, err := Check(ctxt, filename, nil, false, true)
	if err != nil {
		t.Fatalf("Check: %s", err)
	}
	if len(list) != 0 {
		t.Fatalf("Check expected nil error list: %+v", list)
	}
}
