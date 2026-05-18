package cli

import "testing"

func TestResolvePGURLDefaultsToDataDatabase(t *testing.T) {
	t.Setenv("PGURL", "")

	got := ResolvePGURL("")
	want := "postgres:///data"

	if got != want {
		t.Fatalf("ResolvePGURL(\"\") = %q, want %q", got, want)
	}
}

func TestResolvePGURLPreservesExplicitConnections(t *testing.T) {
	t.Setenv("PGURL", "postgres:///envdb")

	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "explicit simple database",
			in:   "data",
			want: "postgres:///data",
		},
		{
			name: "explicit URL",
			in:   "postgresql:///other",
			want: "postgresql:///other",
		},
		{
			name: "keyword value connection",
			in:   "host=/tmp port=5432 dbname=data user=postgres",
			want: "host=/tmp port=5432 dbname=data user=postgres",
		},
		{
			name: "environment fallback",
			in:   "",
			want: "postgres:///envdb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ResolvePGURL(tt.in)
			if got != tt.want {
				t.Fatalf("ResolvePGURL(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
