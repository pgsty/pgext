package cli

import (
	"html"
	"regexp"
	"strings"

	"golang.org/x/text/width"
)

var (
	tableFence     = regexp.MustCompile("^ {0,3}(`{3,}|~{3,})")
	tableDelimiter = regexp.MustCompile(`^:?-+:?$`)
	tableAttribute = regexp.MustCompile(`^\s*\{[^{}\n]*\}\s*$`)
	tableDotClass  = regexp.MustCompile(`(?:^|\s|\{)\.([\w-]+)`)
	tableClass     = regexp.MustCompile(`\bclass\s*=\s*(?:"([^"]*)"|'([^']*)')`)
	tableRaw       = regexp.MustCompile(`\{\{[<%]\s*(/?)\s*(pgext_matrix|pgext_os_matrix|raw|rawhtml|code)\b`)
	tableLink      = regexp.MustCompile(`!?\[([^\]]*)\]\((?:[^()]|\([^)]*\))*\)`)
	tableReference = regexp.MustCompile(`\[([^\]]+)\]\[[^\]]*\]`)
	tableHTML      = regexp.MustCompile(`<[^>]+>`)
	tableURL       = regexp.MustCompile(`https?://\S+`)
	tableShortcode = regexp.MustCompile(`\{\{[<%].*?[%>]\}\}`)
	tableLongToken = regexp.MustCompile(`[A-Za-z0-9_./:-]{28,}`)
)

func tableCells(line string) []string {
	s := strings.TrimSpace(line)
	var parts []string
	start, escaped := 0, false
	for i, r := range s {
		if r == '|' && !escaped {
			parts = append(parts, strings.TrimSpace(s[start:i]))
			start = i + 1
		}
		if r == '\\' && !escaped {
			escaped = true
		} else {
			escaped = false
		}
	}
	parts = append(parts, strings.TrimSpace(s[start:]))
	if len(parts) > 0 && parts[0] == "" {
		parts = parts[1:]
	}
	if len(parts) > 0 && parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}
	return parts
}

func isTableDelimiter(line string) bool {
	parts := tableCells(line)
	if len(parts) < 2 {
		return false
	}
	for _, p := range parts {
		if !tableDelimiter.MatchString(p) {
			return false
		}
	}
	return true
}

func tableClasses(attr string) map[string]bool {
	classes := map[string]bool{}
	for _, m := range tableDotClass.FindAllStringSubmatch(attr, -1) {
		classes[m[1]] = true
	}
	if m := tableClass.FindStringSubmatch(attr); m != nil {
		for _, c := range strings.Fields(m[1] + m[2]) {
			classes[c] = true
		}
	}
	return classes
}

func tableVisibleCell(cell string) string {
	cell = tableLink.ReplaceAllString(cell, "$1")
	cell = tableReference.ReplaceAllString(cell, "$1")
	cell = tableHTML.ReplaceAllString(cell, "")
	cell = tableURL.ReplaceAllString(cell, "")
	cell = tableShortcode.ReplaceAllString(cell, "")
	cell = strings.NewReplacer("*", "", "`", "", "_", "", "\\|", "|").Replace(cell)
	return html.UnescapeString(cell)
}

func needsTableReader(rows []string, attr string) bool {
	classes := tableClasses(attr)
	if classes["fields"] {
		return false
	}
	cols := len(tableCells(rows[0]))
	if classes["wide"] || cols >= 5 || len(rows)-2 >= 10 {
		return true
	}
	if cols < 3 {
		return false
	}
	for i, row := range rows {
		if i == 1 {
			continue
		}
		for _, cell := range tableCells(row) {
			text := tableVisibleCell(cell)
			if tableLongToken.MatchString(text) {
				return true
			}
			n := 0
			for _, r := range text {
				n++
				kind := width.LookupRune(r).Kind()
				if kind == width.EastAsianWide || kind == width.EastAsianFullwidth {
					n++
				}
			}
			if n >= 72 {
				return true
			}
		}
	}
	return false
}

func addTableReaderClasses(attr string) string {
	classes := tableClasses(attr)
	var missing []string
	for _, name := range []string{"wide", "expand"} {
		if !classes[name] {
			missing = append(missing, name)
		}
	}
	if len(missing) == 0 {
		return attr
	}
	if m := tableClass.FindStringSubmatchIndex(attr); m != nil {
		end := m[3]
		if end < 0 {
			end = m[5]
		}
		return attr[:end] + " " + strings.Join(missing, " ") + attr[end:]
	}
	end := strings.LastIndex(attr, "}")
	return strings.TrimRight(attr[:end], " \t") + " ." + strings.Join(missing, " .") + attr[end:]
}

// MarkLargeMarkdownTables retains explicit OINK reader markers on generated
// pages. It operates after stub composition; source stubs and custom matrix
// shortcode data stay unchanged. Link destinations do not count as cell width.
func MarkLargeMarkdownTables(content string) string {
	lines := strings.SplitAfter(content, "\n")
	var out strings.Builder
	i := 0
	if len(lines) > 0 && (strings.TrimSpace(lines[0]) == "---" || strings.TrimSpace(lines[0]) == "+++") {
		marker := strings.TrimSpace(lines[0])
		out.WriteString(lines[0])
		i = 1
		for i < len(lines) {
			out.WriteString(lines[i])
			i++
			if strings.TrimSpace(lines[i-1]) == marker {
				break
			}
		}
	}
	fence, fenceLen, raw, rawHTML := "", 0, "", ""
	for i < len(lines) {
		line := lines[i]
		s := strings.TrimSpace(line)
		m := tableFence.FindStringSubmatch(line)
		if fence != "" {
			if m != nil && strings.HasPrefix(m[1], fence) && len(m[1]) >= fenceLen {
				fence = ""
			}
			out.WriteString(line)
			i++
			continue
		}
		if m != nil {
			fence = m[1][:1]
			fenceLen = len(m[1])
			out.WriteString(line)
			i++
			continue
		}
		if raw != "" {
			if m := tableRaw.FindStringSubmatch(s); m != nil && m[1] == "/" && m[2] == raw {
				raw = ""
			}
			out.WriteString(line)
			i++
			continue
		}
		if m := tableRaw.FindStringSubmatch(s); m != nil && m[1] == "" {
			raw = m[2]
			out.WriteString(line)
			i++
			continue
		}
		if rawHTML != "" {
			if strings.Contains(s, rawHTML) {
				rawHTML = ""
			}
			out.WriteString(line)
			i++
			continue
		}
		for _, pair := range [][2]string{{"<!--", "-->"}, {"<pre", "</pre>"}, {"<code", "</code>"}, {"<table", "</table>"}, {"<script", "</script>"}, {"<style", "</style>"}} {
			if strings.HasPrefix(s, pair[0]) && !strings.Contains(s, pair[1]) {
				rawHTML = pair[1]
				break
			}
		}
		if rawHTML != "" || strings.HasPrefix(s, ">") || strings.HasPrefix(line, "    ") || strings.HasPrefix(line, "\t") {
			out.WriteString(line)
			i++
			continue
		}
		if strings.Contains(s, "|") && !strings.HasPrefix(s, "<") && i+1 < len(lines) && isTableDelimiter(lines[i+1]) && len(tableCells(s)) == len(tableCells(lines[i+1])) {
			end := i + 2
			for end < len(lines) {
				t := strings.TrimSpace(lines[end])
				if t == "" || !strings.Contains(t, "|") || strings.HasPrefix(t, ">") || strings.HasPrefix(t, "<") || strings.HasPrefix(t, "{{") || tableFence.MatchString(lines[end]) {
					break
				}
				end++
			}
			attr := ""
			if end < len(lines) && tableAttribute.MatchString(lines[end]) {
				attr = lines[end]
			}
			for _, row := range lines[i:end] {
				out.WriteString(row)
			}
			if needsTableReader(lines[i:end], attr) {
				if attr != "" {
					attr = addTableReaderClasses(attr)
				} else {
					newline := "\n"
					if strings.HasSuffix(lines[end-1], "\r\n") {
						newline = "\r\n"
					}
					if !strings.HasSuffix(lines[end-1], "\n") {
						out.WriteString(newline)
					}
					out.WriteString("{.wide .expand}" + newline)
				}
			}
			if attr != "" {
				out.WriteString(attr)
				end++
			}
			i = end
		} else {
			out.WriteString(line)
			i++
		}
	}
	return out.String()
}
