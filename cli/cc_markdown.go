/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

CC (pigsty.cc) Markdown generators - Tailwind-inspired inline CSS badges for Hugo/Docsy
Modern, clean design with semantic colors
*/
package cli

import (
	"fmt"
	"strings"
)

// CCBaseURL is the base URL path for extension pages on pigsty.cc
const CCBaseURL = "/ext"

// Tailwind-inspired color palette
var Colors = map[string]string{
	// Semantic colors
	"green":  "#22c55e", // Success, available
	"blue":   "#3b82f6", // Info, PGDG
	"red":    "#ef4444", // Error, missing, PIGSTY
	"amber":  "#f59e0b", // Warning
	"gray":   "#6b7280", // Neutral
	"purple": "#a855f7", // Special
	"cyan":   "#06b6d4", // Alternate
	"rose":   "#f43f5e", // Danger
	"teal":   "#14b8a6", // Alternate success
	"indigo": "#6366f1", // Deep blue
	"orange": "#f97316", // Alert
	"slate":  "#64748b", // Muted
}

// Category colors - vibrant professional palette
var CategoryColors = map[string]string{
	"TIME":  "#0ea5e9", // Sky blue
	"GIS":   "#22c55e", // Green
	"RAG":   "#a855f7", // Purple
	"FTS":   "#f43f5e", // Rose
	"OLAP":  "#f97316", // Orange
	"FEAT":  "#14b8a6", // Teal
	"LANG":  "#6366f1", // Indigo
	"TYPE":  "#06b6d4", // Cyan
	"UTIL":  "#64748b", // Slate
	"FUNC":  "#3b82f6", // Blue
	"ADMIN": "#ef4444", // Red
	"STAT":  "#8b5cf6", // Violet
	"SEC":   "#ea580c", // Orange dark
	"FDW":   "#475569", // Slate dark
	"SIM":   "#d97706", // Amber
	"ETL":   "#71717a", // Zinc
}

// Language colors
var LanguageColors = map[string]string{
	"C":       "#555555",
	"C++":     "#f34b7d",
	"Rust":    "#dea584",
	"SQL":     "#e38c00",
	"PLpgSQL": "#336791",
	"Perl":    "#0298c3",
	"Python":  "#3572a5",
	"Java":    "#b07219",
	"Go":      "#00add8",
	"Data":    "#1e293b",
}

// License colors
var LicenseColors = map[string]string{
	"PostgreSQL":   "#336791",
	"MIT":          "#22c55e",
	"Apache-2.0":   "#ef4444",
	"BSD-2-Clause": "#3b82f6",
	"BSD-3-Clause": "#3b82f6",
	"GPL-2.0":      "#8b5cf6",
	"GPL-3.0":      "#8b5cf6",
	"LGPL-2.1":     "#a855f7",
	"LGPL-3.0":     "#a855f7",
	"AGPLv3":       "#7c3aed",
	"ISC":          "#10b981",
	"MPL-2.0":      "#f97316",
	"Timescale":    "#fbbf24",
	"Elastic-2.0":  "#06b6d4",
}

// CCColorBadge generates a colored badge with semantic color name
func CCColorBadge(text, color string) string {
	bgColor := Colors[color]
	if bgColor == "" {
		bgColor = Colors["gray"]
	}
	style := fmt.Sprintf(
		"display:inline-block;padding:0.2em 0.5em;font-size:0.8em;font-weight:600;"+
			"line-height:1.2;color:#fff;background:%s;border-radius:0.375rem;",
		bgColor)
	return fmt.Sprintf(`<span style="%s">%s</span>`, style, text)
}

// CCBadge generates a modern badge with optional link
func CCBadge(text, bgColor, link string) string {
	style := fmt.Sprintf(
		"display:inline-block;padding:0.2em 0.5em;font-size:0.8em;font-weight:600;"+
			"line-height:1.2;color:#fff;background:%s;border-radius:0.375rem;"+
			"text-decoration:none;",
		bgColor)
	if link != "" {
		return fmt.Sprintf(`<a href="%s" style="%s">%s</a>`, link, style, text)
	}
	return fmt.Sprintf(`<span style="%s">%s</span>`, style, text)
}

// CCCategoryBadge generates a category badge with link
func CCCategoryBadge(category string) string {
	if category == "" {
		return "-"
	}
	color := CategoryColors[strings.ToUpper(category)]
	if color == "" {
		color = Colors["gray"]
	}
	link := fmt.Sprintf("%s/list/cate#%s", CCBaseURL, strings.ToLower(category))
	return CCBadge(strings.ToUpper(category), color, link)
}

// CCLanguageBadge generates a language badge
func CCLanguageBadge(lang string) string {
	if lang == "" {
		return "-"
	}
	color := LanguageColors[lang]
	if color == "" {
		color = Colors["gray"]
	}
	anchor := strings.ToLower(strings.ReplaceAll(lang, "+", ""))
	link := fmt.Sprintf("%s/list/lang#%s", CCBaseURL, anchor)
	return CCBadge(lang, color, link)
}

// CCLicenseBadge generates a license badge
func CCLicenseBadge(license string) string {
	if license == "" {
		return "-"
	}
	color := LicenseColors[license]
	if color == "" {
		color = Colors["gray"]
	}
	anchor := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(license, ".", ""), "-", ""))
	link := fmt.Sprintf("%s/list/license#%s", CCBaseURL, anchor)
	return CCBadge(license, color, link)
}

// CCRepoBadge generates a repo badge (PGDG or PIGSTY)
func CCRepoBadge(repo string) string {
	if repo == "" {
		return "-"
	}
	repoLower := strings.ToLower(repo)
	color := Colors["gray"]
	if repoLower == "pgdg" {
		color = "#336791" // PostgreSQL blue
	} else if repoLower == "pigsty" {
		color = Colors["red"]
	}
	link := fmt.Sprintf("%s/repo/%s", CCBaseURL, repoLower)
	return CCBadge(strings.ToUpper(repo), color, link)
}

// CCPGBadge generates a compact PG version badge
func CCPGBadge(ver string) string {
	style := "display:inline-block;padding:0.15em 0.35em;font-size:0.75em;font-weight:600;" +
		"color:#fff;background:#336791;border-radius:0.25rem;margin:0 1px;"
	return fmt.Sprintf(`<span style="%s">%s</span>`, style, ver)
}

// CCPGBadgeGreen generates a green PG version badge (supported)
func CCPGBadgeGreen(ver string) string {
	style := fmt.Sprintf("display:inline-block;padding:0.15em 0.35em;font-size:0.75em;font-weight:600;"+
		"color:#fff;background:%s;border-radius:0.25rem;margin:0 1px;", Colors["green"])
	return fmt.Sprintf(`<span style="%s">%s</span>`, style, ver)
}

// CCPGBadgeRed generates a red PG version badge (not supported)
func CCPGBadgeRed(ver string) string {
	style := fmt.Sprintf("display:inline-block;padding:0.15em 0.35em;font-size:0.75em;font-weight:600;"+
		"color:#fff;background:%s;border-radius:0.25rem;margin:0 1px;", Colors["red"])
	return fmt.Sprintf(`<span style="%s">%s</span>`, style, ver)
}

// CCPGBadges generates a row of PG version badges
func CCPGBadges(versions []string) string {
	if len(versions) == 0 {
		return "-"
	}
	var badges []string
	for _, v := range versions {
		v = strings.TrimSuffix(v, "+")
		badges = append(badges, CCPGBadge(v))
	}
	return strings.Join(badges, "")
}

// CCExtBadge generates an extension link badge
func CCExtBadge(name string) string {
	if name == "" {
		return ""
	}
	style := "display:inline-block;padding:0.15em 0.4em;font-size:0.8em;" +
		"background:#f1f5f9;border-radius:0.25rem;text-decoration:none;color:#334155;margin:0.1em;"
	return fmt.Sprintf(`<a href="%s/e/%s" style="%s">%s</a>`, CCBaseURL, name, style, name)
}

// CCAvailBadge generates an availability badge
func CCAvailBadge(version, repo string) string {
	repoLower := strings.ToLower(repo)
	bgColor := Colors["green"]
	if repoLower == "pgdg" {
		bgColor = "#336791"
	} else if repoLower == "pigsty" {
		bgColor = Colors["red"]
	}
	style := fmt.Sprintf(
		"display:inline-block;padding:0.1em 0.3em;font-size:0.7em;font-weight:500;"+
			"color:#fff;background:%s;border-radius:0.2rem;",
		bgColor)
	return fmt.Sprintf(`<span style="%s">%s</span>`, style, version)
}

// CCIcon generates a Font Awesome icon
func CCIcon(name string) string {
	return fmt.Sprintf(`<i class="fab fa-%s"></i>`, name)
}

// CCExtLink generates a code-style extension link
func CCExtLink(name, pkg string) string {
	if name == "" {
		return ""
	}
	display := name
	if pkg != "" && pkg != name {
		display = pkg
	}
	return fmt.Sprintf("[`%s`](%s/e/%s)", display, CCBaseURL, name)
}

// CCOSLink generates a link to an OS page
func CCOSLink(os string) string {
	if os == "" {
		return ""
	}
	return fmt.Sprintf("[%s](%s/os/%s)", os, CCBaseURL, os)
}

// CCCategoryLink generates a category badge link
func CCCategoryLink(category string) string {
	return CCCategoryBadge(category)
}

// CCLanguageLink generates a language badge link
func CCLanguageLink(lang string) string {
	return CCLanguageBadge(lang)
}

// CCLicenseLink generates a license badge link
func CCLicenseLink(license string) string {
	return CCLicenseBadge(license)
}

// CCExtLinkWithLabel generates extension link with custom display label
func CCExtLinkWithLabel(name, label string) string {
	if name == "" {
		return ""
	}
	if label == "" {
		label = name
	}
	return fmt.Sprintf("[`%s`](%s/e/%s)", label, CCBaseURL, name)
}
