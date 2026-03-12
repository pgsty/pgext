

## Usage

> [prefix: prefix range type for phone number routing](https://github.com/dimitri/prefix)

The `prefix` extension provides a `prefix_range` data type for efficient prefix matching, particularly useful for telephony call routing.

### Data Type

Create `prefix_range` values using the constructor or text casting:

```sql
CREATE EXTENSION prefix;

SELECT prefix_range('123');           -- 123
SELECT prefix_range('123', '4', '5'); -- 123[4-5]
SELECT '123'::prefix_range;           -- 123
```

### Operators

| Operator | Description |
|----------|-------------|
| `@>` | Contains (prefix contains value) |
| `<@` | Is contained by |
| `&&` | Overlaps |
| `\|` | Union |
| `&` | Intersection |
| `=`, `<>`, `<`, `>`, `<=`, `>=` | Comparison |

### Examples

```sql
-- Find the longest matching prefix for a phone number
SELECT * FROM prefixes
WHERE prefix @> '0123456789'
ORDER BY length(prefix) DESC
LIMIT 1;

-- Containment check
SELECT '123'::prefix_range @> '123456';     -- true

-- Intersection
SELECT '123[4-5]' & '123[2-7]';            -- 123[4-5]

-- Union
SELECT '123' | '124';                       -- 12[3-4]
```

### Index Support

Create a GiST index for efficient prefix lookups:

```sql
CREATE INDEX idx_prefix ON prefixes USING gist(prefix);
```

The index accelerates `@>`, `<@`, `&&`, and `=` operators.
