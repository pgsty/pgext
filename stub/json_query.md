## Usage

Sources:

- [Project README](https://github.com/dgillis/pg_json_query/blob/6aa96058553cfd01a6c9a508d4d925fcbf38a064/README.md)
- [Version 0.9 installation SQL](https://github.com/dgillis/pg_json_query/blob/6aa96058553cfd01a6c9a508d4d925fcbf38a064/build/json_query--0.9.sql)
- [Upstream example](https://github.com/dgillis/pg_json_query/blob/6aa96058553cfd01a6c9a508d4d925fcbf38a064/example.sql)

`json_query` 0.9 is a pure-SQL query-building extension that turns validated JSONB filter objects into predicates over a registered table or composite row type. It is useful when a SQL function needs an application-supplied set of optional filters without maintaining many nearly identical function variants.

### Register a row type and filter it

```sql
CREATE EXTENSION json_query;

CREATE TABLE customer (id integer, status text);
SELECT jq_register_row_type('customer');

SELECT *
FROM customer AS c
WHERE jq_filter(c, '{"status": "active", "id__ge": 100}'::jsonb);
```

`jq_filter(anyelement,jsonb)` accepts Django-style keys such as `id__ge`; `jq_filter_raw(anyelement,jsonb)` accepts an explicit array of field, operator, and value objects. Operators include equality and ordering, membership, pattern matching, and JSONB containment or existence tests. Nested JSONB fields can use `->` and `->>` path syntax.

### Registration lifecycle

`jq_register_row_type(text)` generates type-specific helper functions. Run `jq_unregister_row_type(text)` and register again after changing the table definition. Custom column types may also require `jq_register_column_type(text)`. The extension validates field names and operator names, but applications should still authorize which filters callers may use and cap query cost. The README's published minimum is PostgreSQL 9.4; the reviewed source has no current release compatibility matrix.
