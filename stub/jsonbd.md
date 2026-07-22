## Usage

Sources:

- [Official jsonbd README](https://github.com/postgrespro/jsonbd/blob/6a8aefcca9999596d75212a0f68c28590b5c3570/README.md)
- [Version 0.1 extension SQL](https://github.com/postgrespro/jsonbd/blob/6a8aefcca9999596d75212a0f68c28590b5c3570/jsonbd--0.1.sql)
- [Compression implementation and preload settings](https://github.com/postgrespro/jsonbd/blob/6a8aefcca9999596d75212a0f68c28590b5c3570/jsonbd.c)
- [PostgreSQL column-compression patch proposal](https://commitfest.postgresql.org/15/1294/)

`jsonbd` is an unfinished experiment that compresses repeated JSONB object keys through a shared dictionary. It does not work with an unpatched stock PostgreSQL server: the extension depends on an abandoned column-compression access-method proposal and its non-core SQL syntax.

### Experimental setup

After building a compatible PostgreSQL source tree with the referenced patch, preload the module and restart the server:

```conf
shared_preload_libraries = 'jsonbd'
jsonbd.workers_count = 1
jsonbd.queue_size = 1
```

Then install the extension and select its compression access method for a JSONB column:

```sql
CREATE EXTENSION jsonbd;

CREATE TABLE jsonbd_demo (
    payload jsonb COMPRESSION jsonbd
);
```

The extension creates the `jsonbd` compression access method and a `jsonbd_dictionary` table keyed by compression-object OID, dictionary ID, and JSON key. Background workers coordinate dictionary lookups through shared memory queues. JSON scalar values are not compressed by this scheme.

### Settings and limitations

- `jsonbd.workers_count` controls the background-worker count; `0` disables workers and the source default is `1`.
- `jsonbd.queue_size` configures each shared-memory queue in kilobytes, from `1` KiB to `1024` KiB; the source default is `1` KiB.

Both shared-memory sizing and worker registration happen during preload, so configuration changes require a server restart. Loading the library outside `shared_preload_libraries` raises an error.

The upstream project explicitly says it is unfinished. Its drop callback has no cleanup behavior, so dictionary lifecycle and orphaned rows are unresolved. There is no supported stock-PostgreSQL compatibility target, stable on-disk contract, or production migration path. Use `jsonbd` only to study the historical patch on a disposable cluster, and test restart, crash, backup, restore, and object-drop behavior before trusting any data to it.
