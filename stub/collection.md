

## Usage

> [collection: key-value collection data types for PL/pgSQL](https://github.com/aws/pgcollection)

The `collection` extension provides two memory-optimized collection data types for use within PL/pgSQL functions.

```sql
CREATE EXTENSION collection;
```

### Data Types

- **`collection`**: Key-value pairs with text keys (max 32,767 chars), stored in creation order
- **`icollection`**: Key-value pairs with 64-bit integer keys, enabling sparse arrays

Both types support type modifiers to specify element types:

```sql
DECLARE
    c1  collection('date');
    ic1 icollection('int4');
```

### Subscript Access

```sql
DO $$
DECLARE t_capital collection;
BEGIN
    t_capital['USA'] := 'Washington, D.C.';
    t_capital['Japan'] := 'Tokyo';
    RAISE NOTICE '%', t_capital['USA'];  -- Washington, D.C.
END $$;
```

### Core Functions

| Function | Description |
|----------|-------------|
| `add(coll, key, value)` | Add element |
| `count(coll)` | Element count |
| `delete(coll, key)` | Remove element |
| `exist(coll, key)` | Check key existence |
| `find(coll, key)` | Retrieve value |
| `first(coll)` | Move iterator to start |
| `last(coll)` | Move iterator to end |
| `next(coll)` | Advance iterator |
| `prev(coll)` | Reverse iterator |
| `key(coll)` | Current key |
| `value(coll)` | Current value |
| `copy(coll)` | Create copy |
| `sort(coll)` | Sort by keys |
| `keys_to_table(coll)` | All keys as set |
| `values_to_table(coll)` | All values as set |
| `to_table(coll)` | Keys and values as table |

### Iterator Example

```sql
DO $$
DECLARE t_capital collection;
BEGIN
    t_capital['USA'] := 'Washington, D.C.';
    t_capital['United Kingdom'] := 'London';
    t_capital['Japan'] := 'Tokyo';

    t_capital := first(t_capital);
    WHILE NOT isnull(t_capital) LOOP
        RAISE NOTICE 'Capital of % is %', key(t_capital), value(t_capital);
        t_capital := next(t_capital);
    END LOOP;
END $$;
```

### Sparse Arrays (icollection)

`icollection` supports non-contiguous integer indices and distinguishes between NULL values and uninitialized keys:

```sql
DECLARE sparse icollection;
BEGIN
    sparse[1] := 'first';
    sparse[1000000] := 'millionth';  -- no memory wasted on gaps
END;
```
