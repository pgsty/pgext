


## Usage

> [omni_httpd: HTTP server](https://docs.omnigres.org/omni_httpd/intro/)

The `omni_httpd` extension is an embedded HTTP server for PostgreSQL, allowing HTTP request handling in SQL.

### Define a Handler

```sql
CREATE FUNCTION my_handler(request omni_httpd.http_request)
    RETURNS omni_httpd.http_outcome
    RETURN omni_httpd.http_response(body => request.headers::text);
```

### Register a Route

```sql
INSERT INTO omni_httpd.urlpattern_router (match, handler)
VALUES (omni_httpd.urlpattern('/headers'), 'my_handler'::regproc);
```

### Configuration

- **`omni_httpd.http_workers`** -- Number of HTTP workers (defaults to CPU count, respects `max_worker_processes`)
- **`omni_httpd.temp_dir`** -- Temporary directory for sockets (defaults to `/tmp`)

### Control

```sql
SELECT omni_httpd.stop();   -- Stop the server
SELECT omni_httpd.start();  -- Start the server
-- Pass immediate => true for immediate execution without waiting for transaction commit
```

### Key Types

- **`omni_httpd.http_request`** -- Request type with `path`, `method`, `query_string`, `headers`, `body`
- **`omni_httpd.http_response`** -- Response constructor
- **`omni_httpd.http_outcome`** -- Return type for handlers
