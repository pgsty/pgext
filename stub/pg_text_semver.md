## Usage

- Source: [GitHub repo](https://github.com/bigsmoke/pg_text_semver), [README](https://github.com/bigsmoke/pg_text_semver/blob/master/README.md)
- `pg_text_semver` implements Semantic Versioning 2.0.0 on top of PostgreSQL `text`, using a `semver` domain and related helper types/functions.

```sql
CREATE EXTENSION pg_text_semver;
```

## Core Workflow

The README highlights these capabilities:

- compare `semver` values with the usual comparison operators
- call `semver_cmp(semver, semver)` directly
- validate and inspect version strings with `semver_regexp()`
- cast parsed values to `semver_parsed` for sorting and indexing
- use `semver_prerelease` for prerelease validation and comparison

## Examples

The upstream README points users to the `test__pg_text_semver()` procedure for concrete examples of the types, operators, and functions. It also notes that the extension ships a separate `semver_parsed` type that can be serialized back to `semver` or `text`.

## Notes

The README contrasts this project with C-based `semver` extensions: `pg_text_semver` stays on `text`-backed domains and focuses on a simple, spec-oriented implementation.
