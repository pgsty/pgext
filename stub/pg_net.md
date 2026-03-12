


## Usage

> [pg_net: Async HTTP/HTTPS requests in SQL](https://github.com/supabase/pg_net)

The extension requires `shared_preload_libraries = 'pg_net'` in `postgresql.conf`.

### GET Request

```sql
SELECT net.http_get('https://postman-echo.com/get?foo=bar') AS request_id;
```

With URL-encoded params and headers:

```sql
SELECT net.http_get(
  'https://postman-echo.com/get',
  params := '{"foo": "bar"}'::JSONB,
  headers := '{"API-KEY": "<key>"}'::JSONB
) AS request_id;
```

### POST Request

```sql
SELECT net.http_post(
    'https://postman-echo.com/post',
    '{"key": "value"}'::JSONB,
    headers := '{"Content-Type": "application/json"}'::JSONB
) AS request_id;
```

Send a table row as payload:

```sql
WITH row AS (SELECT * FROM my_table LIMIT 1)
SELECT net.http_post(
    'https://api.example.com/data',
    to_jsonb(row.*)
) AS request_id
FROM row;
```

### DELETE Request

```sql
SELECT net.http_delete('https://api.example.com/resource/42') AS request_id;
```

### Checking Responses

```sql
SELECT * FROM net._http_response;
```

### Configuration

```sql
SHOW pg_net.batch_size;       -- default: 200, max rows processed per cycle
SHOW pg_net.ttl;              -- default: 6 hours, response retention time
SHOW pg_net.database_name;    -- default: 'postgres'
```

Modify settings:

```sql
ALTER SYSTEM SET pg_net.ttl TO '1 hour';
ALTER SYSTEM SET pg_net.batch_size TO 500;
SELECT pg_reload_conf();
```
