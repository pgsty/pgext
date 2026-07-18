## Usage

Sources:

- [Last-known official upstream repository; currently unavailable](https://github.com/dmtkfs/pg-tamper-log)

Current-source status: the official repository and owner return HTTP 404. Checks of GitHub repository, fork, code, and commit search, raw and codeload paths, PGXN, and Software Heritage found no current official source or archive.

The frozen catalog historically described `pg_tamperlog_rust` version `1.0.0` as a pgrx helper for accelerated SHA-256 hash-chain calculations used by `pg_tamperlog`. Those details are historical clues only and could not be revalidated after deletion.

```sql
-- Run only if matching, independently verified artifacts are recovered.
CREATE EXTENSION "pg_tamperlog_rust";
```

Do not install a similarly named crate or binary without provenance. Recover and audit the exact source first, verify PostgreSQL and pgrx ABI compatibility, reproduce builds, compare hash results against an independent implementation, and test NULL, encoding, concurrency, upgrade, and crash behavior before trusting it in an audit chain.
