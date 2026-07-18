## Usage

Sources:

- [vercomp README at the reviewed commit](https://github.com/eendroroy/pg_vercomp/blob/767244c2ab544bbe660e85a4d61cae774ab61441/README.md)
- [vercomp 1.0.0 install SQL at the reviewed commit](https://github.com/eendroroy/pg_vercomp/blob/767244c2ab544bbe660e85a4d61cae774ab61441/vercomp--1.0.0.sql)
- [vercomp C source at the reviewed commit](https://github.com/eendroroy/pg_vercomp/blob/767244c2ab544bbe660e85a4d61cae774ab61441/vercomp.c)
- [Upstream-recommended pg_semver successor](https://github.com/eendroroy/pg_semver)

`vercomp` version 1.0.0 defines a `version` data type, casts, comparison and range operators, and default B-tree and hash operator classes. It supports ordinary comparisons plus the `~`, `!~`, `^`, and `!^` version-range operators, and exposes helpers including `version_cmp` and `version_bump`.

### Store and Compare Versions

```sql
CREATE EXTENSION vercomp;

CREATE TABLE component_versions (
  component text PRIMARY KEY,
  release version NOT NULL
);

INSERT INTO component_versions VALUES
  ('alpha', '1.2.3'),
  ('beta', '2.0.0-rc1'),
  ('gamma', '2.0.0');

SELECT component, release
FROM component_versions
ORDER BY release;

SELECT version_cmp('1.2.3', '2.0.0');
```

### Caveats

- Upstream marks the project obsolete, archived the repository, and recommends `pg_semver` as its successor. Prefer the maintained alternative for new deployments.
- The README contains a legacy instruction to add `$libdir/vercomp` to `shared_preload_libraries` and restart. However, the reviewed control/C source declares no `_PG_init`, hook, worker, or shared-memory setup; catalog evidence records no load requirement. Treat the README step as an unresolved documentation conflict, not as demonstrated runtime necessity.
- Upstream's environment tests target PostgreSQL 10 and do not establish compatibility with current releases.
- Existing indexes and columns depend on this extension's `version` type and operator classes; plan an explicit data conversion before replacing it.
