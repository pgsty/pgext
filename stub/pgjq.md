


## Usage

> [pgjq: Use jq JSON processing language in PostgreSQL](https://github.com/Florents-Tselai/pgJQ)

Provides a `jqprog` data type for jq programs and a `jq()` function to execute them on `jsonb` objects.

### Basic Filtering

```sql
SELECT jq('[{"bar": "baz", "balance": 7.77}]'::jsonb, '.[0].bar');
-- "baz"
```

### Using the `@@` Operator

```sql
SELECT '[{"bar": "baz"}]' @@ '.[0].bar'::jqprog;
-- "baz"
```

### Complex Programs

```sql
SELECT jq('[true,false,[5,true,[true,[false]],false]]',
          '(..|select(type=="boolean")) |= if . then 1 else 0 end');
-- [1, 0, [5, 1, [1, [0]], 0]]

SELECT jq('[1,5,3,0,7]', '(.[] | select(. >= 2)) |= empty');
-- [1, 0]
```

### Passing Arguments

Pass dynamic arguments as a `jsonb` object, referenced as `$var`:

```sql
SELECT jq(
    '{"jobs": [{"id": 9, "ok": true}, {"id": 100, "ok": false}]}'::jsonb,
    '.jobs[] | select(.ok == $ok and .id == 100) | .',
    '{"ok": false}'
);
```

### Chaining with jsonpath

```sql
SELECT jq('[{"cust":"baz","active":true,"trans":{"balance":100}}]',
          '(.[] | select(.active == true))') - '{trans}' @> '{"cust": "baz"}';
-- t
```

### Working with Files

```sql
SELECT jq(pg_read_file('/path/to/data.json'), '.[]');
```
