## Usage

Sources:

- [Official upstream documentation](https://github.com/eendroroy/pg_vercomp/blob/767244c2ab544bbe660e85a4d61cae774ab61441/README.md)
- [Official extension SQL](https://github.com/eendroroy/pg_vercomp/blob/767244c2ab544bbe660e85a4d61cae774ab61441/vercomp--1.0.0.sql)
- [Official range-operator regression tests](https://github.com/eendroroy/pg_vercomp/blob/767244c2ab544bbe660e85a4d61cae774ab61441/test/sql/vercomp_test_caret.sql)

`vercomp` 1.0.0 adds a `version` base type, comparisons, B-tree and hash operator classes, and semver-like range operators. Upstream explicitly marks the project obsolete and points users to `pg_semver`. Use `vercomp` only to maintain data already committed to its exact parsing and ordering semantics.

### Store and Compare Versions

Create the extension before defining columns of its type:

```sql
CREATE EXTENSION vercomp;

CREATE TABLE releases (
  package text NOT NULL,
  release version NOT NULL
);

INSERT INTO releases VALUES
  ('demo', '1.2.3'),
  ('demo', '2.0.0-rc1'),
  ('demo', '2.0.0');

SELECT * FROM releases ORDER BY release;
SELECT version_cmp('1.2.3', '2.0.0');
SELECT version_bump('1.2.3', 2);
```

`version_cmp(version, version)` returns a negative, zero, or positive integer. `version_bump(version, integer)` increments component 0, 1, or 2. Ordinary `=`, `<>`, `<`, `<=`, `>`, and `>=` operators support indexed comparison.

### Range Operators and Casts

`~` and `!~` implement the project's tilde-style match; `^` and `!^` implement its caret-style match. These are binary operators between `version` values, not string prefix or regular-expression operators. The extension also installs implicit casts between text and `version` and an implicit `version`-to-integer cast, plus B-tree and hash operator classes.

Do not assume npm, Cargo, or another semver library will produce the same ranges. The reviewed zero-major regression tests produce surprising matches relative to the README's formulas, so build fixtures for every range used by the application. The integer conversion is a custom, lossy encoding and is not a stable monotonic representation for external systems.

### Migration Boundary

The README includes a legacy preload instruction, but the extension's user-visible type and operators are registered by its SQL script; validate the packaged loading procedure on the exact server rather than copying old configuration blindly. Before replacing it with `pg_semver`, compare accepted syntax, prerelease ordering, casts, range operators, indexes, and dump/restore output. A type change requires an explicit data migration and index rebuild, not just swapping a shared library.
