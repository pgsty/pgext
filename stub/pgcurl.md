## Usage

Sources:

- [Official README](https://github.com/pandrewhk/pgcurl/blob/97d1e7e7447adae345fd90695916031237e1b06b/README.md)
- [Extension SQL](https://github.com/pandrewhk/pgcurl/blob/97d1e7e7447adae345fd90695916031237e1b06b/pgcurl--0.0.1.sql)
- [libcurl implementation](https://github.com/pandrewhk/pgcurl/blob/97d1e7e7447adae345fd90695916031237e1b06b/pgcurl.c)

`pgcurl` 0.0.1 exposes a single `curl(text) RETURNS text` function that performs a synchronous libcurl GET inside the PostgreSQL backend. It can fetch a small text resource for controlled experiments, but the abandoned implementation lacks the controls required for production HTTP integration.

### Basic Call

```sql
CREATE EXTENSION pgcurl;

SELECT curl('https://example.com/');
```

The function follows redirects and asks libcurl to decode supported content encodings such as gzip or deflate. It returns only the response body as PostgreSQL text.

### Failure Semantics

Transport errors are appended to the returned text rather than raised as SQL errors. HTTP status, headers, and final URL are not returned, and HTTP error statuses are not distinguished from successful bodies. A partial response can therefore be followed by an error message in the same value.

There are no SQL options for method, headers, credentials, TLS policy, proxy, redirect limit, response-size limit, connect timeout, or total timeout. The call blocks its database backend while DNS, connection, redirects, and response transfer run. Never pass null: the SQL function is not declared strict, while the C code does not handle a null argument.

### Security Boundary

Any role that can execute `curl(text)` can make the database server request attacker-chosen URLs, including loopback, private network, cloud metadata, or redirected destinations. It can also consume backend memory with a large response. Revoke execution from untrusted roles and enforce destination, DNS, egress, timeout, and response-size controls outside this extension.

The project is abandoned and has no maintained PostgreSQL or libcurl compatibility matrix. For production, perform HTTP work in an application or job worker with explicit authentication, observability, retries, circuit breaking, and transaction-independent failure handling.
