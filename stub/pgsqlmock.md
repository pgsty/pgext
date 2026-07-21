## Usage

Sources:

- [pgSQLMock 1.0.1 documentation](https://github.com/v-maliutin/pgSQLMock/blob/Release_v1.0.1/doc/pgsqlmock.md)
- [pgSQLMock README](https://github.com/v-maliutin/pgSQLMock/blob/Release_v1.0.1/README.md)
- [pgSQLMock control file](https://github.com/v-maliutin/pgSQLMock/blob/Release_v1.0.1/pgsqlmock.control)
- [pgSQLMock 1.0.1 release](https://github.com/v-maliutin/pgSQLMock/releases/tag/Release_v1.0.1)

`pgsqlmock` extends pgTAP with table fakes, function and view mocks, call-count assertions, and debugging helpers. Its helpers alter or replace real database objects, so upstream requires using them inside pgTAP's transaction-based test context, where the changes are rolled back after the test.

```sql
CREATE EXTENSION pgtap;
CREATE EXTENSION pgsqlmock;
```

### Fake Tables

`fake_table(text[], ...)` can isolate a test from foreign keys, primary keys, `NOT NULL` constraints, partitions, or pre-existing rows. Pass schema-qualified table names as a `text[]`:

```sql
SELECT plan(2);

SELECT fake_table(
  _table_ident       => ARRAY['app.accounts', 'app.transactions'],
  _make_table_empty  => true,
  _leave_primary_key => false,
  _drop_not_null     => true
);

INSERT INTO app.transactions(account_id, amount)
VALUES (999, 42.00);

SELECT is(
  (SELECT sum(amount) FROM app.transactions WHERE account_id = 999),
  42.00::numeric,
  'transaction logic is isolated from account fixtures'
);

SELECT * FROM finish();
```

Important options include `make_table_empty`, `leave_primary_key`, `drop_not_null`, `drop_collation`, and `drop_partitions`. Keeping a primary key while dropping the participating columns' `NOT NULL` constraints is contradictory; remove or recreate the key explicitly for that test shape.

### Mock Functions

`mock_func(schema, name, signature, ...)` temporarily replaces a routine while preserving its identity. Supply either a scalar value or SQL/prepared-statement text for a set result:

```sql
CREATE OR REPLACE FUNCTION app.current_business_time()
RETURNS time LANGUAGE sql AS $$ SELECT current_time $$;

SELECT mock_func(
  'app',
  'current_business_time',
  '()',
  _return_scalar_value => '13:00'::time
);

SELECT is(app.current_business_time(), '13:00'::time, 'clock is deterministic');
```

For set-returning routines, pass `_return_set_value` as a SQL query or the name of a prepared statement. Use `get_routine_signature()` when overloaded or defaulted arguments make the stored signature unclear.

### Mock Views

`mock_view(schema, view_name, return_set_sql)` replaces a view with controlled rows:

```sql
SELECT mock_view(
  'app',
  'active_accounts',
  $$SELECT * FROM (VALUES (1, 'test')) AS v(id, name)$$
);

SELECT results_eq(
  'SELECT id, name FROM app.active_accounts',
  $$VALUES (1, 'test')$$,
  'view consumer sees only the fixture'
);
```

### Call Counts and Diagnostics

Set `track_functions = 'all'` before using `call_count()` to assert how often a routine was invoked:

```sql
SET LOCAL track_functions = 'all';

SELECT call_count(
  1,
  'app',
  'current_business_time',
  '()'
);
```

`print_table_as_json()` and `print_query_as_json()` emit reproducible SQL/JSON-style snapshots through `NOTICE`, which is useful when pgTAP's rollback would otherwise hide the state created during a failed test.

### Caveats

- Run mocks and fakes only inside isolated test transactions; they issue real `ALTER`, `DROP`, and replacement DDL.
- pgSQLMock depends on PL/pgSQL and pgTAP. Load pgTAP before running its assertions.
- `call_count()` depends on PostgreSQL function statistics and therefore requires `track_functions = 'all'`.
- Release 1.0.1 fixes `fake_table()` dropping `NOT NULL` constraints on tables without a primary key.
