

## Usage

> [pg_rational: precise fractional arithmetic in 64 bits](https://github.com/begriffs/pg_rational)

The `pg_rational` extension provides exact fractional arithmetic stored in 64 bits (same size as float), with support for Stern-Brocot trees for finding intermediate fractions.

```sql
CREATE EXTENSION pg_rational;
```

### Data Types

- **`rational`**: Fractional type (numerator/denominator)
- **`ratt`**: Helper type for tuple coercion

### Basic Arithmetic

```sql
SELECT '1/3'::rational + '2/7'::rational;  -- 13/21
SELECT '3/4'::rational * '2/3'::rational;  -- 1/2
```

### Functions

```sql
-- Simplify fractions
SELECT rational_simplify('36/12');  -- 3/1

-- Find intermediate fraction (Stern-Brocot tree)
SELECT rational_intermediate('1/2', '2/3');  -- 3/5
```

### Type Conversions

```sql
SELECT 0.263157894737::float::rational;  -- 5/19
SELECT '-1/2'::rational::float;          -- -0.5
SELECT 1::rational;                       -- 1/1
```

### Dynamic Row Ordering

A key use case is maintaining sortable row order without renumbering:

```sql
CREATE TABLE todos (
    prio rational UNIQUE DEFAULT nextval('todos_seq')::integer,
    what text NOT NULL
);

-- Insert between items at priority 1 and 2
INSERT INTO todos VALUES (rational_intermediate('1', '2'), 'new task');
-- prio becomes 3/2, no other rows affected
```

### Index Support

Btree and hash indexes are supported for `rational` columns.
