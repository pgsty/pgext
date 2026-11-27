package repo

import (
	"net/url"
	"slices"
	"strings"
)

type sourceDefinition struct {
	key      string
	provider string
	active   bool
}

type sourceStanza struct {
	content     string
	definitions []sourceDefinition
}

// Recognize the same source on official endpoints and supported mirrors.
// An unrecognized endpoint can conflict, but cannot qualify for migration.
func sourceURL(endpoint string) (identity, provider string) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", ""
	}
	host, path := strings.ToLower(u.Host), strings.TrimRight(u.Path, "/")
	identity = host + path
	if u.User != nil || u.RawQuery != "" || u.Fragment != "" || (u.Scheme != "http" && u.Scheme != "https") {
		return identity, ""
	}
	switch host {
	case "apt.postgresql.org":
		if path == "/pub/repos/apt" {
			return "pgdg/apt", "pgdg"
		}
	case "download.postgresql.org", "yum.postgresql.org":
		if strings.HasPrefix(path, "/pub/repos/yum/") {
			return "pgdg" + strings.TrimPrefix(path, "/pub/repos"), "pgdg"
		}
	case "repo.pigsty.io", "repo.pigsty.cc":
		if strings.HasPrefix(path, "/yum/pgdg/") {
			return "pgdg/yum" + strings.TrimPrefix(path, "/yum/pgdg"), "pgdg"
		}
		for _, prefix := range []string{"/apt/infra", "/apt/pgsql", "/yum/infra", "/yum/pgsql"} {
			if path == prefix || strings.HasPrefix(path, prefix+"/") {
				return "pigsty" + path, "pigsty"
			}
		}
	case "mirrors.cloud.tencent.com", "mirrors.tuna.tsinghua.edu.cn", "mirrors.ustc.edu.cn", "mirrors.xtom.de":
		if path == "/postgresql/repos/apt" || strings.HasPrefix(path, "/postgresql/repos/yum/") {
			return "pgdg" + strings.TrimPrefix(path, "/postgresql/repos"), "pgdg"
		}
	}
	switch path {
	case "/ubuntu", "/ubuntu-ports", "/debian", "/debian-security":
		// A native source on another mirror already supplies the same suite.
		return path, ""
	}
	return identity, ""
}

func aptSource(endpoint, suite string, active bool) sourceDefinition {
	identity, provider := sourceURL(endpoint)
	return sourceDefinition{key: identity + " " + suite, provider: provider, active: active}
}

func sourceStanzas(path, content, family string) []sourceStanza {
	if family == "rpm" {
		return rpmSources(content)
	}
	if strings.HasSuffix(path, ".sources") {
		return deb822Sources(content)
	}
	var stanzas []sourceStanza
	for line := range strings.SplitAfterSeq(content, "\n") {
		stanza := sourceStanza{content: line}
		fields := strings.Fields(line)
		if len(fields) == 0 || strings.HasPrefix(fields[0], "#") {
			stanzas = append(stanzas, stanza)
			continue
		}
		if fields[0] != "deb" && fields[0] != "deb-src" {
			stanza.definitions = []sourceDefinition{{}}
			stanzas = append(stanzas, stanza)
			continue
		}
		fields = fields[1:]
		if len(fields) > 0 && strings.HasPrefix(fields[0], "[") {
			for len(fields) > 0 {
				end := strings.HasSuffix(fields[0], "]")
				fields = fields[1:]
				if end {
					break
				}
			}
		}
		if len(fields) < 2 {
			stanza.definitions = []sourceDefinition{{}}
		} else {
			stanza.definitions = []sourceDefinition{aptSource(fields[0], fields[1], true)}
		}
		stanzas = append(stanzas, stanza)
	}
	return stanzas
}

func deb822Sources(content string) []sourceStanza {
	var stanzas []sourceStanza
	fields := map[string]string{}
	start, offset := 0, 0
	last := ""
	malformed := false
	flush := func() {
		if start == offset {
			return
		}
		stanza := sourceStanza{content: content[start:offset]}
		types, endpoints, suites := strings.Fields(fields["types"]), strings.Fields(fields["uris"]), strings.Fields(fields["suites"])
		if len(fields) > 0 || malformed {
			if malformed || len(types) == 0 || len(endpoints) == 0 || len(suites) == 0 ||
				slices.ContainsFunc(types, func(s string) bool { return s != "deb" && s != "deb-src" }) {
				stanza.definitions = []sourceDefinition{{}}
			} else {
				for _, endpoint := range endpoints {
					for _, suite := range suites {
						stanza.definitions = append(stanza.definitions, aptSource(endpoint, suite, !strings.EqualFold(fields["enabled"], "no")))
					}
				}
			}
		}
		stanzas = append(stanzas, stanza)
		clear(fields)
		start, last, malformed = offset, "", false
	}
	for line := range strings.SplitAfterSeq(content, "\n") {
		offset += len(line)
		switch {
		case strings.TrimSpace(line) == "":
			flush()
		case strings.HasPrefix(strings.TrimSpace(line), "#"):
		case last != "" && (strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t")):
			fields[last] += " " + strings.TrimSpace(line)
		default:
			if name, value, ok := strings.Cut(line, ":"); ok {
				last = strings.ToLower(strings.TrimSpace(name))
				fields[last] = strings.TrimSpace(value)
			} else {
				malformed = true
			}
		}
	}
	flush()
	return stanzas
}

func rpmSources(content string) []sourceStanza {
	var stanzas []sourceStanza
	start, offset := 0, 0
	var name, last string
	var endpoints []string
	alternate := false
	flush := func() {
		if start == offset {
			return
		}
		stanza := sourceStanza{content: content[start:offset]}
		if name != "" {
			key := name
			provider := repoProvider(strings.TrimPrefix(name, "pgext-"))
			if provider != "" {
				key = strings.TrimPrefix(name, "pgext-")
				if alternate || len(endpoints) == 0 || slices.ContainsFunc(endpoints, func(endpoint string) bool {
					_, source := sourceURL(endpoint)
					return source != provider
				}) {
					provider = ""
				}
			}
			stanza.definitions = []sourceDefinition{{key: key, provider: provider, active: true}}
		} else if alternate {
			stanza.definitions = []sourceDefinition{{}}
		}
		stanzas = append(stanzas, stanza)
		start = offset
		name, last, endpoints, alternate = "", "", nil, false
	}
	for raw := range strings.SplitAfterSeq(content, "\n") {
		line := strings.TrimSpace(raw)
		if strings.HasPrefix(line, "[") {
			flush()
		}
		offset += len(raw)
		switch {
		case line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";"):
		case strings.HasPrefix(line, "["):
			var closed bool
			name, _, closed = strings.Cut(line[1:], "]")
			if !closed {
				name, alternate = "", true
			}
		case name == "":
			alternate = true
		default:
			key, value, ok := strings.Cut(line, "=")
			if !ok {
				if last == "baseurl" && (raw[0] == ' ' || raw[0] == '\t') {
					endpoints = append(endpoints, strings.Fields(line)...)
				} else {
					alternate = true
				}
				continue
			}
			last = strings.ToLower(strings.TrimSpace(key))
			switch last {
			case "baseurl":
				endpoints = append(endpoints, strings.Fields(value)...)
			case "metalink", "mirrorlist", "include":
				alternate = true
			}
		}
	}
	flush()
	return stanzas
}
