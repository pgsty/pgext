## Usage

Sources:

- [Official README](https://github.com/shurizzle/pg_url/blob/7fe5192e25b68e2c22b096b74daf84c26bf5e384/README.md)
- [Official extension control file](https://github.com/shurizzle/pg_url/blob/7fe5192e25b68e2c22b096b74daf84c26bf5e384/url.control)
- [Official extension SQL](https://github.com/shurizzle/pg_url/blob/7fe5192e25b68e2c22b096b74daf84c26bf5e384/url--1.0.sql)

`url` (the `pg_url` project) is a pure SQL/PL/pgSQL library for parsing URLs, percent-encoding values, and working with query strings as structured PostgreSQL values. It performs no HTTP request and does not resolve or validate remote hosts.

### Core Workflow

```sql
CREATE EXTENSION url;

SELECT url_encode('hello world/你好');
SELECT url_decode('hello%20world%2F%E4%BD%A0%E5%A5%BD');

SELECT url('https://alice@example.com:8443/docs?q=postgres&q=sql#usage');

SELECT qskv('q=postgres&q=sql&empty') -> 'q';
SELECT qskv('q=postgres&q=sql&empty') ->> 'q';
```

The composite `url` type has `scheme`, `authority`, `domain`, `port`, `path`, `querystring`, and `fragment` fields. Text can be converted with `url(text)` or a cast and reconstructed by casting back to text. `qskv` stores an array of `kventry` values: `->` returns the first value for a key, while `->>` returns all matching values as `text[]`. `unnest(qskv)` emits key/value rows.

### Parsing Boundaries

The implementation is a regular-expression parser written for PostgreSQL 9.4-era syntax, not a complete standards or security validator. Test internationalized domain names, IPv6 literals, unusual user information, invalid percent escapes, duplicate/empty query keys, `+` handling, and round trips before adopting it. Parsed `domain` or `path` values are still untrusted input; applications must separately enforce allowed schemes and destinations to prevent SSRF or open redirects. No network preload, restart, or server filesystem access is required.
