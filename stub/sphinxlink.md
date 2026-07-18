## Usage

Sources:

- [Official extension control file](https://github.com/dimv36/sphinxlink/blob/19760a52df0e2d900bee06c7a0c59896ea408bbd/sphinxlink.control)
- [Official upstream documentation](https://github.com/dimv36/sphinxlink/blob/19760a52df0e2d900bee06c7a0c59896ea408bbd/README.md)

`sphinxlink` — C extension for issuing SphinxSearch or ManticoreSearch queries from PostgreSQL through the MySQL-compatible protocol.

The reviewed catalog snapshot records version `1.4`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "sphinxlink";
SELECT extversion
FROM pg_extension
WHERE extname = 'sphinxlink';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
