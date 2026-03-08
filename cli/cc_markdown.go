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

// CCPkgStateBadge generates an availability badge based on package state and version
func CCPkgStateBadge(state, version string) string {
	switch state {
	case "AVAIL":
		if version != "" {
			return CCAvailBadge(version)
		}
		return CCAvailBadge("✓")
	case "MISS":
		return CCMissBadge()
	case "HIDE":
		return `<span class="ext-badge ext-badge--hide">-</span>`
	case "THROW":
		return `<span class="ext-badge ext-badge--throw">!</span>`
	case "BREAK":
		return `<span class="ext-badge ext-badge--break">!</span>`
	default:
		return `<span class="ext-badge ext-badge--hide">-</span>`
	}
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
		desc := SanitizeText(ext.GetZhDesc())
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

