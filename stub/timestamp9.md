

## Usage

> [timestamp9: nanosecond precision timestamp type for PostgreSQL](https://github.com/optiver/timestamp9)

The `timestamp9` extension provides a nanosecond-precision timestamp type stored as a 64-bit integer (nanoseconds since UNIX epoch).

```sql
CREATE EXTENSION timestamp9;
```

### Data Type

The `timestamp9` type supports timestamps from 1970-01-01 to 2262-04-12 with nanosecond precision.

### Input Formats

```sql
-- Standard PostgreSQL format
SELECT '2019-09-19 08:30:05'::timestamp9;

-- Full nanosecond precision with timezone
SELECT '2019-09-19 08:30:05.123456789 +0200'::timestamp9;

-- Cast from bigint (nanoseconds since epoch)
SELECT 1568878205123456789::bigint::timestamp9;
```

### Type Conversions

- Cast to/from `timestamp` and `timestamptz`
- Cast to/from `date`

Nanosecond precision is preserved throughout conversions.

### Operators

```sql
-- Comparison
SELECT '2019-09-19'::timestamp9 > '2019-09-18'::timestamp9;  -- true

-- Arithmetic with intervals
SELECT '2019-09-19 23:00:00.123456789'::timestamp9 + interval '1 day';

-- Subtraction
SELECT '2019-09-20'::timestamp9 - '2019-09-19'::timestamp9;
```

### Functions

```sql
SELECT greatest('2019-09-19'::timestamp9, '2019-09-20'::timestamp9);
```

### Index Support

Btree and hash indexes are supported.
