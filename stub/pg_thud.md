## Usage

Sources:

- [Official repository](https://github.com/BlueTreble/pg_thud)
- [Official control file](https://github.com/BlueTreble/pg_thud/blob/4d2aba37d20c942f9c05412a6f045f328106cfb3/pg_thud.control)
- [Official PGXN distribution](https://pgxn.org/dist/pg_thud/)

`pg_thud` is a test-data factory for PostgreSQL. It registers named SQL queries for a table's row type and materializes their results on first use, letting tests request repeatable fixture sets. The project is inactive and its prose documentation is only a scaffold, so use the SQL-defined API in isolated test databases.

### Core Workflow

Install the declared `pgtap` dependency, register one or more named data sets, then fetch a set using a typed `NULL` to identify the target row type:

```sql
CREATE EXTENSION pgtap;
CREATE EXTENSION pg_thud;

CREATE TABLE widget (id integer, name text);

SELECT tf.register(
  'widget',
  ARRAY[
    ROW(
      'small',
      'SELECT * FROM (VALUES (1, ''one''), (2, ''two'')) AS v(id, name)'
    )::tf.test_set
  ]
);

SELECT * FROM tf.get(NULL::widget, 'small');
```

The first `tf.get` call creates cached data in the extension's `_test_data` schema; later calls reuse it.

### API and Objects

- `tf.test_set`: composite value containing `set_name` and `insert_sql`.
- `tf.register(table_name text, test_sets tf.test_set[])`: registers factories for a table.
- `tf.get(anyelement, set_name text) RETURNS SETOF anyelement`: returns the named fixture set as the caller's composite type.
- `_tf.test_factory`: internal registry; `_test_data` holds materialized fixture tables.

### Caveats

Both public functions are `SECURITY DEFINER`, and registration stores SQL text that is later executed to build fixture data. Restrict `EXECUTE` privileges and never accept untrusted factory SQL. Cached data can become stale when a table definition or factory query changes; the published API does not document a refresh command, so recreate disposable test state when definitions change.

`pg_thud` creates data for tests but does not itself assert results. Its dependency and source target old PostgreSQL generations, and the upstream project is abandoned. Validate installation against the exact server version and keep it out of production databases.
