## Usage

Sources:

- [Version 1.1 README](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/readme.md)
- [Extension control file](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/pg_https.control)
- [Version 1.1 SQL API](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/sql/pg_https--1.1.sql)
- [SQL entry point and GUC definitions](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/src/pg_https.c)
- [libcurl request, retry, and response implementation](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/src/https_core.c)

`pg_https` 1.1 is a synchronous native HTTP/HTTPS client for PostgreSQL. It exposes libcurl through functions in the `requests` schema, returning status, headers, body, timing, and curl error details; outbound effects are external to the database transaction.

### Preload and Configure

Upstream instructs administrators to add `pg_https` to `shared_preload_libraries` and restart PostgreSQL before creating the extension:

```conf
shared_preload_libraries = 'pg_https'
```

```sql
CREATE EXTENSION pg_https;

SET pg_https.timeout = 10;
SET pg_https.connect_timeout = 5;
SET pg_https.tls_version = 13;
SET pg_https.verify_peer = true;
SET pg_https.max_response_size = 10485760;
SET pg_https.default_headers = '{"User-Agent":"pg_https/1.1"}';
```

Other user-settable GUCs are `pg_https.ca_file`, `pg_https.tcp_keepalive`, `pg_https.http_version`, and `pg_https.connection_reuse`. Keep certificate and hostname verification enabled; the custom CA file must be readable by the PostgreSQL operating-system account.

### Core Workflow

```sql
SELECT status, headers, body, curl_error_code, error_message
FROM requests.get_request(
  'https://api.example.invalid/health',
  '{"Accept":"application/json"}'::jsonb,
  NULL,
  5
);

SELECT status, body, duration_ms
FROM requests.rest_request(
  'POST',
  'https://api.example.invalid/events',
  '{"Content-Type":"application/json"}'::jsonb,
  '{"event":"refresh"}',
  5,
  NULL,
  NULL,
  0,
  300,
  2.0,
  0
);
```

The `http_response` result fields are `status`, `headers`, `body`, `duration_ms`, `response_bytes`, `curl_error_code`, and `error_message`. Response headers are lower-cased into JSON arrays, with the complete header block also retained under a raw field.

### Main Objects and Retry Semantics

- `requests.rest_request` accepts method, URL, JSONB headers, body, timeout override, optional Basic Auth credentials, retry controls, and `cancel_mode`.
- Convenience wrappers are `requests.get_req`, `requests.get_request`, `requests.post_req`, `requests.post_request`, `requests.put_req`, `requests.put_request`, `requests.delete_req`, and `requests.delete_request`.
- `requests.auth_bearer` builds an Authorization header object.
- With `cancel_mode` set to 0, query cancellation aborts the transfer; 1 tells the progress callback to wait for the request instead.

Retries are limited to idempotent `GET`, `HEAD`, `PUT`, `DELETE`, and `OPTIONS` requests. They occur for selected transient libcurl failures and for HTTP `408`, `429`, or `5xx` responses. A `POST` is not retried even when its retry count is nonzero. Requests run synchronously in the calling backend, and retry backoff extends that backend occupancy.

### Security and Transaction Boundaries

This extension has no destination allowlist. A caller can select arbitrary URLs, follow redirects, send write methods, set Basic or bearer credentials, and alter user-settable TLS behavior. That creates an SSRF and data-exfiltration boundary to localhost, cloud metadata endpoints, internal services, and the public internet. Revoke `EXECUTE` from `PUBLIC` on the request functions, grant a narrow role only, and enforce outbound network policy outside PostgreSQL.

Never disable `pg_https.verify_peer` in production or place secrets in URLs, query text, default headers, or ad-hoc SQL arguments. The implementation logs every curl attempt and logs method plus URL for slow or failed requests; ordinary PostgreSQL statement logging can expose still more query text.

An HTTP request may succeed even if the surrounding SQL transaction later rolls back. Design idempotency and reconciliation at the remote API. `pg_https.max_response_size` bounds response buffering, but it does not make untrusted endpoints safe or asynchronous. Version 1.1 is documented as tested on PostgreSQL 17 and depends on the installed libcurl build; validate cancellation, redirects, DNS, TLS, retries, memory limits, and failure behavior on the exact server before use.
