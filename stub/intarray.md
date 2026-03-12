

## Usage

> [intarray: integer array functions and operators with index support](https://www.postgresql.org/docs/current/intarray.html)

Provides functions and operators for manipulating null-free integer arrays, with GiST and GIN index support for fast array searches.

```sql
CREATE EXTENSION intarray;
```

### Functions

| Function | Description | Example |
|---|---|---|
| `icount(int[])` | Number of elements | `icount('{1,2,3}')` -- 3 |
| `sort(int[], dir)` | Sort array (`'asc'` or `'desc'`) | `sort('{3,1,2}','asc')` -- `{1,2,3}` |
| `sort_asc(int[])` | Sort ascending | `sort_asc('{3,1,2}')` -- `{1,2,3}` |
| `sort_desc(int[])` | Sort descending | `sort_desc('{3,1,2}')` -- `{3,2,1}` |
| `uniq(int[])` | Remove adjacent duplicates | `uniq(sort('{1,2,3,2,1}'))` -- `{1,2,3}` |
| `idx(int[], item)` | Index of first match | `idx('{11,22,33}', 22)` -- 2 |
| `subarray(int[], start, len)` | Extract sub-array | `subarray('{1,2,3,4}', 2, 2)` -- `{2,3}` |
| `intset(int)` | Make single-element array | `intset(42)` -- `{42}` |

### Operators

| Operator | Description |
|---|---|
| `&&` | Arrays overlap (have common elements) |
| `@>` | Left array contains right |
| `<@` | Left array is contained in right |
| `#` | Number of elements |
| `+` | Array concatenation / append element |
| `-` | Remove elements |
| `\|` | Union of arrays |
| `&` | Intersection of arrays |
| `@@` | Array matches a query expression |
| `~~` | Query expression matches array |

### Index Support

```sql
-- GiST index for array containment/overlap queries
CREATE INDEX idx ON messages USING GIST (tags gist__intbig_ops);

-- GIN index (alternative)
CREATE INDEX idx ON messages USING GIN (tags gin__int_ops);

-- Query with index support
SELECT * FROM messages WHERE tags && '{1,2}';     -- overlap
SELECT * FROM messages WHERE tags @> '{1,2}';     -- contains
SELECT * FROM messages WHERE tags @@ '1&(2|3)';  -- query expression
```
