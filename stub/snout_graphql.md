## Usage

Sources:

- [README 0.1.0](https://github.com/snoutdata/snout-graphql/blob/4f549d1738072780d9cb0df7fe9d848e6666bbab/README.md)
- [Control](https://github.com/snoutdata/snout-graphql/blob/4f549d1738072780d9cb0df7fe9d848e6666bbab/snout_graphql.control)
- [SQL 0.1.0](https://github.com/snoutdata/snout-graphql/blob/4f549d1738072780d9cb0df7fe9d848e6666bbab/sql/snout_graphql--0.1.0.sql)
- [Cargo](https://github.com/snoutdata/snout-graphql/blob/4f549d1738072780d9cb0df7fe9d848e6666bbab/Cargo.toml)

`snout_graphql` 0.1.0 resolves GraphQL requests inside PostgreSQL from visible tables, views and functions. This Rust extension targets PostgreSQL 17. Installation requires a superuser and creates objects in the fixed `graphql` schema plus two event triggers.

### Core Workflow

```sql
CREATE EXTENSION snout_graphql;
COMMENT ON SCHEMA public IS '@graphql({"inflect_names": true})';
SELECT graphql.resolve($$ { __typename } $$);
```

`graphql.resolve(query, variables, "operationName", extensions)` returns a JSONB response. Values are passed separately from the document. Requests execute with the caller's table privileges and row-level security policies; callers also need access to the extension's schema and objects. An error produces an errors response with null data and rolls back writes made by that request. An HTTP endpoint requires a separate server to invoke the function as the request's database role.

### Objects and Configuration

- `graphql._internal_resolve`: internal entry point that raises errors directly.
- `graphql.comment_directive`: reads JSON directives embedded in object comments.
- `graphql.get_schema_version`, `graphql.increment_schema_version`, `graphql.seq_schema_version`: support reflected-schema cache invalidation.
- `graphql_watch_ddl`, `graphql_watch_drop`: event triggers tracking relevant DDL and dropped objects.
- `graphql.exception`: reports a mutation exceeding its permitted row count.
- Comment directives include `inflect_names`, `max_rows`, `introspection`, `totalCount`, `aggregate`, `primary_key_columns`, and `foreign_keys`. Introspection is off by default and collection pages default to 30 rows. There are no extension GUCs or background workers.

### Boundaries

The entry-point functions and event triggers overlap those used by `pg_graphql`; do not install both into the same schema or assume an automatic migration path. The upstream README's linked divergences document is absent at this revision, so equivalence with another implementation is not established here.

Each connection caches reflected schemas, plans and introspection responses. Object comments affect the exposed API, and table grants and row-level security remain the application's responsibility. GraphQL text is parsed inside a PostgreSQL backend; this source preview is not a production certification.
