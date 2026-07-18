## Usage

Sources:

- [Official extension control file](https://database.dev/toliver38/steampipe_postgres_aws)
- [Official upstream documentation](https://hub.steampipe.io/plugins/turbot/aws#postgres-fdw)
- [Official project or provider page](https://hub.steampipe.io/plugins/turbot/aws)

`steampipe_postgres_aws` — Steampipe AWS native PostgreSQL FDW for querying live AWS APIs as foreign tables.

The reviewed catalog snapshot records version `0.125.0`, kind `standard`, and implementation language `Go`.
The curated compatibility set is `14,15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "steampipe_postgres_aws";
SELECT extversion
FROM pg_extension
WHERE extname = 'steampipe_postgres_aws';
```

The upstream project is associated with `Turbot`; verify its current support, license, packaging, and deployment boundary from the linked source.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
