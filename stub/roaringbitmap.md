

## Usage

> [roaringbitmap: compressed bitmap data type for PostgreSQL](https://github.com/ChenHuajun/pg_roaringbitmap)

The `roaringbitmap` extension provides compressed bitmap data types for efficient set operations on integer collections.

```sql
CREATE EXTENSION roaringbitmap;
SET roaringbitmap.output_format = 'array';

SELECT rb_build('{1,2,3,4,5}'::int[]);  -- {1,2,3,4,5}
```

### Data Types

- **`roaringbitmap`**: 32-bit bitmap, range [0, 4294967296)
- **`roaringbitmap64`**: 64-bit bitmap, range [0, 18446744073709551615)

Output format controlled by: `SET roaringbitmap.output_format = 'array'` or `'bytea'`

### Operators

| Operator | Description |
|----------|-------------|
| `\|` | Bitwise OR (union) |
| `&` | Bitwise AND (intersection) |
| `#` | Bitwise XOR |
| `-` | Difference (ANDNOT) |
| `\|` (with int) | Add element |
| `-` (with int) | Remove element |
| `@>` | Contains |
| `<@` | Is contained by |
| `&&` | Overlap |
| `=`, `<>` | Equality/inequality |

### Core Functions

```sql
-- Construction
SELECT rb_build(ARRAY[1,2,3,4,5]);

-- Analysis
SELECT rb_cardinality(rb_build('{1,2,3}'::int[]));   -- 3
SELECT rb_to_array(rb_build('{1,2,3}'::int[]));      -- {1,2,3}
SELECT rb_iterate(rb_build('{1,2,3}'::int[]));        -- returns set

-- Set operation cardinalities
SELECT rb_and_cardinality(a, b);
SELECT rb_or_cardinality(a, b);
SELECT rb_xor_cardinality(a, b);
SELECT rb_andnot_cardinality(a, b);

-- Range operations
SELECT rb_range(bitmap, 2, 5);   -- extract range [2, 5)
SELECT rb_fill(bitmap, 0, 10);   -- add range [0, 10)
SELECT rb_clear(bitmap, 3, 7);   -- remove range [3, 7)
SELECT rb_flip(bitmap, 0, 10);   -- toggle range [0, 10)

-- Element access
SELECT rb_min(bitmap);            -- minimum element
SELECT rb_max(bitmap);            -- maximum element
SELECT rb_rank(bitmap, 5);        -- count elements <= 5
SELECT rb_index(bitmap, 3);       -- 0-based position of element

-- Utilities
SELECT rb_is_empty(bitmap);
SELECT rb_jaccard_dist(a, b);     -- Jaccard similarity
```

### Aggregate Functions

```sql
SELECT rb_build_agg(id) FROM table;       -- build from rows
SELECT rb_or_agg(bitmap) FROM table;      -- union all bitmaps
SELECT rb_and_agg(bitmap) FROM table;     -- intersect all bitmaps
SELECT rb_xor_agg(bitmap) FROM table;     -- XOR all bitmaps
```

64-bit equivalents use the `rb64_` prefix.
