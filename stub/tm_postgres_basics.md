## Usage

Sources:

- [Upstream README](https://github.com/trimark-jp/tm-postgres-basics/blob/8541898dfe9ace140b0227504cea8b10d11be809/README.md)
- [Version 1.0.0 install SQL](https://github.com/trimark-jp/tm-postgres-basics/blob/8541898dfe9ace140b0227504cea8b10d11be809/tm_postgres_basics--1.0.0.sql)
- [Extension control file](https://github.com/trimark-jp/tm-postgres-basics/blob/8541898dfe9ace140b0227504cea8b10d11be809/tm_postgres_basics.control)
- [MIT license](https://github.com/trimark-jp/tm-postgres-basics/blob/8541898dfe9ace140b0227504cea8b10d11be809/LICENSE)

tm_postgres_basics 1.0.0 is a small 2016 pure-SQL collection for resolving a relation name to an OID, recursively checking pg_inherits ancestry, and listing inherited tables.

### Inheritance check

```sql
CREATE EXTENSION tm_postgres_basics;
CREATE TEMP TABLE tm_parent();
CREATE TEMP TABLE tm_child() INHERITS (tm_parent);
SELECT tm_is_inherit_from(
  'tm_child'::regclass,
  'tm_parent'::regclass
);
```

The direct OID-based ancestry function returns true for this session-local example.

### Caveats

- tm_name_to_oid searches pg_class by relname alone. It ignores schema and relation kind, and can return an arbitrary match when names are repeated. Prefer a schema-qualified regclass or to_regclass lookup.
- tm_find_tables_inherit_from accepts a schema for enumeration but resolves both parent and children through the schema-blind helper, so same-named objects can produce incorrect results.
- The listing function scans every table in the requested schema and recursively queries pg_inherits for each candidate. It can be unnecessarily expensive on large catalogs.
- Functions and catalog references are not schema-qualified and use the caller's search path. Install in a controlled schema and schema-qualify calls.
- All functions retain default volatile and parallel-unsafe markings even though they only read catalogs, limiting planner optimization.
- Upstream has no tests, upgrade scripts, or current compatibility statement and was last changed in 2016. Modern PostgreSQL catalog functions are generally safer and faster.
