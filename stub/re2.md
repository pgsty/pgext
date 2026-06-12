## Usage

Sources: [official README](https://github.com/ClickHouse/pg_re2/blob/main/README.md), [official reference doc](https://github.com/ClickHouse/pg_re2/blob/main/doc/re2.md), [v0.3.0 release](https://github.com/ClickHouse/pg_re2/releases/tag/v0.3.0)

`re2` provides ClickHouse-compatible regular expression functions backed by Google's RE2 engine. It exposes both `text` and `bytea` overloads, so binary data with `\\0` bytes can be searched too. Pigsty packages version `0.3.0` for PostgreSQL 16-18 while upstream documents PostgreSQL 13+ support.

```sql
CREATE EXTENSION re2;

SELECT re2match('hello world', 'h.*o');
SELECT re2extract('Order #123', '(\\d+)');
SELECT re2countmatches('a1 b2 c3', '\\d');
```

### Core Functions

- `re2match(haystack, pattern) -> boolean`
- `re2extract(haystack, pattern) -> text|bytea`
- `re2extractall(haystack, pattern) -> text[]|bytea[]`
- `re2regexpextract(haystack, pattern, index default 1) -> text|bytea`
- `re2extractgroups(haystack, pattern) -> text[]|bytea[]`
- `re2extractallgroupsvertical(haystack, pattern) -> text[]|bytea[]`
- `re2extractallgroupshorizontal(haystack, pattern) -> text[]|bytea[]`
- `re2regexpquotemeta(haystack) -> text|bytea`
- `re2splitbyregexp(pattern, haystack, max_substrings default 0) -> text[]|bytea[]`
- `re2replaceregexpone(haystack, pattern, replacement) -> text|bytea`
- `re2replaceregexpall(haystack, pattern, replacement) -> text|bytea`
- `re2countmatches(...)` and `re2countmatchescaseinsensitive(...)`

```sql
SELECT re2extractallgroupsvertical('a=1 b=2', '(\\w)=(\\d)');
SELECT re2regexpquotemeta('a+b?');
SELECT re2splitbyregexp('\\s+', 'one two three', 2);
```

### Multi-Pattern Matching

The `re2multimatch*` family accepts either multiple pattern arguments or a `VARIADIC` array:

```sql
SELECT re2multimatchany('error: timeout', 'timeout', 'denied');
SELECT re2multimatchanyindex('error: timeout', VARIADIC ARRAY['timeout', 'denied']);
SELECT re2multimatchallindices('error: timeout', 'error', 'timeout', 'panic');
```

### Matching Semantics

- To match ClickHouse behavior, `.` matches line breaks by default.
- Prefix the pattern with `(?-s)` if you want `.` not to cross line breaks.
- Replacement strings support `\\0` through `\\9` backreferences.

### Caveats

- Upstream requires the system `re2` library at build/install time.
- Release `v0.3.0` uses SQL version `0.3`; run `ALTER EXTENSION re2 UPDATE TO '0.3'` after replacing extension binaries from an older minor release.
- `re2splitbyregexp` changed argument order in `v0.3.0` to `pattern, haystack[, max_substrings]`, matching ClickHouse. Earlier `0.2.0` builds used `haystack, pattern`.
- Upstream treats patch releases as binary-only, but minor releases can require SQL upgrade scripts.
