## Usage

Sources:

- [pg_describe 1.0.0 README](https://api.pgxn.org/src/pg_describe/pg_describe-1.0.0/README.md)
- [pg_describe documentation](https://sajonaro.github.io/pg_describe/)
- [pg_describe 1.0.0 control file](https://api.pgxn.org/src/pg_describe/pg_describe-1.0.0/pg_describe.control)
- [pg_describe 1.0.0 SQL](https://api.pgxn.org/src/pg_describe/pg_describe-1.0.0/sql/pg_describe--1.0.0.sql)

`pg_describe` reports the parameters and result columns of a SQL statement without executing it. It uses PostgreSQL parsing and analysis to infer parameter types, wire-visible result types, source-column provenance, and outer-join-aware nullability. Use it for code generation, migration checks, and query-contract tooling.

### Describe a Query

```sql
CREATE EXTENSION pg_describe;

SELECT *
FROM pg_describe(
  'SELECT id, email FROM users WHERE id = $1'
);
```

Rows with `kind = 'param'` describe `$1`, `$2`, and later parameters. Rows with `kind = 'column'` describe result-column order, name, type OID/name, source table/column, base `NOT NULL` status, and whether the final expression is known non-null.

### Check Join Nullability

```sql
SELECT *
FROM pg_describe($query$
  SELECT o.id, c.email
  FROM orders AS o
  LEFT JOIN customers AS c ON c.id = o.customer_id
  WHERE o.placed_at >= $1
$query$);
```

Even when `customers.email` is declared `NOT NULL`, `result_not_null` is false because a left join can null-extend the row. This distinction is useful when generating nullable client types.

### Execution and Security Boundary

- The statement is parsed and analyzed but not executed. Describing a `DELETE`, volatile function call, or expensive query does not run the statement.
- Normal name resolution and privilege checks still apply. Callers cannot use `pg_describe` to inspect objects they could not reference themselves.
- Parameter types must be inferable from context; ambiguous `$n` parameters still produce PostgreSQL analysis errors.
- The result describes PostgreSQL's analyzed output, not dynamic SQL assembled later by an application.

### Requirements and Caveats

- Upstream 1.0.0 requires PostgreSQL 17; PostgreSQL 16 is described as possibly working but untested. Pigsty packages target PostgreSQL 17 and 18.
- The extension is relocatable and does not require preloading or a restart.
- The companion `pg-describe-gen` TypeScript tool is a separate npm package. The PostgreSQL extension works without it.
- This is a young API. Pin the extension/tool versions in CI and review generated changes alongside schema migrations.
