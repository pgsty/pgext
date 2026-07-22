## Usage

Sources:

- [Official README](https://github.com/cyga/www_fdw/blob/2445772dddb44be8534527782deb8230f84015fe/README.md)
- [Official control file](https://github.com/cyga/www_fdw/blob/2445772dddb44be8534527782deb8230f84015fe/www_fdw.control)
- [Official installation SQL](https://github.com/cyga/www_fdw/blob/2445772dddb44be8534527782deb8230f84015fe/sql/www_fdw.sql)
- [Official documentation](https://github.com/cyga/www_fdw/wiki/Documentation)
- [Official examples](https://github.com/cyga/www_fdw/wiki/Examples)

`www_fdw` is an abandoned foreign data wrapper that turns a `SELECT` into an HTTP request and converts JSON, XML, or callback-decoded responses into rows. It targets early PostgreSQL FDW APIs and is best treated as historical integration code, not as a maintained HTTP client for current servers.

### Core Workflow

Create the extension, a server containing the service base URI, a user mapping, and a foreign table whose columns can serve as request parameters and response fields.

```sql
CREATE EXTENSION www_fdw;

CREATE SERVER web_service
  FOREIGN DATA WRAPPER www_fdw
  OPTIONS (
    uri 'https://api.example.invalid/items',
    response_type 'json'
  );

CREATE USER MAPPING FOR current_user
  SERVER web_service;

CREATE FOREIGN TABLE web_items (
  category text,
  item_id text,
  title text
) SERVER web_service;

SELECT item_id, title
FROM web_items
WHERE category = 'database';
```

With the default serializer, equality predicates become URL key/value parameters. Other operators require a custom request callback. The automatic JSON or XML decoder searches breadth-first for the first homogeneous array and maps its fields to the foreign-table columns.

### Server Options and Callbacks

- `uri` is the base endpoint; `uri_select` and `method_select` customize the read path and HTTP method.
- `request_user_agent` and `request_user_header` add request headers.
- `request_serialize_callback` receives serialized query conditions and can construct a custom URL or body.
- `request_serialize_type` supports `log`, `null`, `json`, or `xml` representations.
- `response_type` supports `json`, `xml`, or `other`.
- `response_deserialize_callback` parses a response; `response_iterate_callback` can transform each returned tuple.
- `ssl_cert`, `ssl_key`, `cainfo`, and `proxy` pass connection settings to libcurl.

Although option names for insert, update, and delete exist, the official documentation says FDW writes are not implemented.

### Security and Compatibility

Queries cause the database server to make network requests. Restrict FDW and server ownership, validate every endpoint and proxy, and prevent untrusted roles from turning the database into an SSRF path to internal services. Treat response data, request headers, certificate paths, and callback functions as untrusted integration inputs.

Release `0.1.9` documents PostgreSQL 9.2 support and says the master branch's PostgreSQL 9.5 support was untested. The project is no longer maintained, depends on old libcurl/libjson/libxml integration, and provides limited planning behavior. Use timeouts and external isolation when evaluating it, and prefer a maintained service-ingestion design on supported PostgreSQL releases.
