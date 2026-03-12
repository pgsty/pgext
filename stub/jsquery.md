


## Usage

> [jsquery: data type for jsonb inspection](https://github.com/postgrespro/jsquery)

JsQuery provides a query language for JSONB data, similar to what tsquery does for full-text search. It offers a concise way to search nested objects and arrays with index support via GIN.

### Operators

| Operator | Description |
|----------|-------------|
| `@@` | Match operator: test whether a jsonb value matches a jsquery expression |

### Query Syntax

Expressions follow the pattern `path operator value`:

**Binary operators:**
- `=` (equality), `>`, `>=`, `<`, `<=` (comparison)
- `IN` (list membership)
- `&&` (overlap), `@>` (contains), `<@` (contained in)

**Unary operators:**
- `= *` (existence check)
- `IS ARRAY`, `IS NUMERIC`, `IS OBJECT`, `IS STRING`, `IS BOOLEAN` (type checking)

### Path Expressions

| Symbol | Meaning |
|--------|---------|
| `#` | Any array index |
| `#N` | Specific array index N |
| `%` | Any object key |
| `*` | Any sequence of keys/indexes |
| `@#` | Array/object length |
| `$` | Entire document |

"Every" semantics (all elements must match):
- `#:` -- all array elements
- `%:` -- all object keys
- `*:` -- all nested paths

### Examples

Simple value matching:

```sql
SELECT * FROM jsonb_table WHERE data @@ 'name = "Alice"';
SELECT * FROM jsonb_table WHERE data @@ 'age > 21';
SELECT * FROM jsonb_table WHERE data @@ 'tags.#: IS STRING';
```

Logical combinations:

```sql
SELECT * FROM jsonb_table WHERE data @@ 'a = 1 AND (b = 2 OR c = 3)';
```

Array element matching (find array elements where both conditions hold):

```sql
SELECT * FROM jsonb_table WHERE data @@ '#(a = 1 AND b = 2)';
```

Object key range matching:

```sql
SELECT * FROM jsonb_table WHERE data @@ '%($ >= 10 AND $ <= 20)';
```

### GIN Indexing

Two operator classes for different query patterns:

```sql
-- Best for range and exact searches when full path is known
CREATE INDEX ON jsonb_table USING gin (data jsonb_path_value_ops);

-- Best for exact value searches; supports % and * in paths
CREATE INDEX ON jsonb_table USING gin (data jsonb_value_path_ops);
```

Optimizer hints for index usage:

```sql
SELECT * FROM jsonb_table WHERE data @@ 'x = 1 /*-- index */ AND y = 2';
SELECT * FROM jsonb_table WHERE data @@ 'x = 1 /*-- noindex */ AND y = 2';
```

### Schema Validation via CHECK Constraints

```sql
CREATE TABLE documents (
    id serial PRIMARY KEY,
    data jsonb CHECK (data @@ 'name IS STRING AND similar_ids.#: IS NUMERIC'::jsquery)
);
```
