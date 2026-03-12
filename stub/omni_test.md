


## Usage

> [omni_test: Testing framework](https://docs.omnigres.org/omni_test/guide/)

The `omni_test` extension enables test execution within PostgreSQL databases. It is a templated extension.

### Setup

```sql
CREATE DATABASE myapp_test;
UPDATE pg_database SET datistemplate = true WHERE datname = 'myapp_test';
```

### Writing Tests

**Function tests:**

```sql
CREATE FUNCTION my_test() RETURNS omni_test.test
    LANGUAGE plpgsql AS $$
BEGIN
    -- test assertions here
END;
$$;
```

**Procedure tests** (for non-atomic tests requiring COMMIT/ROLLBACK):

```sql
CREATE PROCEDURE my_test(INOUT omni_test.test)
    LANGUAGE plpgsql AS $$
BEGIN
    -- test logic
END;
$$;
```

### Test Settings

```sql
CREATE PROCEDURE tx_test(INOUT test omni_test.test)
    SET omni_test.transaction_isolation = serializable
    LANGUAGE plpgsql AS $$ ... $$;
```

### Running Tests

```sql
SELECT * FROM omni_test.run_tests('myapp_test');
```

Returns: `name`, `description`, `start_time`, `end_time`, `error_message`.

### Filtering Tests

```sql
-- Exclude tests marked @slow:
SELECT * FROM omni_test.run_tests('myapp_test', filter => '^(?!.*@slow).*');
```

Each test runs in a fresh copy of the template database.
