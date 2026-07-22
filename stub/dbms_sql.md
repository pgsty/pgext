## Usage

Sources:

- [Official README](https://github.com/okbob/orafce_sql/blob/76f311c75074131f4d076bf28981e5f97bc8bf6a/README.md)
- [Official extension SQL](https://github.com/okbob/orafce_sql/blob/76f311c75074131f4d076bf28981e5f97bc8bf6a/dbms_sql--1.0.sql)
- [Official extension control file](https://github.com/okbob/orafce_sql/blob/76f311c75074131f4d076bf28981e5f97bc8bf6a/dbms_sql.control)

`dbms_sql` implements a PostgreSQL-oriented subset of Oracle's `DBMS_SQL` cursor API. It helps migrate PL/SQL that parses dynamic statements, binds scalar or array values, executes them, describes result columns, and fetches rows; it does not promise full Oracle compatibility.

### Core Workflow

Open a cursor, parse a statement with named binds, bind values, execute, and close the cursor:

```sql
CREATE EXTENSION dbms_sql;

DO $$
DECLARE
  c integer;
  inserted bigint;
BEGIN
  c := dbms_sql.open_cursor();
  CALL dbms_sql.parse(c, 'insert into target_table(id, label) values(:id, :label)');
  CALL dbms_sql.bind_variable(c, 'id', 42);
  CALL dbms_sql.bind_variable(c, 'label', 'example'::text);
  inserted := dbms_sql.execute(c);
  CALL dbms_sql.close_cursor(c);
  RAISE NOTICE 'affected rows: %', inserted;
END
$$;
```

For query results, call `define_column` or `define_array`, execute, loop over `fetch_rows`, retrieve values with `column_value`, and always close the cursor in normal and exception paths.

### Main Interface

- `open_cursor`, `is_open`, and `close_cursor` manage extension cursor handles.
- `parse` prepares dynamic SQL; `bind_variable` and `bind_array` supply named parameters.
- `execute`, `execute_and_fetch`, `fetch_rows`, and `last_row_count` drive execution and fetching.
- `define_column`, `define_array`, and `column_value` describe destination buffers and copy fetched values.
- `describe_columns` and `describe_columns_f` return `dbms_sql.desc_rec` metadata. The `col_type` numbers are PostgreSQL type OIDs, not Oracle type codes.

### Migration Boundaries

- Audit every migrated call against the implemented signatures and PostgreSQL transaction/error semantics. Unsupported Oracle overloads, type conversions, privilege behavior, and cursor lifetime must be rewritten explicitly.
- The install SQL creates an unqualified `varchar2` domain as a temporary compatibility object. Upstream instructs users who combine this extension with Orafce to remove that statement from the install script, because Orafce supplies its own `varchar2`.
- Cursor handles are backend-local execution state. They do not cross sessions, pool checkouts, failover, or transaction-retry boundaries; close them deterministically to avoid leaking state in long-lived sessions.
- Dynamic SQL still requires normal injection discipline. Bind data values, validate any identifiers or SQL fragments that cannot be bound, and run with the least-privileged caller.
- The repository is archived. Pin the reviewed commit and test its direct use of PostgreSQL internals against the exact server major version before migration or upgrade.
