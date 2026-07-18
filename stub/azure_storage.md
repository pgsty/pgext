## Usage

Sources:

- [Official upstream documentation](https://learn.microsoft.com/en-us/azure/postgresql/extensions/concepts-storage-extension)

`azure_storage` — Azure Database for PostgreSQL extension for importing from and exporting to Azure Storage blobs from SQL.

The reviewed catalog snapshot records version `1.9`, kind `preload`, and implementation language `C`.
The curated compatibility set is `12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "azure_storage";
SELECT extversion
FROM pg_extension
WHERE extname = 'azure_storage';
```

This is a provider-specific component for `Microsoft Azure`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
