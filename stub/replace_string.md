## Usage

Sources:

- [Official README](https://github.com/helgeho/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/README.md)
- [Function implementation](https://github.com/helgeho/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/src/lib.rs)
- [Build feature matrix](https://github.com/helgeho/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/Cargo.toml)
- [Extension control file](https://github.com/helgeho/replace_string/blob/105b64f8012118edb0b9314e79723ddcba41ee64/replace_string.control)

`replace_string` provides `repl_str(text, text, text) → text`, replacing every occurrence of one string with another through Rust's regex library. The search string is escaped and therefore matched literally, but the replacement string still uses regex replacement expansion rules.

### Core Workflow

```sql
CREATE EXTENSION replace_string;

SELECT repl_str('h@llo', '@', 'e');
SELECT repl_str('a.c a.c', '.', '-');
```

Both examples perform global literal search matching. PostgreSQL already provides `replace(text, from, to)` for fully literal replacement, so prefer the built-in function unless this extension is required for an existing application.

### Replacement Semantics

The implementation applies `regex::escape` to the search argument and then calls `Regex::replace_all`. Regex metacharacters in the search text therefore have no special meaning. By contrast, `$name`, `$1`, and `$$` in the replacement text follow Rust regex capture-expansion syntax. Because the escaped search pattern has no capture groups, capture references may expand to empty text rather than remain literal. Use PostgreSQL's built-in `replace` when replacement text can contain dollar signs.

An empty search string produces an empty regular expression and inserts the replacement at every match boundary, which may be surprising and can greatly enlarge the result. Null SQL inputs follow the pgrx-generated function wrapper's strictness behavior for the built package; verify this if null handling is material.

### Compatibility

The extension has version `0.0.0`, requires superuser installation, and is not relocatable. The current Cargo features target PostgreSQL 13, 14, and 15, defaulting to 14, with `pgrx` 0.16.1 and regex 1.12.2. There is no preload or runtime configuration requirement.

Pin the exact package because regex replacement behavior is part of the result semantics. Test dollar signs, empty search strings, Unicode, large values, and target PostgreSQL compatibility before migrating data in bulk.
