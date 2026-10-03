## Usage

Sources:

- [Official alohadb_cache.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cache/alohadb_cache.control)
- [Official alohadb_cache--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cache/alohadb_cache--1.0.sql)
- [Official alohadb_cache.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cache/alohadb_cache.c)
- [Official cache_store.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cache/cache_store.c)
- [Official alohadb_cache.h](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cache/alohadb_cache.h)

`alohadb_cache` 1.0 is a shared-memory JSONB key/value cache in the AlohaDB source tree. It supports expiration and LRU eviction. These notes describe that fork’s implementation, without asserting compatibility with stock PostgreSQL.

### Enable and Use

Preload the library and restart before creating the extension as a superuser:

```conf
shared_preload_libraries = 'alohadb_cache'
alohadb.cache_max_entries = 1000
```

```sql
CREATE EXTENSION alohadb_cache;
SELECT public.cache_set('demo:item', '{"value":42}'::jsonb, interval '1 minute');
SELECT public.cache_get('demo:item');
SELECT * FROM public.cache_stats();
SELECT public.cache_delete('demo:item');
```

The control file fixes objects in `public`. `cache_set` replaces a value and its TTL; `cache_get` returns NULL for a missing or expired entry. `cache_delete` reports whether an entry existed, while `cache_flush` clears the entire cache and returns a count.

### Capacity and Isolation

`alohadb.cache_max_entries` is a startup-only setting, default 1000, range 16–100000. Keys must fit a 256-byte buffer including the terminator; JSONB text must fit the 8192-byte value buffer. Expired entries are evicted on access and capacity pressure evicts the least recently used entry.

This cache is volatile, not WAL-backed or transactionally rolled back. Its key contains no database or role identifier, so callers across databases share a cluster-wide namespace. The SQL script does not revoke default PUBLIC function execution. Restrict function privileges before use, use explicit key namespaces, and never treat this cache as an authorization boundary or durable source of truth.
