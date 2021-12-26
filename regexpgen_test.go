package regexpgen

import (
	"bytes"
	"testing"
)

func TestGenerate(t *testing.T) {
	s := `foo(-(bar|baz)){2,4}`
	var buf bytes.Buffer
	if err := GenerateString(s, &buf, nil); err != nil {
		t.Fatal(err)
	}
	t.Log(buf.String())
}
