


## Usage

Sources:

- [PGXN biscuit 2.4.1](https://pgxn.org/dist/biscuit/2.4.1/)
- [Biscuit README](https://github.com/CrystallineCore/Biscuit)
- [Biscuit CHANGELOG](https://github.com/CrystallineCore/Biscuit/blob/main/CHANGELOG.md)
- [Biscuit documentation](https://biscuit.readthedocs.io/)

`biscuit` is a PostgreSQL index access method for accelerating `LIKE`, `NOT LIKE`, `ILIKE`, and `NOT ILIKE` pattern matching on text. It uses bitmap-style position indexes to avoid the heap recheck overhead common in trigram searches and supports multi-column indexes for wildcard-heavy workloads.

PGXN package 2.4.1 ships the SQL/control version `2.4.0`; the extension's visible `default_version` is therefore still `2.4.0`. The local Pigsty extension name is `biscuit`, while older package metadata may mention `pg_biscuit`.

> [!WARNING]
> Upstream marks Biscuit as actively developed and recommends thorough staging validation before production use. Test representative datasets, query patterns, upgrades, backup/restore, and performance behavior before relying on it for critical workloads.

### Quick Start

```sql
CREATE EXTENSION IF NOT EXISTS biscuit;

CREATE TABLE users (
  id bigserial PRIMARY KEY,
  name text,
  email text,
  bio text
);

CREATE INDEX users_name_biscuit
ON users USING biscuit (name);

SELECT *
FROM users
WHERE name LIKE '%john%';
```

`biscuit` supports ordinary wildcard patterns with `%` and `_`:

```sql
SELECT * FROM users WHERE name LIKE 'john%';
SELECT * FROM users WHERE name LIKE '%smith';
SELECT * FROM users WHERE name LIKE '%oh_';
SELECT * FROM users WHERE name ILIKE '%john%';
SELECT * FROM users WHERE name NOT LIKE '%test%';
```

### Multi-Column Indexes

```sql
CREATE INDEX users_search_biscuit
ON users USING biscuit (name, email, bio);

SELECT *
FROM users
WHERE name ILIKE '%john%'
  AND email LIKE '%example.com'
  AND bio NOT LIKE '%inactive%';
```

Biscuit can combine bitmap matches from multiple indexed columns and may reorder predicates by estimated selectivity.

### Expression Indexes

Version 2.4.0 adds expression-index support:

```sql
CREATE INDEX users_lower_name_biscuit
ON users USING biscuit (lower(name));

SELECT *
FROM users
WHERE lower(name) LIKE '%john%';
```

For `char(n)` / `bpchar` columns, upstream recommends expression indexes that cast to `text`, because native `bpchar` operator classes are not yet available:

```sql
CREATE INDEX legacy_code_biscuit
ON legacy_table USING biscuit ((code::text));
```

### Introspection

```sql
SELECT *
FROM biscuit_operators;

SELECT *
FROM biscuit_version_history;
```

The `biscuit_operators` view lists the operators registered for the Biscuit access method. In 2.4.0, the view was fixed to remain correct if additional operator classes or families are added.

### Operational Notes

Biscuit's design is optimized for:

- prefix, suffix, substring, and mixed wildcard `LIKE` / `ILIKE` patterns
- multi-column predicates where bitmap intersections can reduce candidate sets
- exact pattern matching without trigram false-positive rechecks
- workloads where text-pattern search dominates query latency

It is not a general full-text search engine and does not replace ranking, stemming, tokenization, or phrase search. Use PostgreSQL full text search, trigram indexes, or dedicated search extensions when those semantics are required.

### Caveats

- Upstream requires PostgreSQL 16 or newer and standard build tools. Pigsty local metadata currently packages Biscuit for PostgreSQL 16-18.
- PGXN package version 2.4.1 carries SQL/control `default_version = '2.4.0'`; this is expected for the current source package.
- Biscuit only targets `LIKE` / `ILIKE`-style wildcard matching. Regular expressions are not the supported search surface.
- Non-text columns should be indexed through explicit text expressions when needed.
- Benchmark against `pg_trgm` and your actual data distribution before replacing existing production indexes.
