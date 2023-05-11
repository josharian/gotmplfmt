package parse

import "testing"

func FuzzParseString(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		root, err := Parse(s)
		if err != nil {
			return
		}
		out := root.String()
		root2, err := Parse(out)
		if err != nil {
			t.Fatal(err)
		}
		out2 := root2.String()
		if out != out2 {
			t.Fatalf("round trip failure: %q != %q", out, out2)
		}
	})
}
