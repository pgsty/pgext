


## Usage

> [omni_web: Common web stack primitives](https://docs.omnigres.org/omni_web/intro/)

The `omni_web` extension provides web stack utility functions, often used together with `omni_httpd`.

### Cookies

Parse a `Cookie` header into key-value pairs:

```sql
SELECT * FROM omni_web.cookies('PHPSESSID=298zf09hf012fh2; csrftoken=u32t4o3tb3gg43; _gat=1');
```

Returns a table with `name` and `value` columns.

### Query Strings

**`omni_web.parse_query_string(text)`** -- Parses a query string into an array of key-value pairs:

```sql
SELECT omni_web.parse_query_string('key=value&a=1&a=2');
```

**`omni_web.param_get(parsed, key)`** -- Get the first value for a parameter:

```sql
SELECT omni_web.param_get(omni_web.parse_query_string('a=1&a=2'), 'a');  -- '1'
```

**`omni_web.param_get_all(parsed, key)`** -- Get all values for a parameter:

```sql
SELECT omni_web.param_get_all(omni_web.parse_query_string('a=1&a=2'), 'a');
-- 1
-- 2
```
