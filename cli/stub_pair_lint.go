package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

var (
	pairURLRE         = regexp.MustCompile(`(?i)https?://[^\s<>\[\]"'` + "`" + `]+`)
	pairHeadingRE     = regexp.MustCompile(`^(#{1,6})[ \t]+(.+?)[ \t]*#*[ \t]*$`)
	pairFenceRE       = regexp.MustCompile("^[ \\t]*(`{3,}|~{3,})(.*)$")
	pairBulletRE      = regexp.MustCompile(`^[-*+]\s+`)
	pairLinkBulletRE  = regexp.MustCompile(`^[-*+]\s+\[[^]]+\]\(`)
	pairHTMLTagRE     = regexp.MustCompile(`<[^>\n]*>`)
	pairTableSepRE    = regexp.MustCompile(`^\s*:?-{3,}:?\s*$`)
	pairEnglishWordRE = regexp.MustCompile(`[A-Za-z]+(?:['\x{2019}][A-Za-z]+)?`)
	pairCreateExtRE   = regexp.MustCompile(`(?i)\bCREATE\s+EXTENSION\b`)
	pairTableProseRE  = regexp.MustCompile(`(?i)\b(?:time-varying|single value|set of|number of|path to|whether to|true if|returns? the|creates? an?|enables? the|disables? the)\b`)
	pairForbiddenRE   = regexp.MustCompile(`(?i)\b(?:TODO|TBD|PLACEHOLDER)\b|lorem\s+ipsum|as\s+an\s+AI|AI\s+language\s+model|I\s+(?:cannot|can't)\s+(?:browse|access)|作为(?:一个|一名)?\s*(?:AI|人工智能)|(?:我|本模型)无法(?:浏览|访问)`)
)

var englishProseMarkers = map[string]struct{}{
	"a": {}, "an": {}, "and": {}, "are": {}, "as": {}, "at": {}, "be": {}, "by": {},
	"can": {}, "for": {}, "from": {}, "has": {}, "have": {}, "in": {}, "into": {},
	"is": {}, "of": {}, "on": {}, "or": {}, "provides": {}, "records": {},
	"requires": {}, "returns": {}, "supports": {}, "that": {}, "the": {}, "this": {},
	"these": {}, "those": {}, "to": {}, "uses": {}, "using": {}, "was": {}, "were": {},
	"with": {}, "without": {},
}

var translatableTableHeaders = map[string]struct{}{
	"default": {}, "description": {}, "effect": {}, "example": {}, "examples": {},
	"feature": {}, "function": {}, "functions": {}, "language": {}, "note": {},
	"notes": {}, "operator": {}, "operators": {}, "option": {}, "options": {},
	"parameter": {}, "parameters": {}, "purpose": {}, "return": {}, "returns": {},
	"type": {}, "types": {},
}

type stubCodeBlock struct {
	info string
	body string
}

type stubDocFacts struct {
	errors        []string
	urls          []string
	blocks        []stubCodeBlock
	inline        []string
	headingLevels []int
}

// LintStubPair checks a bilingual stub pair against the repository document contract.
func LintStubPair(enPath, zhPath, name string, needDDL *bool) error {
	en := inspectStubDoc(enPath, "en", name, needDDL)
	zh := inspectStubDoc(zhPath, "zh", name, needDDL)

	var issues []error
	for _, msg := range en.errors {
		issues = append(issues, fmt.Errorf("%s: en: %s", name, msg))
	}
	for _, msg := range zh.errors {
		issues = append(issues, fmt.Errorf("%s: zh: %s", name, msg))
	}
	if len(en.urls) > 0 && len(zh.urls) > 0 && !stringMultisetEqual(en.urls, zh.urls) {
		issues = append(issues, fmt.Errorf("%s: en/zh URL multisets differ", name))
	}
	if !codeBlocksEqual(en.blocks, zh.blocks) {
		issues = append(issues, fmt.Errorf("%s: en/zh fenced code blocks differ", name))
	}
	if !stringMultisetEqual(en.inline, zh.inline) {
		issues = append(issues, fmt.Errorf("%s: en/zh inline code identifiers differ", name))
	}
	if len(en.headingLevels) > 0 && len(zh.headingLevels) > 0 && !intSliceEqual(en.headingLevels, zh.headingLevels) {
		issues = append(issues, fmt.Errorf("%s: en/zh heading-level structure differs", name))
	}
	return errors.Join(issues...)
}

// LintStubPairs validates bilingual stub pairs in enDir and zhDir.
// When names is empty, every Markdown stem found in either directory is checked.
func LintStubPairs(enDir, zhDir string, names []string) (int, error) {
	if len(names) == 0 {
		set := map[string]struct{}{}
		for _, dir := range []string{enDir, zhDir} {
			entries, err := os.ReadDir(dir)
			if err != nil {
				return 0, fmt.Errorf("read stub directory %s: %w", dir, err)
			}
			for _, e := range entries {
				if e.IsDir() || filepath.Ext(e.Name()) != ".md" {
					continue
				}
				set[strings.TrimSuffix(e.Name(), ".md")] = struct{}{}
			}
		}
		names = make([]string, 0, len(set))
		for name := range set {
			names = append(names, name)
		}
		sort.Strings(names)
	}

	var issues []error
	for _, name := range names {
		if err := SafeStubName(name); err != nil {
			issues = append(issues, err)
			continue
		}
		enPath := filepath.Join(enDir, name+".md")
		zhPath := filepath.Join(zhDir, name+".md")
		if err := LintStubPair(enPath, zhPath, name, nil); err != nil {
			issues = append(issues, err)
		}
	}
	return len(names), errors.Join(issues...)
}

func inspectStubDoc(path, lang, name string, needDDL *bool) stubDocFacts {
	raw, err := os.ReadFile(path)
	if err != nil {
		return stubDocFacts{errors: []string{fmt.Sprintf("missing file: %s", path)}}
	}
	text := strings.ReplaceAll(string(raw), "\r\n", "\n")
	var errs []string
	if strings.TrimSpace(text) == "" {
		errs = append(errs, "file is empty")
	}
	if mdErr := ValidateStubMarkdown(path, text); mdErr != nil {
		for _, line := range strings.Split(mdErr.Error(), "\n") {
			if line = strings.TrimSpace(line); line != "" {
				errs = append(errs, line)
			}
		}
	}

	lines := splitLines(text)
	for _, line := range lines {
		if s := strings.TrimSpace(line); s != "" {
			if s == "---" {
				errs = append(errs, "front matter is forbidden")
			}
			break
		}
	}

	blocks, fenced, fenceErrs := parsePairFences(lines)
	errs = append(errs, fenceErrs...)

	type heading struct {
		level int
		title string
		line  int
	}
	var headings []heading
	for i, line := range lines {
		lineNo := i + 1
		if fenced[lineNo] {
			continue
		}
		if m := pairHeadingRE.FindStringSubmatch(line); m != nil {
			headings = append(headings, heading{len(m[1]), strings.TrimSpace(m[2]), lineNo})
		}
		if i+1 < len(lines) && strings.TrimSpace(line) != "" {
			next := strings.TrimSpace(lines[i+1])
			if isUniformRune(next, '=') {
				errs = append(errs, fmt.Sprintf("setext H1 is forbidden at line %d", lineNo))
			} else if isUniformRune(next, '-') {
				errs = append(errs, fmt.Sprintf("setext H2 is forbidden at line %d", lineNo))
			}
		}
	}

	var h2 []heading
	levels := make([]int, 0, len(headings))
	for _, h := range headings {
		levels = append(levels, h.level)
		if h.level == 1 {
			errs = append(errs, "H1 is forbidden")
		}
		if h.level == 2 {
			h2 = append(h2, h)
		}
	}

	expectedH2, expectedSources := "Usage", "Sources:"
	if lang == "zh" {
		expectedH2, expectedSources = "用法", "来源："
	}
	if len(h2) != 1 || h2[0].title != expectedH2 {
		errs = append(errs, fmt.Sprintf("expected exactly one H2 named %q", expectedH2))
	} else {
		cursor := h2[0].line
		for cursor < len(lines) && strings.TrimSpace(lines[cursor]) == "" {
			cursor++
		}
		actual := ""
		if cursor < len(lines) {
			actual = strings.TrimSpace(lines[cursor])
		}
		if actual != expectedSources {
			errs = append(errs, fmt.Sprintf("first content after H2 must be %q, got %q", expectedSources, actual))
		} else {
			cursor++
			var srcLines []string
			for cursor < len(lines) {
				s := strings.TrimSpace(lines[cursor])
				if s == "" {
					cursor++
					continue
				}
				if pairBulletRE.MatchString(s) {
					srcLines = append(srcLines, s)
					cursor++
					continue
				}
				break
			}
			if len(srcLines) == 0 || len(extractStubURLs(strings.Join(srcLines, "\n"))) == 0 {
				errs = append(errs, "Sources section must begin with a linked bullet list")
			}
		}
	}

	urls := extractStubURLs(text)
	if len(urls) == 0 {
		errs = append(errs, "no source URL found")
	}
	if pairForbiddenRE.MatchString(text) {
		errs = append(errs, "forbidden placeholder or model-disclaimer text found")
	}
	errs = append(errs, stubLanguageErrors(lines, fenced, lang)...)
	if name != "" && !strings.Contains(text, name) {
		errs = append(errs, fmt.Sprintf("extension name %q does not occur verbatim", name))
	}
	if needDDL != nil && !*needDDL {
		for _, b := range blocks {
			if pairCreateExtRE.MatchString(b.body) {
				errs = append(errs, "need_ddl=false document contains CREATE EXTENSION code")
				break
			}
		}
	}

	var outside []string
	for i, line := range lines {
		if !fenced[i+1] {
			outside = append(outside, line)
		}
	}
	return stubDocFacts{
		errors:        errs,
		urls:          urls,
		blocks:        blocks,
		inline:        extractInlineCode(strings.Join(outside, "\n")),
		headingLevels: levels,
	}
}

func splitLines(text string) []string {
	if text == "" {
		return nil
	}
	return strings.Split(strings.TrimSuffix(text, "\n"), "\n")
}

func parsePairFences(lines []string) ([]stubCodeBlock, map[int]bool, []string) {
	var blocks []stubCodeBlock
	fenced := map[int]bool{}
	var errs []string
	var openMarker, info string
	var body []string
	openLine := 0
	for i, line := range lines {
		lineNo := i + 1
		m := pairFenceRE.FindStringSubmatch(line)
		if openMarker == "" {
			if m != nil {
				openMarker = m[1]
				info = strings.TrimSpace(m[2])
				body = nil
				openLine = lineNo
				fenced[lineNo] = true
			}
		} else {
			fenced[lineNo] = true
			if m != nil && m[1][0] == openMarker[0] && len(m[1]) >= len(openMarker) && strings.TrimSpace(m[2]) == "" {
				blocks = append(blocks, stubCodeBlock{info: info, body: strings.Join(body, "\n")})
				openMarker, info, body = "", "", nil
			} else {
				body = append(body, line)
			}
		}
	}
	if openMarker != "" {
		errs = append(errs, fmt.Sprintf("unclosed code fence opened at line %d", openLine))
	}
	return blocks, fenced, errs
}

func extractStubURLs(text string) []string {
	matches := pairURLRE.FindAllString(text, -1)
	urls := make([]string, 0, len(matches))
	for _, cand := range matches {
		balance, end := 0, len(cand)
		for i := 0; i < len(cand); i++ {
			if cand[i] == '(' {
				balance++
			} else if cand[i] == ')' {
				if balance == 0 {
					end = i
					break
				}
				balance--
			}
		}
		u := strings.TrimRight(cand[:end], ".,;:!?，。；：！？")
		if u != "" {
			urls = append(urls, u)
		}
	}
	return urls
}

// extractInlineCode extracts single-backtick spans (`code`), skipping multi-backtick spans.
func extractInlineCode(text string) []string {
	var out []string
	n := len(text)
	for i := 0; i < n; {
		if text[i] != '`' {
			i++
			continue
		}
		j := i
		for j < n && text[j] == '`' {
			j++
		}
		if j-i != 1 {
			i = j
			continue
		}
		end := -1
		for k := j; k < n; k++ {
			if text[k] == '\n' {
				break
			}
			if text[k] == '`' {
				if k+1 < n && text[k+1] == '`' {
					break
				}
				end = k
				break
			}
		}
		if end > j {
			out = append(out, text[j:end])
			i = end + 1
		} else {
			i = j
		}
	}
	return out
}

func stripInlineCode(s string) string {
	var b strings.Builder
	n := len(s)
	for i := 0; i < n; {
		if s[i] == '`' && (i == 0 || s[i-1] != '`') && (i+1 == n || s[i+1] != '`') {
			end := strings.IndexByte(s[i+1:], '`')
			if end > 0 {
				next := i + 1 + end
				if next+1 == n || s[next+1] != '`' {
					b.WriteByte(' ')
					i = next + 1
					continue
				}
			}
		}
		b.WriteByte(s[i])
		i++
	}
	return b.String()
}

func cleanProseFragment(s string) string {
	s = stripInlineCode(s)
	s = pairURLRE.ReplaceAllString(s, " ")
	s = pairHTMLTagRE.ReplaceAllString(s, " ")
	return s
}

func stubLanguageErrors(lines []string, fenced map[int]bool, lang string) []string {
	sourceMarker := "Sources:"
	if lang == "zh" {
		sourceMarker = "来源："
	}
	sourceLines := map[int]bool{}
	for i, line := range lines {
		lineNo := i + 1
		if fenced[lineNo] || strings.TrimSpace(line) != sourceMarker {
			continue
		}
		sourceLines[lineNo] = true
		sawBullet := false
		for c := i + 1; c < len(lines); c++ {
			s := strings.TrimSpace(lines[c])
			if s == "" {
				sourceLines[c+1] = true
				continue
			}
			if pairBulletRE.MatchString(s) {
				sawBullet = true
				sourceLines[c+1] = true
				continue
			}
			if sawBullet && len(lines[c]) > 0 && (lines[c][0] == ' ' || lines[c][0] == '\t') {
				sourceLines[c+1] = true
				continue
			}
			break
		}
		break
	}

	type numberedLine struct {
		num  int
		text string
	}
	proseLines := make([]numberedLine, 0, len(lines))
	for i, line := range lines {
		lineNo := i + 1
		if fenced[lineNo] || sourceLines[lineNo] || strings.HasPrefix(line, "    ") || strings.HasPrefix(line, "\t") {
			proseLines = append(proseLines, numberedLine{lineNo, ""})
			continue
		}
		s := strings.TrimSpace(line)
		if (strings.HasPrefix(s, "|") || strings.HasPrefix(s, "+")) && strings.Count(s, "|") >= 2 {
			proseLines = append(proseLines, numberedLine{lineNo, ""})
			continue
		}
		if pairLinkBulletRE.MatchString(s) && strings.Contains(strings.ToLower(s), "http") {
			proseLines = append(proseLines, numberedLine{lineNo, ""})
			continue
		}
		proseLines = append(proseLines, numberedLine{lineNo, cleanProseFragment(line)})
	}

	var paragraphs []numberedLine
	var parts []string
	start := 0
	for _, nl := range proseLines {
		if s := strings.TrimSpace(nl.text); s != "" {
			if len(parts) == 0 {
				start = nl.num
			}
			parts = append(parts, s)
		} else if len(parts) > 0 {
			paragraphs = append(paragraphs, numberedLine{start, strings.Join(parts, " ")})
			parts = nil
		}
	}
	if len(parts) > 0 {
		paragraphs = append(paragraphs, numberedLine{start, strings.Join(parts, " ")})
	}

	errs := stubTableLanguageErrors(lines, fenced, lang)
	if lang == "en" {
		for _, p := range paragraphs {
			han, maxRun := countHan(p.text)
			latin := countEnglishLetters(p.text)
			if han >= 12 || (maxRun >= 4 && han*2 >= max(latin, 1)) {
				errs = append(errs, fmt.Sprintf("English prose contains untranslated Han text near line %d", p.num))
			}
		}
		return errs
	}

	var substantive strings.Builder
	for _, nl := range proseLines {
		if strings.TrimSpace(nl.text) != "" && !pairHeadingRE.MatchString(nl.text) {
			substantive.WriteString(nl.text)
			substantive.WriteByte('\n')
		}
	}
	if han, _ := countHan(substantive.String()); han < 4 {
		errs = append(errs, "Chinese prose is missing outside headings, sources, and code")
	}
	for _, p := range paragraphs {
		if han, _ := countHan(p.text); han > 0 {
			continue
		}
		words := pairEnglishWordRE.FindAllString(p.text, -1)
		letters := 0
		hasMarker := false
		for _, w := range words {
			letters += len(w)
			if _, ok := englishProseMarkers[strings.ToLower(w)]; ok {
				hasMarker = true
			}
		}
		if len(words) >= 5 && letters >= 24 && hasMarker {
			errs = append(errs, fmt.Sprintf("Chinese document contains an untranslated English paragraph near line %d", p.num))
		}
	}
	return errs
}

func stubTableLanguageErrors(lines []string, fenced map[int]bool, lang string) []string {
	var errs []string
	for i, line := range lines {
		lineNo := i + 1
		s := strings.TrimSpace(line)
		if fenced[lineNo] || !strings.HasPrefix(s, "|") || strings.Count(s, "|") < 2 {
			continue
		}
		rawCells := strings.Split(strings.Trim(s, "|"), "|")
		cells := make([]string, len(rawCells))
		for ci, c := range rawCells {
			cells[ci] = cleanProseFragment(c)
		}
		if lang == "en" {
			for _, c := range cells {
				if han, _ := countHan(c); han > 0 {
					errs = append(errs, fmt.Sprintf("English table contains untranslated Han text at line %d", lineNo))
					break
				}
			}
			continue
		}
		isHeader := false
		if i+1 < len(lines) {
			next := strings.TrimSpace(lines[i+1])
			if strings.HasPrefix(next, "|") {
				sepCells := strings.Split(strings.Trim(next, "|"), "|")
				isHeader = len(sepCells) > 0
				for _, sc := range sepCells {
					if !pairTableSepRE.MatchString(sc) {
						isHeader = false
						break
					}
				}
			}
		}
		for _, c := range cells {
			if han, _ := countHan(c); han > 0 {
				continue
			}
			hit := pairTableProseRE.MatchString(c)
			if !hit && isHeader {
				for _, w := range pairEnglishWordRE.FindAllString(c, -1) {
					if _, ok := translatableTableHeaders[strings.ToLower(w)]; ok {
						hit = true
						break
					}
				}
			}
			if hit {
				errs = append(errs, fmt.Sprintf("Chinese table contains untranslated English text at line %d", lineNo))
				break
			}
		}
	}
	return errs
}

func isHanRune(r rune) bool {
	return (r >= 0x3400 && r <= 0x4dbf) || (r >= 0x4e00 && r <= 0x9fff) || (r >= 0xf900 && r <= 0xfaff)
}

func countHan(s string) (total, maxRun int) {
	run := 0
	for _, r := range s {
		if isHanRune(r) {
			total++
			run++
			if run > maxRun {
				maxRun = run
			}
		} else {
			run = 0
		}
	}
	return total, maxRun
}

func countEnglishLetters(s string) int {
	n := 0
	for _, w := range pairEnglishWordRE.FindAllString(s, -1) {
		for _, r := range w {
			if unicode.IsLetter(r) {
				n++
			}
		}
	}
	return n
}

func isUniformRune(s string, target rune) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r != target {
			return false
		}
	}
	return true
}

func stringMultisetEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	counts := make(map[string]int, len(a))
	for _, v := range a {
		counts[v]++
	}
	for _, v := range b {
		counts[v]--
		if counts[v] < 0 {
			return false
		}
	}
	return true
}

func codeBlocksEqual(a, b []stubCodeBlock) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
