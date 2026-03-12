

## Usage

> [plpgsql: PL/pgSQL procedural language](https://www.postgresql.org/docs/current/plpgsql.html)

PL/pgSQL is PostgreSQL's default procedural language. It extends SQL with control structures, variables, cursors, and exception handling.

```sql
CREATE EXTENSION plpgsql;  -- installed by default

-- Basic function with variables and control flow
CREATE FUNCTION calculate_discount(price numeric, quantity integer) RETURNS numeric
LANGUAGE plpgsql AS $$
DECLARE
  discount numeric := 0;
BEGIN
  IF quantity >= 100 THEN
    discount := 0.20;
  ELSIF quantity >= 50 THEN
    discount := 0.10;
  ELSIF quantity >= 10 THEN
    discount := 0.05;
  END IF;
  RETURN price * quantity * (1 - discount);
END;
$$;

-- Loop and set-returning function
CREATE FUNCTION fibonacci(n integer) RETURNS SETOF integer
LANGUAGE plpgsql AS $$
DECLARE
  a integer := 0;
  b integer := 1;
  tmp integer;
BEGIN
  FOR i IN 1..n LOOP
    RETURN NEXT a;
    tmp := a + b;
    a := b;
    b := tmp;
  END LOOP;
END;
$$;

SELECT * FROM fibonacci(10);

-- Exception handling
CREATE FUNCTION safe_divide(a numeric, b numeric) RETURNS numeric
LANGUAGE plpgsql AS $$
BEGIN
  RETURN a / b;
EXCEPTION
  WHEN division_by_zero THEN
    RAISE NOTICE 'Division by zero, returning NULL';
    RETURN NULL;
END;
$$;

-- Trigger function
CREATE FUNCTION update_modified_column() RETURNS trigger
LANGUAGE plpgsql AS $$
BEGIN
  NEW.modified_at = now();
  RETURN NEW;
END;
$$;

CREATE TRIGGER set_modified
  BEFORE UPDATE ON my_table
  FOR EACH ROW EXECUTE FUNCTION update_modified_column();

-- Procedure with transaction control (PG 11+)
CREATE PROCEDURE batch_archive(batch_size integer)
LANGUAGE plpgsql AS $$
DECLARE
  rows_moved integer;
BEGIN
  LOOP
    WITH moved AS (
      DELETE FROM orders WHERE status = 'completed'
      RETURNING *
    )
    INSERT INTO orders_archive SELECT * FROM moved;

    GET DIAGNOSTICS rows_moved = ROW_COUNT;
    COMMIT;
    EXIT WHEN rows_moved < batch_size;
  END LOOP;
END;
$$;
```
