## Usage

Sources:

- [cat_tools 0.3.0 README](https://github.com/Postgres-Extensions/cat_tools/blob/0.3.0/README.asc)
- [cat_tools 0.3.0 history](https://github.com/Postgres-Extensions/cat_tools/blob/0.3.0/HISTORY.asc)
- [cat_tools 0.3.0 control file](https://github.com/Postgres-Extensions/cat_tools/blob/0.3.0/cat_tools.control)
- [cat_tools 0.3.0 install SQL](https://github.com/Postgres-Extensions/cat_tools/blob/0.3.0/sql/cat_tools--0.3.0.sql.in)

`cat_tools` provides typed views, enums, and helper functions for PostgreSQL catalog introspection. It is designed for database code that needs a more stable and readable interface than repeatedly decoding raw `pg_catalog` fields; the views still track PostgreSQL's catalogs and must be reviewed across major-version upgrades.

### Install and Grant Access

```sql
CREATE EXTENSION cat_tools;
GRANT cat_tools__usage TO app_introspection;
```

The extension installs in the fixed `cat_tools` schema, requires `plpgsql`, and is not relocatable. Grant the `cat_tools__usage` role rather than exposing internal `_cat_tools` helpers directly.

### Inspect Relations and Columns

```sql
SELECT cat_tools.relation__kind(c.relkind::text)
FROM pg_catalog.pg_class AS c
WHERE c.oid = 'public.orders'::regclass;

SELECT cat_tools.relation__column_names('public.orders'::regclass);
SELECT cat_tools.pg_attribute__get('public.orders'::regclass, 'id');
```

Useful relation helpers include `pg_class(regclass)`, `relation__is_catalog`, `relation__is_temp`, `relation__kind`, and `relation__relkind`. Typed mapping functions make the one-character catalog codes explicit.

### Inspect Routines

Version 0.3 adds functions and types that cover both functions and procedures:

```sql
SELECT cat_tools.routine__arg_types(
  'public.calculate_total(integer, numeric)'::regprocedure
);

SELECT cat_tools.routine__parse_arg_names(
  'IN account_id integer, INOUT total numeric'
);
```

The routine surface includes `routine__parse_arg_types`, `routine__parse_arg_names`, `routine__arg_types`, `routine__arg_names`, their text variants, and mappings for routine kind, argument mode, volatility, and parallel safety. `function__arg_types` and `function__arg_types_text` are deprecated; use the routine parsers.

### Version 0.3.0 and Caveats

- Version 0.3.0 supports PostgreSQL 12-18+ upstream; current Pigsty packages cover PostgreSQL 14-18.
- The release corrects the `c`, `f`, and `m` mappings for composite types, foreign tables, and materialized views. Re-test any code that worked around the old mapping.
- Internal `_cat_tools` helpers now revoke `EXECUTE` from `PUBLIC`; callers should inherit `cat_tools__usage` and use the supported surface.
- The 0.2.3-to-0.3.0 update adds enum values and therefore cannot run on PostgreSQL 11 or earlier. Upgrade the database major version and extension in the order documented upstream.
- PostgreSQL does not promise catalog compatibility across major releases. Pin tests to every supported PostgreSQL major even when using these wrappers.
