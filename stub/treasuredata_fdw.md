## Usage

Sources:

- [Pinned current upstream README](https://github.com/treasure-data/treasuredata_fdw/blob/afbecb707e3757b43a9c7dcfe729721f8269cbd0/README.md)
- [Pinned extension control file](https://github.com/treasure-data/treasuredata_fdw/blob/afbecb707e3757b43a9c7dcfe729721f8269cbd0/treasuredata_fdw.control)
- [Pinned distribution metadata](https://github.com/treasure-data/treasuredata_fdw/blob/afbecb707e3757b43a9c7dcfe729721f8269cbd0/META.json)
- [Pinned option and credential handling](https://github.com/treasure-data/treasuredata_fdw/blob/afbecb707e3757b43a9c7dcfe729721f8269cbd0/option.c)
- [Pinned Rust bridge dependencies](https://github.com/treasure-data/treasuredata_fdw/blob/afbecb707e3757b43a9c7dcfe729721f8269cbd0/bridge_td_client_rust/Cargo.toml)
- [Formal PGXN distribution](https://pgxn.org/dist/treasuredata_fdw/)

treasuredata_fdw connects PostgreSQL to the Treasure Data service using Hive or Presto jobs. Foreign tables can query a remote table or fixed remote SQL, push selected filters and aggregates, stream or download results, insert data in chunks, and be generated with IMPORT FOREIGN SCHEMA.

### Define a read-only foreign table

The reviewed validator accepts credentials and connection options on the foreign table itself:

```sql
CREATE EXTENSION treasuredata_fdw;

CREATE SERVER treasuredata_server
FOREIGN DATA WRAPPER treasuredata_fdw;

CREATE FOREIGN TABLE sample_access (
    time integer,
    host text,
    path text,
    code integer
) SERVER treasuredata_server
OPTIONS (
    apikey '<api-key>',
    database 'sample_datasets',
    query_engine 'presto',
    table 'www_access'
);

SELECT code, count(*)
FROM sample_access
WHERE time BETWEEN 1412121600 AND 1414800000
GROUP BY code;
```

The extension object version is 1.2, matching this catalog. The pinned repository metadata identifies distribution 1.2.16, while the captured PGXN page lists 1.2.15 as its latest stable release. Do not compare those package versions directly with the SQL object version.

### Remote transaction and secret boundaries

The required API key is stored in foreign-table options rather than a protected user mapping. Catalog access, schema dumps, backups, support bundles, IMPORT-generated commands, and administrative logs can expose it. Restrict catalog and table ownership, use a narrowly scoped key, avoid sharing definitions, rotate regularly, and verify the service endpoint.

Queries and imports are synchronous remote jobs that can run for hours while holding a PostgreSQL backend. The optional result-download directory writes compressed files on the database server. INSERT with atomic_import disabled uploads chunks that may become remotely visible before the local statement commits; aborting PostgreSQL does not roll back prior remote imports. Atomic mode stages a remote table and performs a final copy, but it is still not PostgreSQL two-phase commit and can leave remote resources after failures.

The Rust bridge uses old HTTP, time, retry, and Treasure Data client dependencies, contains unwrap and panic paths, and the repository's last source commit is from 2023. Test TLS, cancellation, timeouts, retry duplication, partial import, cleanup, unsupported types, remote SQL injection boundaries, local disk exhaustion, schema drift, and current Treasure Data APIs. Grant foreign-table access only after documenting cost and data-egress policy.
