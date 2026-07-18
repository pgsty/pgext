## Usage

Sources:

- [Official extension control file](https://github.com/b4oshany/livewire/blob/547e59a951fe5d088d89039a2d078186e73c7862/livewire.control)

`livewire` — power delivery modeling

The reviewed catalog snapshot records version `0.2.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `pgrouting`, `plpgsql`, `postgis`, `postgis_topology`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "livewire";
SELECT extversion
FROM pg_extension
WHERE extname = 'livewire';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
