## Usage

Sources:

- [Official README](https://github.com/olhamide/pg-query-schema/blob/v1.1.0/README.md)
- [Extension control](https://github.com/olhamide/pg-query-schema/blob/v1.1.0/query_schema.control)
- [Extension SQL](https://github.com/olhamide/pg-query-schema/blob/v1.1.0/sql/query_schema--1.1.sql)
- [Upgrade SQL](https://github.com/olhamide/pg-query-schema/blob/v1.1.0/sql/query_schema--1.0--1.1.sql)
- [Implementation](https://github.com/olhamide/pg-query-schema/blob/v1.1.0/src/query-schema.c)
- [Supported build target](https://github.com/olhamide/pg-query-schema/blob/v1.1.0/Dockerfile)

`query_schema` SQL version `1.1`, distributed in upstream release `1.1.0`, uses PostgreSQL's parser and analyzer to describe query results as Effect schema source text. Upstream limits this extension to local development; the supplied release artifacts target PostgreSQL 18 on Alpine for amd64 and arm64.

### Workflow

Install in a development database as a superuser. The module does not require preload or a server restart:

```sql
CREATE EXTENSION query_schema;
CREATE TEMP TABLE example_accounts (id integer NOT NULL, label text);
SELECT query_schema('SELECT id, label FROM example_accounts');
SELECT * FROM query_shape('SELECT id, label FROM example_accounts WHERE id = $1');
```

`query_schema(text)` returns a text expression based on `Schema.Struct`. `query_shape(text)` returns `path`, `pg_type`, `nullable`, `schema` and `note` columns, including inferred parameter types and nested JSON/array paths. Parameter rows have a null `nullable` value.

After replacing the module files, update an existing installation with `ALTER EXTENSION query_schema UPDATE TO '1.1';`. The upgrade SQL changes the recorded extension version without changing function definitions.

### JSON and client decoding

Version `1.1` handles nullability in `LATERAL` joins and set operations, including `UNION`, `INTERSECT` and `EXCEPT`. It also distinguishes top-level binary client values from values inside JSON.

Inside JSON, `int8` values are JSON numbers. The generated schema uses `Schema.Int` and `SchemaGetter` to decode them to a bigint; values outside JavaScript's safe integer range are rejected after JSON parsing. Cast an `int8` to `::text` before JSON construction to preserve its full range; that shape uses `Schema.BigIntFromString`. Generated schemas may require both imports:

```typescript
import { Schema, SchemaGetter } from "effect"
```

`@effect/sql-pg` cannot decode every PostgreSQL type in binary format. For unsupported columns, such as `money`, ranges, composite types or multidimensional arrays, `query_schema` emits a warning and `query_shape` adds a note. Select these values as `::text` or wrap them in `to_json`, then verify the resulting client schema.

### Analysis boundaries

Pass exactly one statement. The implementation accepts `SELECT`, `INSERT`, `UPDATE`, `DELETE` and `MERGE`; data-modifying statements describe their `RETURNING` fields. Utility commands and multiple statements are rejected. It analyzes and, for selects, rewrites the query rather than executing its result plan.

The analyzer uses the current database objects and name resolution. Outer joins, conditions, aggregates, CTEs, enum values and JSON constructors influence the output. Unknown nullability is represented conservatively as nullable. Top-level type mapping follows `@effect/sql-pg`, while nested JSON follows PostgreSQL's JSON conversion, so the two representations can differ.

This is a development aid, not a proof that arbitrary input is safe or that an inferred schema fully validates runtime results. Review generated schemas and notes against the real client decoder. SQL functions retain ordinary PostgreSQL execution grants; restrict access before sharing a development database. The inspected source does not declare a license, so reuse rights remain unverified.
