package soap

import "testing"

func TestEscapeXMLText(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"abc123", "abc123"},
		{"<foo>&", "&lt;foo&gt;&amp;"},
		{"\"foo'", "\"foo'"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.input, func(t *testing.T) {
			got := EscapeXMLText(test.input)
			if got != test.want {
				t.Errorf("want %q, got %q", test.want, got)
			}
		})
	}
}
