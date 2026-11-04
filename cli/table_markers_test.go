package cli

import (
	"strings"
	"testing"
)

func TestMarkLargeMarkdownTables(t *testing.T) {
	wide := "| A | B | C | D | E |\n|---|---|---|---|---|\n| 1 | 2 | 3 | 4 | 5 |\n"
	short := "| Name | Version | Link |\n|---|---|---|\n| lib | 1 | [docs](https://example.org/an-extremely-long-reference-target-that-is-not-visible) |\n"
	long := "| Parameter | Default | Description |\n|---|---|---|\n| `shared_preload_libraries_for_a_very_long_module` | false | Setting |\n"
	tall := "| A | B |\n|---|---|\n" + strings.Repeat("| x | y |\n", 10)
	cases := []struct{ name, input, want string }{
		{"columns", wide, wide + "{.wide .expand}\n"},
		{"rows", tall, tall + "{.wide .expand}\n"},
		{"visible identifier", long, long + "{.wide .expand}\n"},
		{"link target", short, short},
		{"existing attributes", wide + "{#catalog .matrix caption=\"Packages\"}\n", wide + "{#catalog .matrix caption=\"Packages\" .wide .expand}\n"},
		{"quoted class", wide + "{class=\"matrix wide\" #catalog}\n", wide + "{class=\"matrix wide expand\" #catalog}\n"},
		{"single quoted class", wide + "{class='matrix wide' #catalog}\n", wide + "{class='matrix wide expand' #catalog}\n"},
		{"fields", wide + "{.fields}\n", wide + "{.fields}\n"},
		{"fenced source", "```markdown\n" + wide + "```\n", "```markdown\n" + wide + "```\n"},
		{"nested fence", "````markdown\n```\n" + wide + "```\n````\n", "````markdown\n```\n" + wide + "```\n````\n"},
		{"raw matrix", "{{< pgext_matrix >}}\n" + wide + "{{< /pgext_matrix >}}\n", "{{< pgext_matrix >}}\n" + wide + "{{< /pgext_matrix >}}\n"},
		{"indented source", "    " + strings.ReplaceAll(strings.TrimSuffix(wide, "\n"), "\n", "\n    ") + "\n", "    " + strings.ReplaceAll(strings.TrimSuffix(wide, "\n"), "\n", "\n    ") + "\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := MarkLargeMarkdownTables(tc.input)
			if got != tc.want {
				t.Fatalf("got %q; want %q", got, tc.want)
			}
			if twice := MarkLargeMarkdownTables(got); twice != got {
				t.Fatal("not idempotent")
			}
		})
	}
}
