## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/repology/postgresql-libversion/blob/215a2d54573deb560b5165eb681a766085603829/README.md)
- [Extension control file](https://github.com/repology/postgresql-libversion/blob/215a2d54573deb560b5165eb681a766085603829/libversion.control)
- [C implementation](https://github.com/repology/postgresql-libversion/blob/215a2d54573deb560b5165eb681a766085603829/libversion.c)
- [libversion dependency](https://github.com/repology/libversion)

`libversion` compares package-style version strings through the external libversion library. It provides comparison functions and a `versiontext` type whose equality and ordering use version semantics rather than lexical text order.

```sql
CREATE EXTENSION libversion;
SELECT version_compare2('1.10', '1.2');
SELECT '1.10'::versiontext > '1.2'::versiontext;
SELECT '1.0'::versiontext = '1.0.0'::versiontext;
```

`version_compare4` accepts a flag for each operand; helper functions expose libversion flags such as `VERSIONFLAG_P_IS_PATCH`, `VERSIONFLAG_LOWER_BOUND`, and `VERSIONFLAG_UPPER_BOUND`.

The extension requires the matching libversion development and runtime library. Its cataloged compatibility ends at PostgreSQL 16, so validate source and binary builds on newer majors. A `versiontext` index depends on this comparison behavior: keep library versions consistent across primary, replicas, dump/restore targets, and logical subscribers to avoid ordering disagreement.
