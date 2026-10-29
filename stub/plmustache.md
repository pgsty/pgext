## Usage

Sources:

- [README.md](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/README.md)
- [LICENSE](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/LICENSE)
- [Makefile](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/Makefile)
- [plmustache.control.in](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/plmustache.control.in)
- [sql/plmustache.sql](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/sql/plmustache.sql)
- [src/plmustache.c](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/src/plmustache.c)
- [.github/workflows/ci.yaml](https://github.com/PostgREST/plmustache/blob/024b550d4fe2bbb4e122d2b3216a9ab7f514837a/.github/workflows/ci.yaml)

`plmustache` renders Mustache templates through PostgreSQL functions. The current source Makefile generates control and installation SQL version 0.1; the historical release tag v0.1 instead contains SQL version 0.0.1. This page describes the pinned source snapshot, without claiming a supported upgrade between those versions. Its C library requires libmustach. Current CI covers PostgreSQL 14–17; the older README lists 12–16.

### Render a Template

After the matching generated control, SQL and library files are installed, a superuser creates the extension. No shared preload or server restart is required for ordinary use:

```sql
CREATE EXTENSION plmustache;
CREATE FUNCTION greeting(person text) RETURNS text
LANGUAGE plmustache AS $$Hello {{person}}!$$;
SELECT greeting('PostgreSQL');
```

The extension's control does not mark installation trusted; the installed procedural language itself is TRUSTED. These are separate permission boundaries. Named function arguments provide template variables, with returned output as text. The library supplies call, inline and validation handlers, but the inline handler rejects DO blocks.

### Escaping and Sections

Double braces escape HTML; triple braces return the raw value. For example:

```sql
CREATE FUNCTION display_value(value text) RETURNS text
LANGUAGE plmustache AS $$Escaped: {{value}} Raw: {{{value}}}$$;
CREATE FUNCTION visible_message(show bool) RETURNS text
LANGUAGE plmustache AS $${{#show}}Visible{{/show}}{{^show}}Hidden{{/show}}$$;
SELECT display_value('<script>example</script>');
SELECT visible_message(false);
```

Treat unescaped output as untrusted content at the rendering boundary. HTML escaping is not SQL quoting, JSON encoding or a general injection defense. The README demonstrates boolean sections and array iteration, while its roadmap is incomplete; verify exact argument types against the source tests. Partials, inheritance and lambdas are not established supported features in this snapshot.

### Maintenance

Keep generated SQL/control versions and the installed C library aligned. Existing backends retain a loaded old library until they disconnect. The source layout and version have moved since the old release; review the installation and any migration in an isolated database instead of assuming the old v0.1 artifact installs the current SQL version.
