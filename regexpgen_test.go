package regexpgen

import (
	"bytes"
	"regexp/syntax"
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

func BenchmarkString(b *testing.B) {
	s := `foo(-(bar|baz)){2,4}`
	re, err := syntax.Parse(s, syntax.Perl)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		s, err := String(re)
		if err != nil {
			b.Fatal(err)
		}
		if s == "" {
			b.Fatal("empty string")
		}
	}
}
