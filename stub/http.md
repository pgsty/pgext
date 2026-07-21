## Usage

Sources:

- [pgsql-http v1.7.2 README](https://github.com/pramsey/pgsql-http/blob/v1.7.2/README.md)
- [Extension control file](https://github.com/pramsey/pgsql-http/blob/v1.7.2/http.control)
- [v1.7.1 to v1.7.2 comparison](https://github.com/pramsey/pgsql-http/compare/v1.7.1...v1.7.2)

`http` lets PostgreSQL issue synchronous HTTP requests through libcurl. It is useful for controlled integrations and administrative calls, but the backend waits for the remote service inside the SQL statement and transaction. Restrict who can call it, set short timeouts, and do not let untrusted input choose arbitrary URLs.

### Core Workflow

```sql
CREATE EXTENSION http;

SELECT status, content_type, content
FROM http_get('https://httpbingo.org/get');
```

Send JSON and inspect the response:

```sql
SELECT status, content::jsonb
FROM http_post(
  'https://httpbingo.org/post',
  '{"event":"invoice.paid"}',
  'application/json'
);
```

The generic entry point accepts a complete request:

```sql
SELECT (http((
  'GET',
  'https://httpbingo.org/headers',
  http_headers('Authorization', 'Bearer example'),
  NULL,
  NULL
)::http_request)).status;
```

### Important Objects

- `http_request` contains `method`, `uri`, `headers`, `content_type`, and `content`.
- `http_response` contains `status`, `content_type`, `headers`, and `content`.
- `http_header`, `http_header(...)`, and `http_headers(...)` build request headers; `unnest(response.headers)` exposes response headers as rows.
- `http(...)` executes a complete `http_request`.
- `http_get`, `http_post`, `http_put`, `http_patch`, `http_delete`, and `http_head` are convenience wrappers.
- `urlencode(text)` and `urlencode(jsonb)` encode query data.
- `http_set_curlopt`, `http_list_curlopt`, and `http_reset_curlopt` manage supported session-level libcurl settings.

### Timeouts, Connections, and Security

Each request uses a fresh connection by default. Enable persistent connections only after measuring backend lifetime and remote-server behavior:

```sql
SET http.curlopt_timeout_ms = 1000;
SET http.curlopt_connecttimeout_ms = 250;
SET http.curlopt_tcp_keepalive = 1;
```

The default request timeout is five seconds. A timeout raises a SQL error, so callers must handle transaction rollback. Network latency in triggers or long transactions can hold locks and exhaust database connections; prefer an outbox plus an external worker for durable asynchronous delivery.

Keep TLS verification enabled, protect credential-bearing curl settings, validate response status and content before use, and limit outbound destinations at both SQL privilege and network layers. Version 1.7.2 contains build, test, and curl-option constant maintenance relative to 1.7.1; it does not introduce a material SQL API change. The control file still declares SQL extension version 1.7.
