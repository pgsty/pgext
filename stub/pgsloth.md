## Usage

Sources:

- [Official README](https://github.com/jamessewell/pgsloth/blob/cf5fbf1f68a7527bbe75c3eaf4c28bad92a582a0/README.md)
- [Executor-hook implementation](https://github.com/jamessewell/pgsloth/blob/cf5fbf1f68a7527bbe75c3eaf4c28bad92a582a0/src/lib.rs)
- [Extension control file](https://github.com/jamessewell/pgsloth/blob/cf5fbf1f68a7527bbe75c3eaf4c28bad92a582a0/pgsloth.control)

`pgsloth` version `0.0.0` is a demonstration module that deliberately delays every query executor start by a random 100 through 9,999 milliseconds. It is a test toy for observing latency behavior, not a production extension.

### Enablement

The library installs only an executor hook and has no SQL objects, so do not use `CREATE EXTENSION` as its activation mechanism. Add the library to server configuration and restart PostgreSQL:

```conf
shared_preload_libraries = 'pgsloth'
```

After restart, ordinary queries emit a notice and sleep before execution:

```sql
SELECT 1;
```

### Operational Boundary

The delay applies globally in every backend that loads the hook. There is no GUC, role exemption, relation filter, or per-session switch. The README's relaxed-replication and index features are explicitly marked “to be implemented”; the current source implements only query sleep.

Remove `pgsloth` from `shared_preload_libraries` and restart to disable it. Never enable it on a production or shared server: it adds unpredictable latency to application queries and administrative work, and could amplify connection backlog and timeout failures.
