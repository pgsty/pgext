


## Usage

> [omni_httpc: HTTP client](https://docs.omnigres.org/omni_httpc/reference/)

The `omni_httpc` extension provides HTTP/1, HTTP/2, and HTTP/3 (experimental) client functionality.

### Prepare and Execute Requests

```sql
SELECT version >> 8 AS http_version, status, headers, convert_from(body, 'utf-8')
FROM omni_httpc.http_execute(
    omni_httpc.http_request('https://example.com'),
    omni_httpc.http_request('https://example.org')
);
```

**`http_request(url, method, headers, body)`** -- Prepares a request. Method defaults to GET.

**`http_execute(VARIADIC requests)`** -- Executes one or more requests. Returns `version`, `status`, `headers`, `body`, and `error` columns.

### Execution Options

```sql
SELECT * FROM omni_httpc.http_execute_with_options(
    omni_httpc.http_execute_options(http2_ratio => 100),
    omni_httpc.http_request('https://example.com')
);
```

| Option                | Type     | Default | Description                        |
|:----------------------|:---------|:--------|:-----------------------------------|
| `http2_ratio`         | smallint | 0       | 0-100, percentage of HTTP/2 use    |
| `http3_ratio`         | smallint | 0       | 0-100, percentage of HTTP/3 use    |
| `force_cleartext_http2` | bool  | false   | Use h2c                            |
| `first_byte_timeout`  | int      | 5000    | Milliseconds                       |
| `timeout`             | int      | 5000    | Milliseconds                       |
| `follow_redirects`    | bool     | true    | Follow HTTP redirects              |
| `cacerts`             | text[]   | null    | PEM-encoded CA certificates        |
| `clientcert`          | client_certificate | null | PEM-encoded client cert   |

The sum of `http2_ratio` and `http3_ratio` must not exceed 100.

### Connection Pool

```sql
SELECT * FROM omni_httpc.http_connections;
-- Returns: http_protocol, url
```
