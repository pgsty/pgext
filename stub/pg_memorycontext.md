## Usage

Sources:

- [Pinned upstream README](https://github.com/jackgo73/pg_memorycontext/blob/ac9dc1d208f712dcdee2bb5dffada347148b4b80/README.md)
- [Version 1.0 installation SQL](https://github.com/jackgo73/pg_memorycontext/blob/ac9dc1d208f712dcdee2bb5dffada347148b4b80/pg_memorycontext--1.0.sql)
- [Pinned private-memory implementation](https://github.com/jackgo73/pg_memorycontext/blob/ac9dc1d208f712dcdee2bb5dffada347148b4b80/pg_memorycontext.c)
- [Formal PGXN distribution](https://pgxn.org/dist/pg_memorycontext/)

pg_memorycontext version 1.0 aggregates the current backend's memory-context tree by context name. Its view returns how many contexts share each name and the combined allocated block size in bytes, ordered by total size and count.

### Configuration and query

Upstream instructs users to preload the library and restart before creating it:

```conf
shared_preload_libraries = 'pg_memorycontext'
```

```sql
CREATE EXTENSION pg_memorycontext;

SELECT memorycontextname, count, totalsize
FROM pg_memorycontext
ORDER BY totalsize DESC, count DESC;
```

The result describes only the backend executing the query. It is a point-in-time diagnostic and cannot measure another session or provide historical retention.

Despite that dynamic behavior, the installation SQL declares `pg_memorycontext()` `IMMUTABLE`; do not treat that declaration as a caching or planner-safety guarantee. It also exposes `totalsize` as a 32-bit `integer` even though the C code accumulates into a wider `long`, so an aggregate above 2 GiB can fail conversion instead of being reported.

### Unsafe compatibility boundary

The source duplicates a private AllocSetContext layout, casts every traversed context to that structure, and follows private first-child and next-child pointers from TopMemoryContext. PostgreSQL has multiple memory-context implementations and their layouts change between server versions. Upstream's broad PostgreSQL 9.2+ claim is not a reliable modern compatibility guarantee; a mismatched build can misread memory or crash the backend.

The install SQL grants SELECT on the diagnostic view to PUBLIC while revoking the underlying function. Ordinary users can still invoke it through the view and observe backend-internal allocation names and sizes. Revoke the view grant and limit it to a diagnostics role. Rebuild against the exact PostgreSQL server headers, validate representative workloads under a crash-safe test environment, and do not preload this legacy 2018 code on production solely on the strength of the README claim.
