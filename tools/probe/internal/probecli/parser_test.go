package probecli

import "testing"

func TestParseCheckFlags(t *testing.T) {
	cases := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{name: "single url", args: []string{"--url", "https://example.com"}},
		{name: "multiple urls", args: []string{"--url", "https://a.com", "--url", "https://b.com"}},
		{name: "no url", args: []string{}, wantErr: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ParseCheckFlags(c.args)

			if c.wantErr {
				if err == nil {
					t.Fatalf("ParseCheckFlags(%v) = %+v, want error", c.args, got)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseCheckFlags(%v) unexpected error: %v", c.args, err)
			}
		})
	}
}

func TestParseRecoverFlags(t *testing.T) {
	cases := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{name: "url given", args: []string{"--url", "https://example.com"}},
		{name: "no url", args: []string{}, wantErr: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ParseRecoverFlags(c.args)

			if c.wantErr {
				if err == nil {
					t.Fatalf("ParseRecoverFlags(%v) = %+v, want error", c.args, got)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseRecoverFlags(%v) unexpected error: %v", c.args, err)
			}
		})
	}
}
