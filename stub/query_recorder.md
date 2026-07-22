## Usage

Sources:

- [Official README](https://github.com/tvondra/query_recorder/blob/14e6509a0b92b962111647b48aa3fed9e909c060/README)
- [Official extension control file](https://github.com/tvondra/query_recorder/blob/14e6509a0b92b962111647b48aa3fed9e909c060/query_recorder.control)
- [Official PGXN distribution](https://pgxn.org/dist/query_recorder/)

`query_recorder` is a PostgreSQL query-capture module that writes executed SQL and timing metadata to rotating files. Use it for controlled workload analysis or replay preparation, not as a general audit log: it can capture sensitive SQL text, file ordering is not guaranteed, and the reviewed `1.0.1` code is an old preview release.

### Core Workflow

The module allocates shared memory, so add it to `shared_preload_libraries` and restart PostgreSQL. Configure a writable base filename and bounded rotation before enabling capture:

```ini
shared_preload_libraries = 'query_recorder'
query_recorder.filename = '/var/log/postgresql/recorded-query'
query_recorder.max_files = 20
query_recorder.size_limit = '256MB'
query_recorder.buffer_size = '8MB'
query_recorder.enabled = false
query_recorder.normalize = true
```

After restart, install the extension and enable recording only for the sessions or interval you intend to observe:

```sql
CREATE EXTENSION query_recorder;

SET query_recorder.enabled = true;
SELECT current_database();
SET query_recorder.enabled = false;
```

The output uses numeric suffixes such as `.000`. Each record contains a timestamp in microseconds, backend identifier, duration, query length, and query text. When `query_recorder.normalize` is enabled, embedded line endings are replaced with spaces; otherwise consumers must use the recorded length rather than line boundaries to parse SQL safely.

### Configuration Index

- `query_recorder.filename` sets the output base path.
- `query_recorder.max_files` bounds the rotating file set.
- `query_recorder.size_limit` sets the rotation threshold per file.
- `query_recorder.buffer_size` controls the shared output buffer.
- `query_recorder.enabled` starts or stops capture and can be changed with `SET`.
- `query_recorder.normalize` converts query line endings to spaces.

### Operational Boundaries

Provision directory ownership, restrictive file permissions, retention, disk monitoring, and secure disposal before enabling the recorder. The stream may contain credentials, personal data, and very large statements, while concurrent backends may appear out of order. Test the exact build and parser on representative traffic, and keep capture windows short; `query_recorder` does not provide audit-grade integrity, attribution, or tamper resistance.
