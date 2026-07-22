## Usage

Sources:

- [Official README](https://github.com/jnidzwetzki/pg_debug_scan/blob/6ac69c0f81de24957130e8dd28f39256056f67ad/README.md)
- [Scan implementation](https://github.com/jnidzwetzki/pg_debug_scan/blob/6ac69c0f81de24957130e8dd28f39256056f67ad/src/lib.rs)
- [Extension control file](https://github.com/jnidzwetzki/pg_debug_scan/blob/6ac69c0f81de24957130e8dd28f39256056f67ad/pg_debug_scan.control)
- [Build feature matrix](https://github.com/jnidzwetzki/pg_debug_scan/blob/6ac69c0f81de24957130e8dd28f39256056f67ad/Cargo.toml)

`pg_debug_scan` performs a raw heap scan using either the current transaction snapshot or caller-supplied snapshot fields. It is an educational MVCC inspection tool that can reveal normally invisible or deleted tuple versions; it is not a time-travel, recovery, or access-control mechanism.

### Core Workflow

```sql
CREATE EXTENSION pg_debug_scan;

-- Use the current transaction snapshot.
SELECT * FROM pg_debug_scan('public.temperature', NULL);

-- Supply xmin:xmax:xip1,xip2 explicitly.
SELECT * FROM pg_debug_scan('public.temperature', '775:775:');
```

`pg_debug_scan(table text, snapshot text DEFAULT NULL)` returns `xmin bigint`, `xmax bigint`, and `data text`. The custom snapshot syntax is `xmin:xmax:xip1,xip2`; each listed in-progress transaction ID must satisfy `xmin <= xip < xmax`.

The function converts every user column through its type output function and serializes the result as JSON-shaped text. All non-null values are JSON strings, and an SQL null is represented by the string `"NULL"`, not a JSON null. Cast and interpret `data` accordingly.

### Security Boundary

The current implementation resolves and opens the relation through internal APIs and does not perform the normal SQL table privilege check. Although extension installation requires superuser, function execution is public by default unless revoked. Immediately restrict it to a dedicated diagnostic role:

```sql
REVOKE EXECUTE ON FUNCTION pg_debug_scan(text, text) FROM PUBLIC;
GRANT EXECUTE ON FUNCTION pg_debug_scan(text, text) TO mvcc_diagnostics;
```

A grantee may otherwise inspect relations it cannot query normally, including obsolete tuple versions that still contain sensitive values. Treat access as equivalent to low-level database inspection and use it only on controlled systems.

### Limitations and Compatibility

The scan works only on heap relations, takes a full `AccessShareLock`, and accumulates every result row in memory before returning. Large relations can cause substantial I/O and backend memory use. User snapshot values are 32-bit transaction IDs, while modern `pg_current_snapshot()` exposes epoch-aware `xid8` values; transaction wraparound and epoch information are not modeled.

The parser validates individual `xip` values but does not validate the overall `xmin`/`xmax` relationship, and malformed numeric input may abort with an error. The source uses PostgreSQL internal heap and snapshot APIs, so each major version needs a matching build. Cargo declares PostgreSQL 11 through 16, but the relation-name code has explicit branches only for 12 through 16; verify the exact target build, especially PostgreSQL 11, before use.
