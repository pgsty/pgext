## Usage

Sources:

- [Official upstream README](https://gitlab.com/pierre_forstmann/pg_documents/-/blob/main/README.md)
- [Official extension control file](https://gitlab.com/pierre_forstmann/pg_documents/-/blob/main/pg_documents.control)
- [Official project page](https://gitlab.com/pierre_forstmann/pg_documents)

`pg_documents` embeds an HTTP daemon in PostgreSQL and exposes tables from one configured database as JSON. It is a preload-only server component, not a SQL extension with a confirmed `CREATE EXTENSION` workflow.

### Core Workflow

After building and installing the library and its `json-c` dependency, configure the target database and listening port:

```ini
shared_preload_libraries = 'pg_documents'
pg_documents.database = 'documents'
pg_documents.port = '8000'
```

Restart PostgreSQL, create test tables in the configured database, and query the daemon:

```sh
curl http://127.0.0.1:8000/_all_dbs
```

The upstream example returns a JSON array of visible table names.

### Important Settings

- `pg_documents.database` selects the single database opened by the worker.
- `pg_documents.port` selects the HTTP listening port.

### Requirements and Caveats

- The reviewed control, registry, or catalog evidence identifies version `1.0.0`.
- The control file marks the extension as non-relocatable.
- The README documents PostgreSQL 15.1 and a locally modified `json-c` 0.16 build. It does not establish compatibility with other combinations.
- Changing preload configuration requires a server restart.
- The reviewed README does not document authentication, authorization, or TLS for the HTTP endpoint. Bind it to a protected interface and validate its security behavior before exposing it beyond an isolated test host.
- Test malformed requests, large results, transaction behavior, database restart, and worker failure before relying on the daemon.
