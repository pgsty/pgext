## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/README.md)
- [Exported Rust modules](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/src/lib.rs)
- [Case-conversion implementation](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/src/modules/case.rs)
- [Markdown implementation](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/src/modules/markdown.rs)
- [Cargo metadata](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/Cargo.toml)

`pg_str` is a Rust/pgrx collection of Laravel-inspired string helpers. Version `0.3.1` installs functions in the `public` schema, mostly with a `str_` prefix, for case conversion, searching, slicing, transliteration, inflection, padding, Markdown rendering, random strings, and UUIDs. It is convenient application-style text processing, not a replacement for PostgreSQL collation-aware operators or full-text search.

### Core Workflow

```sql
CREATE EXTENSION pg_str;

SELECT str_snake('Hello PostgreSQL');
SELECT str_contains_all(
  'postgres extensions',
  ARRAY['postgres', 'ext']
);
SELECT str_uuid();
```

### Function Groups

- Position and slicing: `str_after`, `str_before`, `str_before_last`, `str_start`, `str_substr`, `str_split`, `str_length`, `str_replace`, and `str_append`.
- Case and word forms: `str_lower`, `str_upper`, `str_title`, `str_camel`, `str_kebab`, `str_snake`, `str_studly`, `str_scream`, `str_slug`, `str_singular`, and `str_plural`.
- Tests and conversion: `str_contains`, `str_contains_all`, `str_is_ascii`, and `str_ascii`.
- Generation and presentation: `str_random`, `str_uuid`, `str_markdown`, `pad_left`, `pad_right`, and `pad_both`.

The README checklist mixes implemented and planned names. The current Rust modules are authoritative; verify an exported `#[pg_extern]` function and its generated SQL signature before depending on it.

### Behavioral and Compatibility Notes

- `str_markdown` converts Markdown to HTML and enables strikethrough, but it is not an HTML sanitizer. Treat untrusted input as active content at the rendering boundary and sanitize according to the application policy.
- Case conversion, word inflection, transliteration, and slug generation follow Rust library behavior rather than the database collation and locale. Test multilingual and domain-specific terms.
- `str_random` casts its signed length to an allocation size; reject negative or unreasonably large values before calling it. Do not assume a generated value is a password or token without evaluating the required security properties.
- Padding code mixes character counts and byte-oriented string operations in some paths. Test multibyte padding strings and lengths; prefer core PostgreSQL functions where Unicode behavior is critical.
- Version `0.3.1` pins pgrx `0.11.2` and declares PostgreSQL 11 through 15 build features. A binary built for one PostgreSQL major is not portable to another; rebuild and test for the exact server major, especially newer versions.
