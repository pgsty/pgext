## Usage

Sources:

- [1.0.2 release](https://github.com/bitsinfra/gnomon/releases/tag/v1.0.2)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/src/c/README.md)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/src/c/gnomon_c.control)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/LICENSE)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/src/c/gnomon_c--1.0.2.sql)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/docs/installation.md)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/docs/security.md)

gnomon_c 1.0.2 is Gnomon's optional C companion for PostgreSQL 15–18. It depends on gnomon. The C versioning trigger can be used without preload, but routing hooks, shared-memory optimizations and the heartbeat worker require preload and a restart.

### Enablement

```ini
shared_preload_libraries = 'gnomon_c'
gnomon.heartbeat_databases = 'app'
gnomon.heartbeat_interval = '1s'
```

```sql
CREATE EXTENSION gnomon;
CREATE EXTENSION gnomon_c;
SELECT gnomon.use_c_trigger('public.orders'::regclass, true);
```

The table must already be tracked and the caller must hold its management rights. The C trigger supports known-time, row-mode tables without wide keys; unsupported combinations are rejected. Configure the actual database name and preserve other preload entries. Creation and shared configuration require administrator control.

### Operations

With preloading, hooks route branch reads/writes and provide additional locking, point-read and upsert behavior. Keep the Gnomon security model: never grant internal reader/search roles to applications. Upgrade the SQL core first, then gnomon_c with matching files and restart processes that loaded replaced libraries. The companion does not make unlogged scratch history durable or remove retention and heartbeat requirements.

### 1.0.2 correction

The SQL upgrade repairs read sets recorded by 1.0.1 when a branch predates table tracking. Merges now restore the session's selected branch. The C companion rejects system-column reads on branches instead of crashing. Update the SQL core to 1.0.2 first, then the matching C companion if installed; restart processes using a replaced preloaded library.
