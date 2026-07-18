## Usage

Sources:

- [Project README](https://github.com/japinli/pg_dbsm/blob/4d8f0928f3cc4923d4ba6b6d1c358a8fdba7d00e/README.md)
- [Extension control file](https://github.com/japinli/pg_dbsm/blob/4d8f0928f3cc4923d4ba6b6d1c358a8fdba7d00e/pg_dbsm.control)
- [Background-worker implementation](https://github.com/japinli/pg_dbsm/blob/4d8f0928f3cc4923d4ba6b6d1c358a8fdba7d00e/pg_dbsm.c)

`pg_dbsm` 0.0.1 is a background-worker extension that periodically records the size of every database in the cluster. The worker connects to one configured database, creates an unqualified `dbsm` table there, and appends each database name, collection timestamp, total size, and size increment.

### Configure the worker

```conf
shared_preload_libraries = 'pg_dbsm'
dbsm.database = 'postgres'
dbsm.naptime = 86400
```

Restart PostgreSQL after adding `pg_dbsm` to `shared_preload_libraries`. `dbsm.database` selects the storage database and defaults to `postgres`; `dbsm.naptime` is the collection interval in seconds and defaults to 86400. Then inspect the worker-created table in that database:

```sql
CREATE EXTENSION pg_dbsm;

SELECT datname, created, datsize, incsize
FROM dbsm
ORDER BY created DESC;
```

### Ownership and failure behavior

The version 0.0.1 installation SQL is effectively empty: `CREATE EXTENSION` neither registers the background worker nor creates or owns the `dbsm` table. Preloading the library is what registers the worker, and the worker creates the table and index. Consequently, dropping the extension does not stop a still-preloaded library or clean up those objects.

The source configures the worker not to restart automatically after failure. Monitor its process and PostgreSQL log, protect the central table, and define retention because it grows indefinitely. The project has no published modern compatibility matrix; verify this old C background-worker code on the exact PostgreSQL major version before production use.
