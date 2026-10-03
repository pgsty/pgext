## Usage

Sources:

- [Official setup.sh](https://github.com/mradul-001/dbisFinal/blob/b1aa6af716ce5c906c0d18766258277d2a64fd64/setup.sh)
- [Official smart_router.control](https://github.com/mradul-001/dbisFinal/blob/b1aa6af716ce5c906c0d18766258277d2a64fd64/dbis-project/contrib/smart_router/smart_router.control)
- [Official Makefile](https://github.com/mradul-001/dbisFinal/blob/b1aa6af716ce5c906c0d18766258277d2a64fd64/dbis-project/contrib/smart_router/Makefile)
- [Official smart_router.c](https://github.com/mradul-001/dbisFinal/blob/b1aa6af716ce5c906c0d18766258277d2a64fd64/dbis-project/contrib/smart_router/smart_router.c)
- [Official smart_router--1.0.sql](https://github.com/mradul-001/dbisFinal/blob/b1aa6af716ce5c906c0d18766258277d2a64fd64/dbis-project/contrib/smart_router/smart_router--1.0.sql)

`smart_router` 1.0 is a research preload module for a fixed two-cluster cache demonstration. It is built inside the repository’s PostgreSQL tree. The install SQL contains no SQL objects; the official setup activates the module by preload, without a database-level extension creation step.

### Enablement Boundary

```conf
shared_preload_libraries = 'smart_router'
```

A restart registers an executor hook and a schema-synchronization background worker. The supplied setup starts a remote cluster on port 6000 with database `company_remote` and a local cluster on port 6001. The source hardcodes the remote login as `mradul` and the worker’s local database as `postgres`; these are demonstration assumptions, not configurable routing APIs.

### Behavior and Limits

The module fetches remote table definitions and data through libpq, populates local tables and tracks five table entries with an LRU policy. Executor hooks inspect query text and attempt remote or cached handling. There are no documented management functions or GUCs for changing the fixed connection and cache layout.

Use only disposable demonstration clusters and trusted input. Eviction issues DELETE against local tables, fetched values are concatenated into SQL without escaping, and the query recognizer is not a general SQL parser. Do not point the module at authoritative local tables or assume transparent routing, transaction consistency or safe multi-user operation. No license or stock PostgreSQL version compatibility is declared for this module.
