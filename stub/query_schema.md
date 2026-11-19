## Usage

Sources:

- [Official source](https://github.com/olhamide/pg-query-schema/blob/151f38957e389c0f0d4e5d377249b484030fae4c/README.md)
- [Official source](https://github.com/olhamide/pg-query-schema/blob/151f38957e389c0f0d4e5d377249b484030fae4c/query_schema.control)
- [Official source](https://github.com/olhamide/pg-query-schema/blob/151f38957e389c0f0d4e5d377249b484030fae4c/sql/query_schema--1.0.sql)
- [Official source](https://github.com/olhamide/pg-query-schema/blob/151f38957e389c0f0d4e5d377249b484030fae4c/src/query-schema.c)
- [Official source](https://github.com/olhamide/pg-query-schema/blob/151f38957e389c0f0d4e5d377249b484030fae4c/Dockerfile)

`query_schema` `1.0` uses PostgreSQL's parser and analyzer to describe query results as Effect schema source text. Upstream limits this extension to local development; the supplied release artifacts target PostgreSQL 18 on Alpine for amd64 and arm64.

### Workflow

Install in a development database as a superuser. The module does not require preload or a server restart:

```sql
CREATE EXTENSION query_schema;
CREATE TEMP TABLE example_accounts (id integer NOT NULL, label text);
SELECT query_schema('SELECT id, label FROM example_accounts');
SELECT * FROM query_shape('SELECT id, label FROM example_accounts WHERE id = $1');
```

`query_schema(text)` returns a text expression based on `Schema.Struct`. `query_shape(text)` returns `path`, `pg_type`, `nullable`, `schema` and `note` columns, including inferred parameter types and nested JSON/array paths. Parameter rows have a null `nullable` value.

### Analysis boundaries

Pass exactly one statement. The implementation accepts `SELECT`, `INSERT`, `UPDATE`, `DELETE` and `MERGE`; data-modifying statements describe their `RETURNING` fields. Utility commands and multiple statements are rejected. It analyzes and, for selects, rewrites the query rather than executing its result plan.

The analyzer uses the current database objects and name resolution. Outer joins, conditions, aggregates, CTEs, enum values and JSON constructors influence the output. Unknown nullability is represented conservatively as nullable. Top-level type mapping follows `@effect/sql-pg`, while nested JSON follows PostgreSQL's JSON conversion, so the two representations can differ.

This is a development aid, not a proof that arbitrary input is safe or that an inferred schema fully validates runtime results. Review generated schemas and notes against the real client decoder. SQL functions retain ordinary PostgreSQL execution grants; restrict access before sharing a development database. The inspected source does not declare a license, so reuse rights remain unverified.
