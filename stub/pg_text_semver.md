## Usage

Source: [README](https://github.com/bigsmoke/pg_text_semver/blob/master/README.md), [META.json](https://github.com/bigsmoke/pg_text_semver/blob/master/META.json), [Tag v1.2.1](https://github.com/bigsmoke/pg_text_semver/tree/v1.2.1)

`pg_text_semver` implements Semantic Versioning 2.0.0 on top of PostgreSQL `text` using a `semver` domain rather than a custom C type.

### Core types and functions

```sql
CREATE EXTENSION pg_text_semver;

SELECT '1.2.3'::semver < '2.0.0'::semver;
SELECT semver_cmp('1.2.3'::semver, '1.2.4'::semver);
SELECT semver_regexp(true);
SELECT '1.2.3-alpha.1+build5'::semver::semver_parsed;
```

- `semver`: domain over `text` with SemVer validation.
- `semver_parsed`: parsed composite type that supports sorting and indexing.
- `semver_prerelease`: domain for prerelease identifiers.
- `semver_cmp(...)`: comparison function for `semver` and `semver_parsed`.
- `semver_regexp(include_captures boolean)`: exposes the validation regex.

### Extra helpers

The current README also documents PGXN-version-range helpers:

- `meta_pgxn_version_range(text)`
- `meta_pgxn_version_range_cmp(text, text)`
- `nonsemver_cmp(text, text, text)`

### Caveats

- This extension favors a spec-oriented, text-backed implementation over the lower-level C-based alternatives.
- The upstream README remains the authoritative user-facing reference; the current stub already matched that surface closely, so this refresh mainly aligns it with the documented 1.2.1 helper set.
