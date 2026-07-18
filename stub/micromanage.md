## Usage

Sources:

- [Pinned upstream README](https://github.com/lithp/pg_micromanage/blob/1ed36659995a82a6346305f276762c8eed65d32b/README.md)
- [Version 0.0.1 installation SQL](https://github.com/lithp/pg_micromanage/blob/1ed36659995a82a6346305f276762c8eed65d32b/micromanage--0.0.1.sql)
- [Pinned planner implementation](https://github.com/lithp/pg_micromanage/blob/1ed36659995a82a6346305f276762c8eed65d32b/micromanage.c)

micromanage is an experimental PostgreSQL 9.6 planner-hook demonstration. It accepts a base64-encoded protobuf plan, intercepts a narrowly shaped call, and constructs a PostgreSQL plan for sequence scans, simple expressions and filters, nested-loop joins, and sorting.

### Example

The extension build requires protobuf-c and protoc. This upstream example runs a pre-encoded plan that scans the one-column table:

```sql
CREATE EXTENSION micromanage;

CREATE TABLE a (a integer);
INSERT INTO a VALUES (1);

SELECT * FROM run_select('Cg0KBwoFCAESAWEaAggBEgMKAWE=');
```

The result is the value 1. The installed API also includes dump_query for inspecting the plan and encode_protobuf for converting the text form described by the repository's queries.proto file.

### Experimental limitations

The hook recognizes only a SELECT whose outer shape matches the constant run_select call expected by the implementation. It resolves relation names in the public schema and rejects installation alongside another planner hook instead of chaining it.

Source comments explicitly note incomplete cost population, uncertain transaction handling, memory leaks, and likely missing security checks. The code targets old planner internals and upstream documents only PostgreSQL 9.6. The repository is not archived but has not changed since 2018. Use micromanage only as instructional or research code in an isolated database; do not accept untrusted protobuf plans or use it for production query execution.
