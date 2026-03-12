


## Usage

> [rum: RUM index access method](https://github.com/postgrespro/rum)

RUM is an index access method that extends GIN by storing additional information in the posting tree. This enables direct access to positional data, avoiding extra heap scans for ranking, phrase searches, and timestamp ordering.

### Index Creation

```sql
CREATE INDEX idx ON table_name USING rum (column operator_class);
```

With addon operators (e.g., ordering by a timestamp alongside full-text search):

```sql
CREATE INDEX tsts_idx ON tsts USING rum (t rum_tsvector_addon_ops, d)
    WITH (attach = 'd', to = 't');
```

### Operator Classes

| Operator Class | Description |
|---------------|-------------|
| `rum_tsvector_ops` | Stores tsvector lexemes with positions. Supports `<=>` ordering and prefix search. |
| `rum_tsvector_hash_ops` | Stores hashed tsvector lexemes with positions. Supports `<=>` ordering, no prefix search. |
| `rum_tsvector_addon_ops` | Combines tsvector with additional fields (timestamps, integers, etc.) for filtering and ordering. |
| `rum_tsvector_hash_addon_ops` | Hashed variant supporting addon fields, no prefix search. |
| `rum_tsquery_ops` | Stores tsquery branches for fast query matching against indexed documents. |
| `rum_anyarray_ops` | Indexes array types. Supports `&&`, `@>`, `<@`, `=`, `%` and `<=>` ordering. |
| `rum_anyarray_addon_ops` | Combines array elements with additional fields. |
| `rum_TYPE_ops` | Generic ops for int2, int4, int8, float4, float8, money, oid, time, timetz, date, interval, macaddr, inet, cidr, text, varchar, char, bytea, bit, varbit, numeric, timestamp, timestamptz. |

### Ordering Operators

| Operator | Description |
|----------|-------------|
| `<=>` | Distance operator for tsvector, timestamp, numeric types, arrays |
| `<=\|` | Left-side distance for timestamp, int, float, money, oid |
| `\|=>` | Right-side distance for timestamp, int, float, money, oid |

### Examples

Full-text search with ranking:

```sql
SELECT t, a <=> to_tsquery('english', 'beautiful | place') AS rank
FROM test_rum
WHERE a @@ to_tsquery('english', 'beautiful | place')
ORDER BY a <=> to_tsquery('english', 'beautiful | place');
```

Timestamp-ordered full-text search:

```sql
SELECT id, d, d <=> '2016-05-16 14:21:25' FROM tsts
WHERE t @@ 'wr&qh'
ORDER BY d <=> '2016-05-16 14:21:25'
LIMIT 5;
```

Array matching with distance ordering:

```sql
SELECT * FROM test_array
WHERE i && '{1}'
ORDER BY i <=> '{1}' ASC;
```
