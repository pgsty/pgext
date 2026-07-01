
## Usage

Sources: [README](https://github.com/pramsey/pgsql-http/blob/v1.7.1/README.md), [v1.7.1 release](https://github.com/pramsey/pgsql-http/releases/tag/v1.7.1)

`http` lets SQL code make HTTP requests through libcurl. Use it for controlled integration points such as triggers that notify an external service, SQL jobs that fetch a small remote payload, or database-side webhook calls.

```sql
CREATE EXTENSION http;
```

### Request And Response Types

Every request uses `http_request` and returns `http_response`:

```text
http_request(method http_method, uri varchar, headers http_header[], content_type varchar, content varchar)
http_response(status integer, content_type varchar, headers http_header[], content varchar)
```

Convenience wrappers call the same underlying `http(http_request)` function:

- `http_get(uri varchar)`
- `http_get(uri varchar, data jsonb)`
- `http_post(uri varchar, content varchar, content_type varchar)`
- `http_post(uri varchar, data jsonb)`
- `http_put(uri varchar, content varchar, content_type varchar)`
- `http_patch(uri varchar, content varchar, content_type varchar)`
- `http_delete(uri varchar)`
- `http_head(uri varchar)`

### Examples

```sql
SELECT status, content_type, content
FROM http_get('https://httpbun.com/ip');

SELECT content::json->'headers'->>'Authorization'
FROM http((
  'GET',
  'https://httpbun.com/headers',
  http_headers('Authorization', 'Bearer token'),
  NULL,
  NULL
)::http_request);

SELECT status, content::json->'form' AS form
FROM http_post(
  'https://httpbun.com/post',
  jsonb_build_object('myvar', 'myval', 'foo', 'bar')
);

SELECT status, content_type, content::json->>'data' AS data
FROM http_put('https://httpbun.com/put', 'some text', 'text/plain');
```

Inspect response headers by unnesting the `headers` array:

```sql
SELECT (unnest(headers)).*
FROM http_get('https://httpbun.com/');
```

### Binary Content

The README warns that `varchar::bytea` is not safe for binary response bodies because it stops at zero-valued bytes. Use `text_to_bytea(content)` for response content and `bytea_to_text(bytea)` when sending binary request bodies.

```sql
WITH http AS (
  SELECT * FROM http_get('https://httpbingo.org/image/png')
)
SELECT content_type, length(text_to_bytea(content)) AS bytes
FROM http;
```

### Timeout And Version Notes

`pg_http` 1.7.1 is a compatibility and documentation release: it adds timeout examples, adds PostgreSQL 17 wait-event hooks, and includes PostgreSQL 19 support fixes. The user-facing SQL API remains the README surface above.
