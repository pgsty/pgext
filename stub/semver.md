## Usage

Sources:

- [doc/semver.md](https://api.pgxn.org/src/semver/semver-0.43.0/doc/semver.md)
- [Changes](https://api.pgxn.org/src/semver/semver-0.43.0/Changes)
- [semver.control](https://api.pgxn.org/src/semver/semver-0.43.0/semver.control)
- [sql/semver.sql](https://api.pgxn.org/src/semver/semver-0.43.0/sql/semver.sql)

`semver` provides a Semantic Versioning 2.0.0 data type, comparison operators and the `semverrange` range type. Version 0.43.0 keeps the hash behavior introduced in 0.42.0; indexes created under older hash semantics still need the historical REINDEX.

### Core Workflow

```sql
CREATE EXTENSION semver;
SELECT '1.2.0'::semver > '1.2.0-b1'::semver;
SELECT '1.2.0'::semver = '1.2.0+build1'::semver;
SELECT to_semver('1.0'), is_semver('1.2.0');
SELECT get_semver_major('4.2.0'::semver);
SELECT '1.0.5'::semver <@ '[1.0.0,2.0.0)'::semverrange;
```

### Operational Boundaries

`semver(text)` and casts from `text` require valid semantic versions; `to_semver(text)` performs more permissive conversion, and `is_semver(text)` checks validity. `get_semver_major`, `get_semver_minor` and `get_semver_patch` return integer components; `get_semver_prerelease` returns the prerelease text. Casts also accept `numeric`, `real`, `double precision`, `integer`, `bigint` and `smallint`. Version 0.43.0 retains the 0.42.0 rejection of Infinity and NaN instead of producing malformed versions.

The comparison operators `=`, `<>`, `<`, `<=`, `>` and `>=`, plus `MIN(semver)` and `MAX(semver)`, use semantic ordering. A prerelease sorts before its final release; build metadata does not affect precedence. Btree and hash indexes are available. `semverrange` supports standard range operations.

Install this native, untrusted extension with administrator privileges. After installing 0.43.0, run `ALTER EXTENSION semver UPDATE` and use REINDEX on every existing hash index containing semver keys. The hash change is not repaired by updating the extension’s SQL definitions alone. The release also fixes ordering of alphanumeric prerelease identifiers beginning with a digit and validation of strings with multiple plus signs. When upgrading from old releases, inspect the upstream compatibility notes for previously accepted invalid values.
