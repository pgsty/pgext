/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>

CC (pigsty.cc) Markdown generators - CSS class-based badges for Hugo/Docsy
Uses semantic CSS classes instead of inline styles
*/
package cli

import (
	"fmt"
	"strings"
)

// CCBaseURL is the base URL path for extension pages on pigsty.cc
const CCBaseURL = "/ext"

// CCCategoryBadge generates a category badge with CSS class
func CCCategoryBadge(category string) string {
	if category == "" {
		return "-"
	}
	lower := strings.ToLower(category)
	return fmt.Sprintf(`<a class="ext-badge ext-badge--cate %s" href="%s/cate/%s">%s</a>`,
		lower, CCBaseURL, lower, strings.ToUpper(category))
}

// CCLanguageBadge generates a language badge with CSS class
func CCLanguageBadge(lang string) string {
	if lang == "" {
		return "-"
	}
	anchor := strings.ToLower(strings.ReplaceAll(lang, "+", ""))
	return fmt.Sprintf(`<a class="ext-badge ext-badge--lang %s" href="%s/language#%s">%s</a>`,
		anchor, CCBaseURL, anchor, lang)
}

// CCLicenseBadge generates a license badge with CSS class
func CCLicenseBadge(license string) string {
	if license == "" {
		return "-"
	}
	cls := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(license, ".", ""), "-", ""))
	anchor := strings.ReplaceAll(cls, " ", "")
	return fmt.Sprintf(`<a class="ext-badge ext-badge--license %s" href="%s/license#%s">%s</a>`,
		cls, CCBaseURL, anchor, license)
}

// CCRepoBadge generates a repo badge (PGDG or PIGSTY) with CSS class
func CCRepoBadge(repo string) string {
	if repo == "" {
		return "-"
	}
	repoLower := strings.ToLower(repo)
	return fmt.Sprintf(`<a class="ext-badge ext-badge--repo %s" href="%s/repo#%s">%s</a>`,
		repoLower, CCBaseURL, repoLower, strings.ToUpper(repo))
}

// CCPGVerBadge generates a PG version badge with ok/na CSS class
func CCPGVerBadge(ver string, ok bool) string {
	cls := "ext-pgver--na"
	if ok {
		cls = "ext-pgver--ok"
	}
	return fmt.Sprintf(`<span class="ext-pgver %s">%s</span>`, cls, ver)
}

// CCFlagBadge generates a yes/no flag badge with CSS class
func CCFlagBadge(val bool) string {
	if val {
		return `<span class="ext-flag ext-flag--yes">是</span>`
	}
	return `<span class="ext-flag ext-flag--no">否</span>`
}

// CCOptFlagBadge generates a yes/no/unknown flag badge
func CCOptFlagBadge(val NullBool) string {
	if !val.Valid {
		return `<span class="ext-flag ext-flag--no">否</span>`
	}
	return CCFlagBadge(val.Bool)
}

// CCAvailBadge generates an availability badge with CSS class
func CCAvailBadge(version string) string {
	return fmt.Sprintf(`<span class="ext-badge ext-badge--avail">%s</span>`, version)
}

// CCMissBadge generates a missing badge
func CCMissBadge() string {
	return `<span class="ext-badge ext-badge--miss">✗</span>`
}

// CCExtLink generates a plain markdown extension link
func CCExtLink(name string) string {
	if name == "" {
		return ""
	}
	return fmt.Sprintf("[`%s`](%s/e/%s)", name, CCBaseURL, name)
}

// CCExtBoldLink generates a bold code extension link
func CCExtBoldLink(name, url string) string {
	if url != "" {
		return fmt.Sprintf("[**`%s`**](%s)", name, url)
	}
	return fmt.Sprintf("**`%s`**", name)
}

// CCOSLink generates a link to an OS page
func CCOSLink(os string) string {
	if os == "" {
		return ""
	}
	return fmt.Sprintf("[**%s**](%s/os/%s)", os, CCBaseURL, os)
}

// CCColorBadge kept for backward compatibility with list/os generators
func CCColorBadge(text, color string) string {
	Colors := map[string]string{
		"green":  "#22c55e",
		"blue":   "#3b82f6",
		"red":    "#ef4444",
		"amber":  "#f59e0b",
		"gray":   "#6b7280",
		"purple": "#a855f7",
	}
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

// CCBadge kept for backward compatibility
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

// CCPGBadge kept for backward compatibility
func CCPGBadge(ver string) string {
	style := "display:inline-block;padding:0.15em 0.35em;font-size:0.75em;font-weight:600;" +
		"color:#fff;background:#336791;border-radius:0.25rem;margin:0 1px;"
	return fmt.Sprintf(`<span style="%s">%s</span>`, style, ver)
}

// CCPGBadgeGreen kept for backward compatibility
func CCPGBadgeGreen(ver string) string {
	style := "display:inline-block;padding:0.15em 0.35em;font-size:0.75em;font-weight:600;" +
		"color:#fff;background:#22c55e;border-radius:0.25rem;margin:0 1px;"
	return fmt.Sprintf(`<span style="%s">%s</span>`, style, ver)
}

// CCPGBadgeRed kept for backward compatibility
func CCPGBadgeRed(ver string) string {
	style := "display:inline-block;padding:0.15em 0.35em;font-size:0.75em;font-weight:600;" +
		"color:#fff;background:#ef4444;border-radius:0.25rem;margin:0 1px;"
	return fmt.Sprintf(`<span style="%s">%s</span>`, style, ver)
}

// CCPGBadges kept for backward compatibility
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

// CCCategoryLink kept for backward compatibility
func CCCategoryLink(category string) string {
	return CCCategoryBadge(category)
}

// CCLanguageLink kept for backward compatibility
func CCLanguageLink(lang string) string {
	return CCLanguageBadge(lang)
}

// CCLicenseLink kept for backward compatibility
func CCLicenseLink(license string) string {
	return CCLicenseBadge(license)
}

// CCIcon generates a Font Awesome icon
func CCIcon(name string) string {
	return fmt.Sprintf(`<i class="fab fa-%s"></i>`, name)
}

// CCExtensionTable generates a markdown extension table for a list of extensions
// Used by both category pages and extension index page
func CCExtensionTable(exts []*Extension) string {
	var b strings.Builder
	b.WriteString("| **扩展** | **包** | **版本** | **许可证** | **语言** | **描述** |\n")
	b.WriteString("|:---------|:-------|:--------:|:----------:|:--------:|:---------|\n")
	for _, ext := range exts {
		version := "-"
		if ext.Version.Valid {
			version = ext.Version.String
		}
		license := "-"
		if ext.License.Valid {
			license = CCLicenseBadge(ext.License.String)
		}
		lang := "-"
		if ext.Lang.Valid {
			lang = CCLanguageBadge(ext.Lang.String)
		}
		desc := SanitizeText(ext.Name)
		if ext.ZhDesc.Valid && ext.ZhDesc.String != "" {
			desc = SanitizeText(ext.ZhDesc.String)
		} else if ext.EnDesc.Valid && ext.EnDesc.String != "" {
			desc = SanitizeText(ext.EnDesc.String)
		}
		pkgLink := fmt.Sprintf("`%s`", ext.Pkg)
		if ext.URL.Valid && ext.URL.String != "" {
			pkgLink = fmt.Sprintf("[`%s`](%s)", ext.Pkg, ext.URL.String)
		}
		b.WriteString(fmt.Sprintf("| [`%s`](/ext/e/%s) | %s | `%s` | %s | %s | %s |\n",
			ext.Name, ext.Name, pkgLink, version, license, lang, desc))
	}
	b.WriteString("{.ext-table}\n\n")
	return b.String()
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
