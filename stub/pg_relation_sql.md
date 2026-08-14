## Usage

Sources:

- [pg_relation_sql 0.2.2 on PGXN](https://pgxn.org/dist/pg_relation_sql/0.2.2/)
- [pg_relation_sql 0.2.2 README](https://api.pgxn.org/src/pg_relation_sql/pg_relation_sql-0.2.2/README.md)
- [pg_relation_sql 0.2.2 SQL script](https://api.pgxn.org/src/pg_relation_sql/pg_relation_sql-0.2.2/relation_sql.sql)
- [pg_relation_sql 0.2.2 plan comparison](https://api.pgxn.org/src/pg_relation_sql/pg_relation_sql-0.2.2/EXPLAIN.md)

`pg_relation_sql` 0.2.2 generates pairs of SQL functions from PostgreSQL foreign keys: a lookup follows a reference, while a list function returns rows that point back. The generated `LANGUAGE sql` functions are designed to be inlined by the planner, allowing queries to navigate declared relations without repeating join conditions.

Upstream deliberately ships one standalone `relation_sql.sql` file rather than a control file. There is no `CREATE EXTENSION pg_relation_sql`; execute the packaged script in every database where the functions are needed.

```bash
psql app -f /usr/pgsql-17/share/pg_relation_sql/relation_sql.sql
psql app -f /usr/share/postgresql/17/pg_relation_sql/relation_sql.sql
```

The script creates `relation_sql(text)` in the current schema and finishes by requesting `relation_sql('install')`.

### Generate and Use Relations

```sql
CREATE TABLE profile (
  id bigint PRIMARY KEY,
  name text
);

CREATE TABLE address (
  id bigint PRIMARY KEY,
  profile_id bigint REFERENCES profile(id),
  city text
);

SELECT status, command FROM relation_sql('sync');

SELECT a.city, p.name
FROM address AS a, profile(a) AS p;

SELECT p.name, a.city
FROM profile AS p, address_list(p) AS a;
```

For each foreign key, the lookup function follows the referenced row and the reverse function uses a `_list` suffix unless the foreign key is one-to-one. Composite and cross-schema foreign keys are supported, and several foreign keys to the same target receive role-specific names.

### Generator Modes

- `relation_sql()` returns a status dashboard.
- `relation_sql('show')` reports the computed functions and ready-to-run synchronization commands without changing objects.
- `relation_sql('sync')` creates, replaces, or removes marked relation functions to match current foreign keys.
- `relation_sql('install')` adds a `ddl_command_end` event trigger and synchronizes immediately.
- `relation_sql('uninstall')` removes the event trigger; `relation_sql('drop')` removes generated functions.

### Operational Boundaries

- Creating the event trigger requires superuser privileges. Without them, installation emits a warning and the one-time synchronization still runs with the caller's object privileges.
- Install the generator in a trusted schema with a controlled `search_path`: automatic mode creates a `SECURITY DEFINER` event-trigger helper that preserves the installation-time path.
- Generated functions depend on table row types. Dropping a table whose row type is used by them can require `CASCADE`; inspect dependencies before destructive DDL.
- The generated bodies use `SELECT *`, so column-level `SELECT` grants do not combine cleanly with them. Row-level security continues to apply.
- Put relation functions in `FROM` for plan-sensitive queries. Attribute notation in a select list becomes a `ProjectSet`, and `NOT EXISTS (SELECT FROM relation_function(row))` can remain a correlated probe instead of becoming the equivalent anti-join.
- Queries depend on generated functions just as they depend on views. Run `relation_sql('sync')` in the migration path when not using the event trigger.
- Upstream requires PostgreSQL 11 or later; Pigsty packages cover PostgreSQL 14–18.
