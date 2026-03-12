


## Usage

> [pgtap: Unit testing for PostgreSQL](https://github.com/theory/pgtap)

`pgtap` is a unit testing framework for PostgreSQL that produces TAP (Test Anything Protocol) output, providing hundreds of assertion functions for testing database objects and query results.

```sql
CREATE EXTENSION pgtap;
```

### Test Structure

```sql
BEGIN;
SELECT plan(3);  -- declare how many tests to run

SELECT ok(1 = 1, 'one equals one');
SELECT is(1 + 1, 2, 'addition works');
SELECT isnt(1, 2, 'one is not two');

SELECT * FROM finish();
ROLLBACK;
```

Use `no_plan()` when the test count is not known in advance:

```sql
BEGIN;
SELECT * FROM no_plan();
-- ... tests ...
SELECT * FROM finish();
ROLLBACK;
```

### Basic Assertions

```sql
SELECT ok(expression, description);           -- boolean test
SELECT is(got, expected, description);         -- equality test
SELECT isnt(got, unexpected, description);     -- inequality test
SELECT matches(value, regex, description);     -- regex match
```

### Schema Testing

```sql
SELECT has_table('users');
SELECT has_table('myschema', 'users', 'users table exists');
SELECT has_column('users', 'email');
SELECT col_type_is('users', 'email', 'text');
SELECT col_not_null('users', 'id');
SELECT col_has_default('users', 'created_at');
SELECT has_function('calculate_total');
SELECT has_function('calculate_total', ARRAY['integer', 'numeric']);
SELECT has_index('users', 'users_email_idx');
SELECT has_pk('users');
SELECT has_fk('orders');
```

### Error Testing

```sql
SELECT lives_ok('INSERT INTO t(id) VALUES (1)', 'insert succeeds');
SELECT throws_ok(
  'SELECT 1/0',
  '22012',          -- SQLSTATE for division by zero
  'division by zero'
);
```

### Query Result Testing

```sql
-- Compare ordered result sets
SELECT results_eq(
  'SELECT * FROM active_users()',
  'SELECT * FROM users WHERE active',
  'active_users returns correct rows'
);

-- Compare unordered result sets
SELECT set_eq(
  'SELECT * FROM active_ids()',
  ARRAY[2, 3, 4, 5]
);

-- Check query returns no rows
SELECT is_empty('SELECT * FROM users WHERE id = -1');

-- Compare bag (multiset) results
SELECT bag_eq(
  'SELECT color FROM items',
  $$VALUES ('red'), ('blue'), ('red')$$
);
```

### Running Tests with pg_prove

```bash
pg_prove -d mydb tests/*.sql
pg_prove -d mydb --ext .sql --recurse tests/
```

### xUnit Style

```sql
CREATE FUNCTION test_my_feature() RETURNS SETOF text AS $$
  RETURN NEXT ok(1 = 1, 'basic check');
  RETURN NEXT is(my_func(1), 42, 'function works');
$$ LANGUAGE plpgsql;

SELECT * FROM runtests('test_my_feature');
```
