/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"fmt"
	"strings"
)

// Markdown shortcode generators for Hugo/Hextra theme

// Badge generates badge shortcode
func Badge(content, color, alt, link string) string {
	parts := []string{fmt.Sprintf(`{{< badge content="%s"`, content)}

	if color != "" && color != "default" {
		parts = append(parts, fmt.Sprintf(`color="%s"`, color))
	}
	if alt != "" {
		parts = append(parts, fmt.Sprintf(`alt="%s"`, alt))
	}
	if link != "" {
		parts = append(parts, fmt.Sprintf(`link="%s"`, link))
	}

	return strings.Join(parts, " ") + " >}}"
}

// BgShortcode generates bg shortcode - a simplified badge with only 3 positional parameters
// Parameters: text, alt, color
func BgShortcode(text, alt, color string) string {
	// Build the shortcode based on what parameters are provided
	if alt == "" && color == "" {
		// Only text
		return fmt.Sprintf(`{{< bg "%s" >}}`, text)
	} else if color == "" {
		// Text and alt
		return fmt.Sprintf(`{{< bg "%s" "%s" >}}`, text, alt)
	} else {
		// All three parameters
		return fmt.Sprintf(`{{< bg "%s" "%s" "%s" >}}`, text, alt, color)
	}
}

// PkgShortcode generates package shortcode with 5 parameters:
// name, version, org, state, count
func PkgShortcode(name, version, org, state, count string) string {
	parts := []string{fmt.Sprintf(`{{< pkg "%s"`, name)}

	// Always include all 5 parameters for consistency
	if version != "" {
		parts = append(parts, fmt.Sprintf(`"%s"`, version))
	} else {
		parts = append(parts, `""`)
	}

	if org != "" {
		parts = append(parts, fmt.Sprintf(`"%s"`, org))
	} else {
		parts = append(parts, `""`)
	}

	if state != "" {
		parts = append(parts, fmt.Sprintf(`"%s"`, state))
	} else {
		parts = append(parts, `""`)
	}

	if count != "" {
		parts = append(parts, fmt.Sprintf(`"%s"`, count))
	} else {
		parts = append(parts, `""`)
	}

	return strings.Join(parts, " ") + " >}}"
}

// ExtShortcode generates extension link shortcode
func ExtShortcode(target, label string) string {
	if target == "" {
		return `{{< ext "" >}}`
	}

	// Escape quotes in target and label
	target = strings.ReplaceAll(target, `"`, `\"`)

	if label != "" && label != target {
		label = strings.ReplaceAll(label, `"`, `\"`)
		return fmt.Sprintf(`{{< ext "%s" "%s" >}}`, target, label)
	}

	return fmt.Sprintf(`{{< ext "%s" >}}`, target)
}

// AliasShortcode generates alias shortcode (similar to ext but for listing pages)
func AliasShortcode(name, pkg string) string {
	if name == "" {
		return `{{< alias "" >}}`
	}

	// Escape quotes
	name = strings.ReplaceAll(name, `"`, `\"`)

	if pkg != "" && pkg != name {
		pkg = strings.ReplaceAll(pkg, `"`, `\"`)
		return fmt.Sprintf(`{{< alias "%s" "%s" >}}`, name, pkg)
	}

	return fmt.Sprintf(`{{< alias "%s" >}}`, name)
}

// CategoryShortcode generates category badge shortcode
func CategoryShortcode(category string) string {
	if category == "" {
		return ""
	}
	return fmt.Sprintf(`{{< category "%s" >}}`, strings.ToUpper(category))
}

// LicenseShortcode generates license badge shortcode
func LicenseShortcode(license string) string {
	if license == "" {
		return ""
	}
	return fmt.Sprintf(`{{< license "%s" >}}`, license)
}

// LanguageShortcode generates language badge shortcode
func LanguageShortcode(lang string) string {
	if lang == "" {
		return ""
	}
	return fmt.Sprintf(`{{< language "%s" >}}`, lang)
}

// PGVerShortcode generates PostgreSQL version badges
func PGVerShortcode(versions []int, supported []int) string {
	if len(versions) == 0 {
		return Badge("N/A", "gray", "", "")
	}

	// Build comma-separated version list
	versionStrs := make([]string, len(versions))
	for i, v := range versions {
		versionStrs[i] = fmt.Sprintf("%d", v)
	}
	versionList := strings.Join(versionStrs, ",")

	// Build comma-separated color list (g=green for supported, r=red for unsupported)
	supportedMap := make(map[int]bool)
	for _, v := range supported {
		supportedMap[v] = true
	}

	colors := make([]string, len(versions))
	for i, v := range versions {
		if supportedMap[v] {
			colors[i] = "g"
		} else {
			colors[i] = "r"
		}
	}
	colorList := strings.Join(colors, ",")

	return fmt.Sprintf(`{{< pgver "%s" "%s" >}}`, versionList, colorList)
}

// CardsShortcode wraps content in cards shortcode
func CardsShortcode(cols int, content string) string {
	return fmt.Sprintf("{{< cards cols=%d >}}\n%s\n{{< /cards >}}", cols, content)
}

// CardShortcode generates a single card
func CardShortcode(title, subtitle, link, icon string) string {
	parts := []string{}

	if link != "" {
		parts = append(parts, fmt.Sprintf(`link="%s"`, link))
	}
	if title != "" {
		parts = append(parts, fmt.Sprintf(`title="%s"`, title))
	}
	if icon != "" {
		parts = append(parts, fmt.Sprintf(`icon="%s"`, icon))
	}
	if subtitle != "" {
		parts = append(parts, fmt.Sprintf(`subtitle="%s"`, subtitle))
	}

	return fmt.Sprintf(`{{< card %s >}}`, strings.Join(parts, " "))
}

// TabsShortcode generates tabs container
func TabsShortcode(items []string, content string) string {
	itemList := strings.Join(items, ",")
	return fmt.Sprintf(`{{< tabs items="%s" >}}%s{{< /tabs >}}`, itemList, content)
}

// TabShortcode generates individual tab content
func TabShortcode(content string) string {
	return fmt.Sprintf("\n{{< tab >}}\n%s\n{{< /tab >}}", content)
}

// TripleQuote wraps content in triple backticks (```) for code blocks
func TripleQuote(content string) string {
	return "```\n" + content + "\n```"
}

// TripleQuoteLang wraps content in triple backticks with language specification
func TripleQuoteLang(lang, content string) string {
	return fmt.Sprintf("```%s\n%s\n```", lang, content)
}

// TripleQuoteBash is a convenience function for bash code blocks
func TripleQuoteBash(content string) string {
	return TripleQuoteLang("bash", content)
}

// TripleQuoteSQL is a convenience function for SQL code blocks
func TripleQuoteSQL(content string) string {
	return TripleQuoteLang("sql", content)
}