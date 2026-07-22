## Usage

Sources:

- [Official README](https://github.com/fabriziomello/pg_normalize_query/blob/bd73d4979dbdb89e152b2a776390369992c8aa42/README.md)
- [Official extension SQL](https://github.com/fabriziomello/pg_normalize_query/blob/bd73d4979dbdb89e152b2a776390369992c8aa42/pg_normalize_query--1.0.sql)
- [Official implementation](https://github.com/fabriziomello/pg_normalize_query/blob/bd73d4979dbdb89e152b2a776390369992c8aa42/pg_normalize_query.c)

`pg_normalize_query` replaces constants in a valid SQL statement with numbered parameters, producing a fingerprint-like text form similar to the normalization used by `pg_stat_statements`. It is useful for grouping statement shapes in logs or application telemetry.

### Core Workflow

```sql
CREATE EXTENSION pg_normalize_query;

SELECT pg_normalize_query(
  $$SELECT * FROM pg_proc WHERE proname = 'pg_normalize_query'$$
);
-- SELECT * FROM pg_proc WHERE proname = $1

SELECT pg_normalize_query(
  $$SELECT oid FROM pg_class WHERE relkind IN ('r', 'i') LIMIT 10$$
);
-- SELECT oid FROM pg_class WHERE relkind IN ($1, $2) LIMIT $3
```

### API

The extension exposes one strict, immutable function: `pg_normalize_query(query text) RETURNS text`. The implementation invokes PostgreSQL's raw parser, walks constant locations in the parse tree, and rewrites those token spans as `$1`, `$2`, and so on. Existing external parameters are accounted for when numbering replacements.

### Interpretation and Compatibility

Input must be syntactically valid for the PostgreSQL parser used to build the extension; parse errors are raised normally. The result preserves most original spelling and whitespace, so it is a normalization of constants rather than a canonical SQL formatter. The implementation also notes that equivalent queries can still yield different normalized representations.

The reviewed upstream snapshot targeted PostgreSQL `9.4` through pre-release `13`. Because it copies parser-adjacent logic from a historical PostgreSQL/libpg_query implementation, verify both compilation and representative syntax on newer server majors before relying on stable fingerprints across upgrades.
