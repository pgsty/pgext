## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/README.md)
- [Exported Rust functions](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/src/lib.rs)
- [Case-conversion module](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/src/modules/case.rs)
- [Cargo metadata](https://github.com/publicmatt/pg_str/blob/5c49c2397a987feb6bcdc248a53abaf4297b7e36/Cargo.toml)

`pg_str` is a Rust/pgrx collection of Laravel-inspired string helpers. Functions are installed in `public` with a `str_` prefix and cover case conversion, containment, slicing, Markdown rendering, random strings, and UUID generation.

```sql
CREATE EXTENSION pg_str;
SELECT str_snake('Hello PostgreSQL');
SELECT str_contains_all('postgres extensions', ARRAY['postgres', 'ext']);
SELECT str_uuid();
```

The README checklist is the best API inventory, but it includes both implemented and planned functions. Confirm a function in the current Rust modules before depending on it. Markdown rendering produces HTML, so escape or sanitize the result according to the consuming application's trust model.

Version 0.3.1 uses an older pgrx release and declares PostgreSQL 11–15 features. The GitHub repository is a mirror of the upstream Gitea project; validate builds on newer PostgreSQL versions rather than assuming ABI compatibility.
