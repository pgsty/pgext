


## Usage

> [omni_http: Basic HTTP types](https://docs.omnigres.org/omni_httpc/reference/)

The `omni_http` extension provides foundational HTTP types used by `omni_httpc` and `omni_httpd`.

### Key Types

- **`omni_http.http_method`** -- HTTP method enum (GET, POST, PUT, DELETE, etc.)
- **`omni_http.http_headers`** -- Array type for HTTP headers. If header name is null, no header is created; if value is null, it serializes as empty string.
- **`omni_http.http_request`** -- Composite type representing an HTTP request
- **`omni_http.http_response`** -- Composite type representing an HTTP response

These types form the shared interface between the HTTP client (`omni_httpc`) and HTTP server (`omni_httpd`) extensions.
