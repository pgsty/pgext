

## Usage

> [seg: line segment / floating-point interval data type](https://www.postgresql.org/docs/current/seg.html)

The `seg` extension provides a data type for representing line segments or floating-point intervals, useful for representing measurements with uncertainty.

```sql
CREATE EXTENSION seg;
```

### Input Formats

```sql
SELECT '5.0'::seg;           -- zero-length interval (point)
SELECT '5.0 .. 7.0'::seg;    -- interval from 5.0 to 7.0
SELECT '5(+-)0.3'::seg;      -- interval 4.7 .. 5.3
SELECT '50 ..'::seg;          -- open interval >= 50
SELECT '.. 0'::seg;           -- open interval <= 0
```

Certainty indicators (`<`, `>`, `~`) can be prepended but are treated as comments and ignored by operators.

### Operators

**Spatial operators (GiST-indexed):**

| Operator | Description |
|----------|-------------|
| `<<` | Entirely left of |
| `>>` | Entirely right of |
| `&<` | Does not extend right of |
| `&>` | Does not extend left of |
| `=` | Equal |
| `&&` | Overlaps |
| `@>` | Contains |
| `<@` | Contained in |

**Comparison operators** (`<`, `<=`, `>`, `>=`) are available for sorting.

### Index Support

```sql
CREATE INDEX idx ON measurements USING gist (reading);
```

### Precision

Values are stored as 32-bit floating-point pairs, retaining up to 7 significant digits. Trailing zeroes are preserved.
