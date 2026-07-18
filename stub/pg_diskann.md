## Usage

Sources:

- [Official upstream documentation](https://learn.microsoft.com/en-us/azure/postgresql/extensions/how-to-use-pgdiskann)

`pg_diskann` — Azure Database for PostgreSQL extension for DiskANN approximate-nearest-neighbor vector indexes.

The reviewed catalog snapshot records version `0.6.5`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `vector`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_diskann";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_diskann';
```

This is a provider-specific component for `Microsoft Azure`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
