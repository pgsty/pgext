


## Usage

> [omni_regex: PCRE-compatible regular expressions](https://docs.omnigres.org/omni_regex/regex/)

The `omni_regex` extension provides PCRE2-based regular expression support with named capture groups.

### Operators

| Operator | Description                    |
|:---------|:-------------------------------|
| `~`      | Matches string against regex   |
| `!~`     | Non-matching operator          |
| `=~`     | Alternative matching (same as `~`) |

```sql
SELECT 'foo' ~ regex 'fo+';   -- true
SELECT 'bar' !~ regex 'foo';  -- true
```

### Functions

**`regex_match(text, regex)`** -- Returns captured groups from the first match:

```sql
SELECT regex_match('ABC123', '([A-Z]*)(\d+)');  -- {ABC,123}
```

**`regex_matches(text, regex)`** -- Returns all matches as a set of text arrays:

```sql
SELECT regex_matches('foo1bar', '(fo+|bar)(\d?)');
-- {foo,1}
-- {bar,""}
```

**`regex_named_groups(regex)`** -- Extracts named capture groups:

```sql
SELECT index FROM regex_named_groups('(fo+|bar)(?<num>\d?)')
WHERE name = 'num';  -- 2
```

Named capture group syntax: `(?<name>REGEX)`
