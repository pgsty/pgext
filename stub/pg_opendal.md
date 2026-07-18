## Usage

Sources:

- [Pinned upstream README](https://github.com/yihong0618/pg_opendal/blob/71f78ac7414c82425578060a5d1d233ce696a957/README.md)
- [Pinned Rust implementation](https://github.com/yihong0618/pg_opendal/blob/71f78ac7414c82425578060a5d1d233ce696a957/src/lib.rs)
- [Pinned Cargo metadata](https://github.com/yihong0618/pg_opendal/blob/71f78ac7414c82425578060a5d1d233ce696a957/Cargo.toml)
- [Pinned extension control file](https://github.com/yihong0618/pg_opendal/blob/71f78ac7414c82425578060a5d1d233ce696a957/pg_opendal.control)

pg_opendal version 0.0.0 exposes selected Apache OpenDAL storage operations as SQL functions. The reviewed build enables local filesystem, S3, and in-memory services and provides read, write, existence, delete, stat, create-directory, list, copy, rename, and capability calls.

### Isolated filesystem example

Create a server-owned disposable directory first, then use paths relative to that root:

```sql
CREATE EXTENSION pg_opendal;

SELECT pg_opendal_write(
    'fs',
    'demo.txt',
    'hello',
    '{"root":"/tmp/pg-opendal-sandbox"}'::jsonb
);

SELECT pg_opendal_read(
    'fs',
    'demo.txt',
    '{"root":"/tmp/pg-opendal-sandbox"}'::jsonb
);
```

Service configuration must be a JSON object whose values are strings. Read returns UTF-8 text rather than arbitrary bytes.

### Filesystem, credential, and backend risks

Each call builds a new OpenDAL operator and Tokio runtime, then blocks the PostgreSQL backend until storage I/O completes. Reads buffer the whole object and require valid UTF-8; writes accept complete text values; directory listing performs an additional stat request per entry. Large objects, slow endpoints, retries, and large directories can consume backend memory and connection time.

The local filesystem service acts with the PostgreSQL operating-system account and includes destructive write, delete, copy, and rename operations. S3 configuration passes access keys and secret keys as SQL arguments, making secrets vulnerable to query text, activity views, logs, audit systems, and error reporting. The generated functions have no application authorization model. Revoke execution from PUBLIC, expose fixed-service SECURITY DEFINER wrappers only after a search-path review, keep credentials outside caller SQL where possible, and enforce filesystem and network sandboxes.

Upstream labels the project as still in progress. The 0.0.0 snapshot uses pgrx 0.14.3 and declares PostgreSQL 13 through 17, not 18. Test cancellation, timeouts, partial writes, path traversal, symlinks, endpoint redirects, credential rotation, backend crashes, and provider consistency before any non-disposable use.
