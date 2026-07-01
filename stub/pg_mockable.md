


## Usage

> Sources: [pg_mockable upstream README](https://github.com/bigsmoke/pg_mockable/blob/v1.1.0/README.md), [v1.1.0 tag](https://github.com/bigsmoke/pg_mockable/tree/v1.1.0), [PGXN pg_mockable](https://pgxn.org/dist/pg_mockable/), [local metadata](../db/extension.csv), local source tarball `pg_mockable-1.1.0.tar.gz`.

`pg_mockable` creates mockable wrapper functions for PostgreSQL routines. It is mainly useful in database tests where application code should call a stable wrapper, while tests temporarily replace the wrapper's return value.

```sql
CREATE EXTENSION pg_mockable CASCADE;
```

The extension installs into the fixed `mockable` schema and is not relocatable.

### Mock Built-In Time Functions

`mockable.now()` is pre-created because mocking `now()` also covers the related current-time wrappers exposed by this extension.

```sql
SELECT mockable.now();

SELECT mockable.mock(
  'pg_catalog.now()',
  '2026-06-17 10:00:00+08'::timestamptz
);

SELECT mockable.now();
SELECT mockable.current_timestamp();
SELECT mockable.current_date();

CALL mockable.unmock('pg_catalog.now()');
```

`mockable.mock(regprocedure, anyelement)` stores the mock value and returns it. `mockable.unmock(regprocedure)` clears the mock and restores the wrapper to call the original routine.

### Wrap Application Functions

Use `mockable.wrap_function()` to create a thin wrapper in the `mockable` schema:

```sql
CREATE SCHEMA app;

CREATE FUNCTION app.answer()
RETURNS int
LANGUAGE sql
RETURN 42;

SELECT mockable.wrap_function('app.answer()');

SELECT mockable.answer();                 -- 42
SELECT mockable.mock('app.answer()', 7);   -- 7
SELECT mockable.answer();                 -- 7

CALL mockable.unmock('app.answer()');
SELECT mockable.answer();                 -- 42
```

The first argument is a `regprocedure`, so include argument types when the function is overloaded:

```sql
SELECT mockable.wrap_function('pg_catalog.current_setting(text, boolean)');
```

If automatic wrapper generation is not sufficient, pass the exact `CREATE OR REPLACE FUNCTION` statement as the second argument:

```sql
SELECT mockable.wrap_function(
  'app.answer()',
  $$
  CREATE OR REPLACE FUNCTION mockable.answer()
  RETURNS int
  LANGUAGE sql
  RETURN app.answer();
  $$
);
```

Version 1.1.0 also adds optional debug logging for wrapped/mockable routines through `raise_debug_messages$` on `mockable.wrap_function(...)` and the `mock_memory.raise_debug_messages` column.

### Mock Lifetime

The default mock lifetime is transaction-scoped. For values that must survive dump/restore or later transactions, create the wrapper with a persistent lifetime:

```sql
SELECT mockable.wrap_function(
  'app.answer()',
  mock_duration$ => 'PERSISTENT'
);
```

Persistent mocks should be explicitly cleared when the test fixture no longer needs them:

```sql
CALL mockable.unmock('app.answer()');
```

### Search Path Caveat

Application code must actually call the wrapper, for example `mockable.now()` or `mockable.answer()`, for the mock to apply. Some PL/pgSQL code can be redirected by adjusting `search_path`, but expressions such as table defaults are compiled to function OIDs; adding `mockable` to `search_path` later does not rewrite those references. Prefer explicit `mockable.*` calls in code that is meant to be testable.

### Caveats

- Version 1.1.0 supports PostgreSQL 14-18. It is a SQL extension and does not need `shared_preload_libraries`.
- `pg_mockable` owns the `mockable` schema; installing it in another schema is not supported by the control file.
- Wrapper privileges are derived from the wrapped routine. The tests verify that wrapping a private function does not grant execute privilege to roles that could not call the original function.
